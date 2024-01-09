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

type NjTweetStatus struct {
	CreateTime int64 `json:"create_time" firestore:"create_time"`
	VoteCount  int   `json:"vote_count" firestore:"vote_count"`
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

	var ts = NjTweetStatus{content.CreateAt, 0}

	statusDoc := dm.fileCli.Collection(DBTableTweetsStatus).Doc(createTime)
	_, err = statusDoc.Set(opCtx, ts)
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
func (dm *DbManager) UpdateTweetVoteStatic(createAt int64, amount int) (int, error) {
	opCtx, cancel := context.WithTimeout(dm.ctx, DefaultDBTimeOut)
	defer cancel()
	docRef := dm.fileCli.Collection(DBTableTweetsStatus).Doc(fmt.Sprintf("%d", createAt))

	docSnapshot, err := docRef.Get(opCtx)
	if err != nil {
		if status.Code(err) != codes.NotFound {
			util.LogInst().Err(err).Int64("createAt", createAt).Msg("Failed to get tweet document to update vote")
			return 0, err
		}

		_, err := docRef.Set(opCtx, map[string]interface{}{
			"vote_count":  amount,
			"create_time": createAt,
		})
		if err != nil {
			util.LogInst().Err(err).Int64("createAt", createAt).Msg("Failed to create new document")
			return 0, err
		}
		return amount, nil
	}

	var existingData NjTweetStatus
	err = docSnapshot.DataTo(&existingData)
	if err != nil {
		util.LogInst().Err(err).Int64("createAt", createAt).Msg("Failed to decode document data")
		return 0, err
	}

	newFieldValue := existingData.VoteCount + amount
	_, err = docRef.Update(opCtx, []firestore.Update{
		{Path: "vote_count", Value: newFieldValue},
	})

	return newFieldValue, err
}

func (dm *DbManager) QueryTweetStatus(createTimes []int64) (map[int64]*NjTweetStatus, error) {
	opCtx, cancel := context.WithTimeout(dm.ctx, DefaultDBTimeOut)
	defer cancel()
	iter := dm.fileCli.Collection(DBTableTweetsStatus).Where("create_time", "in", createTimes).Documents(opCtx)
	result := make(map[int64]*NjTweetStatus)

	for {
		doc, err := iter.Next()
		if errors.Is(err, iterator.Done) {
			return result, nil
		}
		if err != nil {
			util.LogInst().Err(err).Msg("failed to fetch tweet status data")
			return nil, err
		}
		var ts NjTweetStatus
		err = doc.DataTo(&ts)
		if err != nil {
			util.LogInst().Err(err).Msg("data to tweet status object failed")
			return nil, err
		}
		result[ts.CreateTime] = &ts
	}
}
