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

type TweetQueryParm struct {
	StartID  int64   `json:"start_id"`
	Web3ID   string  `json:"web3_id"`
	VotedIDs []int64 `json:"voted_ids"`
}

type TweetVoted struct {
	CreateTime int64 `json:"create_time" firestore:"create_time"`
	VoteCount  int   `json:"vote_count" firestore:"vote_count"`
}

func (p *TweetQueryParm) String() string {
	bts, _ := json.Marshal(p)
	return string(bts)
}

func (p *TweetQueryParm) createFilter(pageSize int, doc *firestore.CollectionRef, opCtx context.Context) *firestore.DocumentIterator {

	if len(p.VotedIDs) > 0 {
		return doc.Where("create_time", "in", p.VotedIDs).OrderBy("create_time", firestore.Desc).Documents(opCtx)
	}

	var query = doc.Limit(pageSize)

	if len(p.Web3ID) == 0 {
		query = query.Where("payment_status", "==", TxStSuccess)
	} else {
		query = query.Where("web3_id", "==", p.Web3ID)
	}

	if p.StartID == 0 {
		query = query.OrderBy("create_time", firestore.Desc)
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

func (dm *DbManager) DelUnpaidTweet(createTime int64, addr string) error {
	opCtx, cancel := context.WithTimeout(dm.ctx, DefaultDBTimeOut)
	defer cancel()

	query := dm.fileCli.Collection(DBTableTweetsPosted).
		Where("payment_status", "==", TxStNotPay).
		Where("create_time", "==", createTime).
		Where("web3_id", "==", addr)

	iter := query.Documents(opCtx)
	defer iter.Stop()

	doc, err := iter.Next()
	if err != nil {
		util.LogInst().Err(err).Msg("no such item to delete")
		return err
	}

	_, err = doc.Ref.Delete(opCtx)
	return err
}

type TweetVoteAction struct {
	CreateTime    int64 `json:"create_time"`
	VoteCount     int   `json:"vote_count"`
	UserVoteCount int   `json:"user_vote_count"`
}

func (dm *DbManager) UpdateTweetVoteStatic(vote *TweetVoteAction, voter string) error {
	opCtx, cancel := context.WithTimeout(dm.ctx, DefaultDBTimeOut)
	defer cancel()

	err := dm.fileCli.RunTransaction(opCtx, func(ctx context.Context, tx *firestore.Transaction) error {
		createTime := fmt.Sprintf("%d", vote.CreateTime)
		tweetDoc := dm.fileCli.Collection(DBTableTweetsPosted).Doc(createTime)
		voteDoc := dm.fileCli.Collection(DBTableTweetsVoted).Doc(voter).Collection(DBTableTweetsSubStatus).Doc(createTime)

		tweetSnapshot, err := tx.Get(tweetDoc)
		if err != nil {
			util.LogInst().Err(err).Int64("createAt", vote.CreateTime).Msg("Failed to get tweet document to update vote")
			return err
		}

		var existingTweet NinjaTweet
		err = tweetSnapshot.DataTo(&existingTweet)
		if err != nil {
			util.LogInst().Err(err).Int64("createAt", vote.CreateTime).Msg("Failed to decode tweet document data")
			return err
		}

		voteSnapshot, voteErr := tx.Get(voteDoc)
		var votedObj TweetVoted
		if voteErr != nil {
			if status.Code(voteErr) != codes.NotFound {
				util.LogInst().Err(voteErr).Int64("create_time", vote.CreateTime).Msg("Failed to get tweet vote document")
				return voteErr
			}
			votedObj = TweetVoted{
				CreateTime: vote.CreateTime,
				VoteCount:  vote.VoteCount,
			}
		} else {
			err = voteSnapshot.DataTo(&votedObj)
			if err != nil {
				util.LogInst().Err(err).Int64("create_time", vote.CreateTime).Msg("parse tweet vote obj failed")
				return err
			}

			votedObj.VoteCount += vote.VoteCount
		}

		var newFieldValue = existingTweet.VoteCount + vote.VoteCount
		err = tx.Update(tweetDoc, []firestore.Update{{Path: "vote_count", Value: newFieldValue}})
		if err != nil {
			util.LogInst().Err(err).Msg("update vote count on tweet failed")
			return err
		}

		err = tx.Set(voteDoc, votedObj)
		if err != nil {
			util.LogInst().Err(err).Msg("update vote status doc err")
			return err
		}
		vote.VoteCount = newFieldValue
		vote.UserVoteCount = votedObj.VoteCount
		return nil
	})

	return err
}

func (dm *DbManager) QueryVotedTweetID(pageSize int, startID int64, voter string) ([]*TweetVoted, error) {

	opCtx, cancel := context.WithTimeout(dm.ctx, DefaultDBTimeOut)
	defer cancel()

	voteRef := dm.fileCli.Collection(DBTableTweetsVoted).Doc(voter).Collection(DBTableTweetsSubStatus)
	var query = voteRef.Limit(pageSize)

	if startID == 0 {
		query = query.OrderBy("create_time", firestore.Desc)
	} else {
		query = query.Where("create_time", "<", startID).OrderBy("create_time", firestore.Desc)
	}

	iter := query.Documents(opCtx)
	defer iter.Stop()
	var result = make([]*TweetVoted, 0)
	for {
		doc, err := iter.Next()
		if errors.Is(err, iterator.Done) {
			return result, nil

		}
		if err != nil {
			util.LogInst().Err(err).Msg("query vote status failed")
			return nil, err
		}

		var voteStatus TweetVoted
		err = doc.DataTo(&voteStatus)
		if err != nil {
			util.LogInst().Err(err).Msg("data to obj failed")
			continue
		}
		result = append(result, &voteStatus)
	}
}
