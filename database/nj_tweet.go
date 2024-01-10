package database

import (
	"cloud.google.com/go/firestore"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/ninjahome/web-bridge/util"
	"google.golang.org/api/iterator"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type TxStatus int8

const (
	TxStNotPay TxStatus = iota
	TxStPending
	TxStSuccess
	TxStFailed
)

func (ts TxStatus) String() string {
	switch ts {
	case TxStNotPay:
		return "not paid"
	case TxStPending:
		return "pending"
	case TxStSuccess:
		return "success"
	case TxStFailed:
		return "failed"
	default:
		return "unknown"
	}
}

type NinjaTweet struct {
	Txt           string   `json:"text" firestore:"text"`
	CreateAt      int64    `json:"create_time" firestore:"create_time"`
	Web3ID        string   `json:"web3_id" firestore:"web3_id"`
	TweetUsrId    string   `json:"twitter_id" firestore:"twitter_id"`
	TweetId       string   `json:"tweet_id,omitempty" firestore:"tweet_id"`
	Signature     string   `json:"signature,omitempty" firestore:"signature"`
	PrefixedHash  string   `json:"prefixed_hash" firestore:"prefixed_hash"`
	PaymentStatus TxStatus `json:"payment_status" firestore:"payment_status"`
	VoteCount     int      `json:"vote_count" firestore:"vote_count"`
}

type TweetVoted struct {
	VotedTweets map[string]int `json:"voted_tweets" firestore:"voted_tweets"`
}

type TweetQueryParm struct {
	StartID  int64   `json:"start_id"`
	Web3ID   string  `json:"web3_id"`
	Newest   bool    `json:"newest"`
	VotedIDs []int64 `json:"voted_i_ds"`
}

func (p *TweetQueryParm) String() string {
	bts, _ := json.Marshal(p)
	return string(bts)
}

func (p *TweetQueryParm) createFilter(pageSize int, doc *firestore.CollectionRef, opCtx context.Context) *firestore.DocumentIterator {

	if len(p.VotedIDs) > 0 {
		return doc.Where("create_time", "in", p.VotedIDs).Documents(opCtx)
	}

	var query = doc.Limit(pageSize)

	if len(p.Web3ID) == 0 {
		query = query.Where("payment_status", "==", TxStSuccess)
	} else {
		query = query.Where("web3_id", "==", p.Web3ID)
	}

	if p.Newest {
		query = query.Where("create_time", ">", p.StartID).OrderBy("create_time", firestore.Asc)
	} else {
		query = query.Where("create_time", "<", p.StartID).OrderBy("create_time", firestore.Desc)
	}

	return query.Documents(opCtx)
}

type TweetsOfUser struct {
	Tweets map[string]struct{} `json:"tweets"  firestore:"tweets"`
}

func (nt *NinjaTweet) IsValid() bool {
	return nt.CreateAt > 0 && len(nt.Txt) > 0 &&
		len(nt.TweetUsrId) > 0 && len(nt.Web3ID) > 0
}

func (nt *NinjaTweet) String() string {
	bts, _ := json.Marshal(nt)
	return string(bts)
}

func (dm *DbManager) SaveTweet(content *NinjaTweet) error {
	opCtx, cancel := context.WithTimeout(dm.ctx, DefaultDBTimeOut)
	defer cancel()

	var createTime = fmt.Sprintf("%d", content.CreateAt)
	tweetsDoc := dm.fileCli.Collection(DBTableTweetsPosted).Doc(createTime)

	_, err := tweetsDoc.Set(opCtx, content)
	if err != nil {
		util.LogInst().Err(err).Msg("save ninja tweet failed:" + content.String())
		return err
	}
	return err
}

func (dm *DbManager) QueryTweetsByFilter(pageSize int, param *TweetQueryParm) ([]*NinjaTweet, error) {

	opCtx, cancel := context.WithTimeout(dm.ctx, DefaultDBTimeOut)
	defer cancel()

	var doc = dm.fileCli.Collection(DBTableTweetsPosted)

	var iter = param.createFilter(pageSize, doc, opCtx)
	defer iter.Stop()

	var tweets = make([]*NinjaTweet, 0)
	for {
		doc, err := iter.Next()
		if errors.Is(err, iterator.Done) {
			return tweets, nil
		}
		if err != nil {
			util.LogInst().Err(err).Msgf("Failed to iterate: %v", err)
			return nil, err
		}

		var tweet NinjaTweet
		err = doc.DataTo(&tweet)
		if err != nil {
			util.LogInst().Err(err).Msgf("Failed to convert document to NinjaUsrInfo: %v", err)
			return nil, err
		}
		tweets = append(tweets, &tweet)
	}
}

func (dm *DbManager) NjTweetDetails(createAt int64) (*NinjaTweet, error) {
	opCtx, cancel := context.WithTimeout(dm.ctx, DefaultDBTimeOut)
	defer cancel()

	var doc = dm.fileCli.Collection(DBTableTweetsPosted).Doc(fmt.Sprintf("%d", createAt))
	data, err := doc.Get(opCtx)
	if err != nil {
		util.LogInst().Err(err).Int64("create_time", createAt).Msg("query ninja tweet detail failed")
		return nil, err
	}
	var obj NinjaTweet
	err = data.DataTo(&obj)
	return &obj, err
}

func (dm *DbManager) UpdateTweetPaymentStatus(createAt int64, s TxStatus) error {
	opCtx, cancel := context.WithTimeout(dm.ctx, DefaultDBTimeOut)
	defer cancel()
	tweetsDoc := dm.fileCli.Collection(DBTableTweetsPosted).Doc(fmt.Sprintf("%d", createAt))
	_, err := tweetsDoc.Update(opCtx, []firestore.Update{
		{Path: "payment_status", Value: s},
	})
	return err
}

func (dm *DbManager) UpdateTweetVoteStatic(createAt int64, amount int, voter string) (int, error) {
	opCtx, cancel := context.WithTimeout(dm.ctx, DefaultDBTimeOut)
	defer cancel()

	var newFieldValue int
	err := dm.fileCli.RunTransaction(opCtx, func(ctx context.Context, tx *firestore.Transaction) error {
		createTime := fmt.Sprintf("%d", createAt)
		tweetDoc := dm.fileCli.Collection(DBTableTweetsPosted).Doc(createTime)
		voteDoc := dm.fileCli.Collection(DBTableTweetsVoted).Doc(voter)

		tweetSnapshot, err := tx.Get(tweetDoc)
		if err != nil {
			util.LogInst().Err(err).Int64("createAt", createAt).Msg("Failed to get tweet document to update vote")
			return err
		}

		var existingTweet NinjaTweet
		err = tweetSnapshot.DataTo(&existingTweet)
		if err != nil {
			util.LogInst().Err(err).Int64("createAt", createAt).Msg("Failed to decode tweet document data")
			return err
		}

		voteSnapshot, voteErr := tx.Get(voteDoc)
		var votedObj TweetVoted
		if voteErr != nil {
			if status.Code(voteErr) != codes.NotFound {
				util.LogInst().Err(voteErr).Int64("create_time", createAt).Msg("Failed to get tweet vote document")
				return voteErr
			}
			votedObj = TweetVoted{VotedTweets: make(map[string]int)}
			votedObj.VotedTweets[createTime] = amount
		} else {
			err = voteSnapshot.DataTo(&votedObj)
			if err != nil {
				util.LogInst().Err(err).Int64("create_time", createAt).Msg("parse tweet vote obj failed")
				return err
			}
			votedObj.VotedTweets[createTime] += amount
		}

		newFieldValue = existingTweet.VoteCount + amount
		err = tx.Update(tweetDoc, []firestore.Update{{Path: "vote_count", Value: newFieldValue}})
		if err != nil {
			util.LogInst().Err(err).Msg("update vote count on tweet failed")
			return err
		}

		return tx.Set(voteDoc, votedObj)
	})

	if err != nil {
		util.LogInst().Err(err).Msg("update vote status transaction failed")
		return 0, err
	}

	return newFieldValue, nil
}

func (dm *DbManager) QueryVotedTweetID(voter string) (map[string]int, error) {

	opCtx, cancel := context.WithTimeout(dm.ctx, DefaultDBTimeOut)
	defer cancel()

	voteDoc := dm.fileCli.Collection(DBTableTweetsVoted).Doc(voter)
	var votedObj TweetVoted

	voteSnapshot, voteErr := voteDoc.Get(opCtx)
	if voteErr != nil {
		if status.Code(voteErr) != codes.NotFound {
			util.LogInst().Err(voteErr).Msg("query vote status ")
			return nil, voteErr
		}
		return make(map[string]int), nil
	}

	voteErr = voteSnapshot.DataTo(&votedObj)
	if voteErr != nil {
		util.LogInst().Err(voteErr).Msg("DataTo vote status error")
		return nil, voteErr
	}

	return votedObj.VotedTweets, nil
}
