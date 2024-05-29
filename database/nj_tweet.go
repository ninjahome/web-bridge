package database

import (
	"cloud.google.com/go/firestore"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/ninjahome/web-bridge/util"
	"google.golang.org/api/iterator"
	"strings"
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
	Txt           string     `json:"text" firestore:"text"`
	TxtList       []string   `json:"text_list"  firestore:"-"`
	ImageThumb    [][]string `json:"-"  firestore:"-"`
	ImageHash     [][]string `json:"-"  firestore:"image_hash"`
	ImageRaw      [][]string `json:"-"  firestore:"-"`
	CreateAt      int64      `json:"create_time" firestore:"create_time"`
	Web3ID        string     `json:"web3_id" firestore:"web3_id"`
	TweetUsrId    string     `json:"twitter_id" firestore:"twitter_id"`
	TweetId       string     `json:"tweet_id,omitempty" firestore:"tweet_id"`
	Signature     string     `json:"signature,omitempty" firestore:"signature"`
	PrefixedHash  string     `json:"prefixed_hash" firestore:"prefixed_hash"`
	PaymentStatus TxStatus   `json:"payment_status" firestore:"payment_status"`
	VoteCount     int        `json:"vote_count" firestore:"vote_count"`
}

type TweetQueryParm struct {
	StartID  int64    `json:"start_id"`
	Web3ID   string   `json:"web3_id"`
	VotedIDs []int64  `json:"voted_ids"`
	HashArr  []string `json:"hash_arr"`
	IsOwner  bool     `json:"is_owner"`
}

type TweetImgRaw struct {
	Raw  string `json:"raw" firestore:"raw"`
	Hash string `json:"hash" firestore:"hash"`
}

func (p *TweetQueryParm) String() string {
	bts, _ := json.Marshal(p)
	return string(bts)
}

type TweetPaymentStatus struct {
	CreateTime int64    `json:"create_time" firestore:"create_time"`
	TxHash     string   `json:"tx_hash" firestore:"tx_hash"`
	Status     TxStatus `json:"status,omitempty"  firestore:"status"`
}

func (p *TweetQueryParm) createFilter(pageSize int, doc *firestore.CollectionRef, opCtx context.Context) *firestore.DocumentIterator {

	if len(p.VotedIDs) > 0 {
		return doc.Where("create_time", "in", p.VotedIDs).OrderBy("create_time", firestore.Desc).Documents(opCtx)
	}

	if len(p.HashArr) > 0 {
		return doc.Where("prefixed_hash", "in", p.HashArr).Documents(opCtx)
	}

	var query = doc.Limit(pageSize)

	if len(p.Web3ID) == 0 {
		query = query.Where("payment_status", "==", TxStSuccess)
	} else {
		query = query.Where("web3_id", "==", p.Web3ID)
		if p.IsOwner == false {
			query = query.Where("payment_status", "==", TxStSuccess)
		}
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
	return nt.CreateAt > 0 && len(nt.TxtList) >= 0 &&
		len(nt.TweetUsrId) > 0 && len(nt.Web3ID) > 0
}

func (nt *NinjaTweet) String() string {
	bts, _ := json.Marshal(nt)
	return string(bts)
}

func (dm *DbManager) SaveRawImg(hash, raw string) error {
	opCtx, cancel := context.WithTimeout(dm.ctx, DefaultDBTimeOut)
	defer cancel()
	imgDoc := dm.fileCli.Collection(DBTableTweetsImages).Doc(hash)

	var obj = TweetImgRaw{
		raw,
		hash,
	}
	_, err := imgDoc.Set(opCtx, obj)
	return err
}

func (dm *DbManager) GetRawImg(hash string) (*TweetImgRaw, error) {
	opCtx, cancel := context.WithTimeout(dm.ctx, DefaultDBTimeOut)
	defer cancel()
	imgDoc := dm.fileCli.Collection(DBTableTweetsImages).Doc(hash)

	docSnapshot, err := imgDoc.Get(opCtx)
	if err != nil {
		util.LogInst().Err(err).Msg("not found image raw obj :" + hash)
		return nil, err
	}
	var imgRaw TweetImgRaw
	err = docSnapshot.DataTo(&imgRaw)
	if err != nil {
		util.LogInst().Err(err).Msg("parse image raw obj failed:" + hash)
		return nil, err
	}

	return &imgRaw, nil // 返回TweetImgRaw结构体中的Raw字段
}

func (dm *DbManager) SaveThumbImg(hash, raw string) error {
	opCtx, cancel := context.WithTimeout(dm.ctx, DefaultDBTimeOut)
	defer cancel()
	imgDoc := dm.fileCli.Collection(DBTableTweetsThumbImages).Doc(hash)

	var obj = TweetImgRaw{
		raw,
		hash,
	}
	_, err := imgDoc.Set(opCtx, obj)
	return err
}

func (dm *DbManager) GetThumbImg(hash string) (*TweetImgRaw, error) {
	opCtx, cancel := context.WithTimeout(dm.ctx, DefaultDBTimeOut)
	defer cancel()
	imgDoc := dm.fileCli.Collection(DBTableTweetsThumbImages).Doc(hash)

	docSnapshot, err := imgDoc.Get(opCtx)
	if err != nil {
		util.LogInst().Err(err).Msg("not found image raw obj :" + hash)
		return nil, err
	}
	var imgRaw TweetImgRaw
	err = docSnapshot.DataTo(&imgRaw)
	if err != nil {
		util.LogInst().Err(err).Msg("parse image raw obj failed:" + hash)
		return nil, err
	}
	return &imgRaw, nil
}

func (dm *DbManager) updateNjUserForTweet(web3ID string, opCtx context.Context) error {
	docRef := dm.fileCli.Collection(DBTableNJUser).Doc(strings.ToLower(web3ID))
	var nu NinjaUsrInfo
	doc, err := docRef.Get(opCtx)
	if err != nil {
		util.LogInst().Err(err).Str("web3-id", web3ID).Msg("query nj user failed")
		return err
	}
	err = doc.DataTo(&nu)
	if err != nil {
		util.LogInst().Err(err).Str("web3-id", web3ID).Msg("parse nj user failed")
		return err
	}
	nu.TweetCount += 1
	nu.Points += __dbConf.PointForPost
	_, err = docRef.Update(opCtx, []firestore.Update{
		{Path: "tweet_count", Value: nu.TweetCount},
		{Path: "points", Value: nu.Points},
	})

	return nil
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

func (dm *DbManager) UpdateTweetPaymentStatus(status *TweetPaymentStatus, web3ID string) error {
	opCtx, cancel := context.WithTimeout(dm.ctx, DefaultDBTimeOut*2)
	defer cancel()
	tweetId := fmt.Sprintf("%d", status.CreateTime)

	tweetsDoc := dm.fileCli.Collection(DBTableTweetsPosted).Doc(tweetId)
	_, err := tweetsDoc.Update(opCtx, []firestore.Update{
		{Path: "payment_status", Value: status.Status},
	})
	if err != nil {
		util.LogInst().Err(err).Str("tweet", tweetId).Msg("update tweet status failed")
		return err
	}

	if status.Status != TxStSuccess {
		return nil
	}

	if err := dm.checkTransactionStatus(opCtx, status); err != nil {
		return err
	}
	return dm.updateNjUserForTweet(web3ID, opCtx)
}

func (dm *DbManager) checkTransactionStatus(opCtx context.Context, status *TweetPaymentStatus) error {
	statusDoc := dm.fileCli.Collection(DBTableTweetsStatus).Doc(status.TxHash)
	_, err := statusDoc.Get(opCtx)
	if err == nil {
		return fmt.Errorf("duplicate update tweet:%d with transaction:%s", status.CreateTime, status.TxHash)
	}
	_, err = statusDoc.Set(opCtx, status)
	if err != nil {
		util.LogInst().Err(err).Int64("tweet", status.CreateTime).Str("tx", status.TxHash).Msg("add tweet update failed")
		return err
	}
	return nil
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

func (dm *DbManager) QueryTwUserByTweetHash(tHash string) (*TWUserInfo, error) {
	opCtx, cancel := context.WithTimeout(dm.ctx, DefaultDBTimeOut)
	defer cancel()
	query := dm.fileCli.Collection(DBTableTweetsPosted).
		Where("prefixed_hash", "==", tHash)
	iter := query.Documents(opCtx)
	defer iter.Stop()
	doc, err := iter.Next()
	if err != nil {
		util.LogInst().Err(err).Str("tweet-hash", tHash).Msg("no such tweet")
		return nil, err
	}

	var obj NinjaTweet
	err = doc.DataTo(&obj)
	if err != nil {
		util.LogInst().Err(err).Msg("parse nj tweet failed")
		return nil, err
	}

	twitterDoc := dm.fileCli.Collection(DBTableTWUser).Doc(obj.TweetUsrId)
	doc, err = twitterDoc.Get(opCtx)
	if err != nil {
		util.LogInst().Err(err).Str("twitter-id", obj.TweetUsrId).Msg("twitterDoc get failed")
		return nil, err
	}
	tu := &TWUserInfo{}
	err = doc.DataTo(tu)
	if err != nil {
		util.LogInst().Err(err).Str("twitter-id", obj.TweetUsrId).Msg("twitter Doc to TWUserInfo failed")
		return nil, err
	}
	return tu, nil
}

func (dm *DbManager) NjTweetDetailsByHash(tHash string) (*NinjaTweet, error) {
	opCtx, cancel := context.WithTimeout(dm.ctx, DefaultDBTimeOut)
	defer cancel()
	query := dm.fileCli.Collection(DBTableTweetsPosted).
		Where("prefixed_hash", "==", tHash)
	iter := query.Documents(opCtx)
	defer iter.Stop()
	doc, err := iter.Next()
	if err != nil {
		util.LogInst().Err(err).Str("tweet-hash", tHash).Msg("no such tweet")
		return nil, err
	}

	var obj NinjaTweet
	err = doc.DataTo(&obj)
	if err != nil {
		util.LogInst().Err(err).Msg("parse nj tweet failed")
		return nil, err
	}

	return &obj, nil
}
