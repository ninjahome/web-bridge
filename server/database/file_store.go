package database

import (
	"cloud.google.com/go/firestore"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/dghubble/oauth1"
	"github.com/ninjahome/web-bridge/util"
	"golang.org/x/oauth2"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"os"
	"sync"
	"time"
)

const (
	DefaultFirestoreProjectID = "dessage"
	DefaultDatabaseID         = "dessage-release"
	DefaultDBTimeOut          = 10 * time.Second
	DBTableNJUser             = "ninja-user"
	DBTableTWUser             = "twitter-user"
	DBTableTWUserAccToken     = "twitter-user-access-token"
	DBTableTWUserAccTokenV2   = "twitter-user-access-token_v2"
	DBTableWeb3Bindings       = "twitter-eth-binding"
	DBTableTweetsPosted       = "tweets-posted"
	DBTableTweetsStatus       = "tweets-status"
)

/*******************************************************************************************************
*
* Database Logic
*
 ******************************************************************************************************/
var _dbInst *DbManager
var databaseOnce sync.Once

type DbManager struct {
	fileCli *firestore.Client
	ctx     context.Context
	cancel  context.CancelFunc
}

func DbInst() *DbManager {
	databaseOnce.Do(func() {
		_dbInst = newDb()
	})
	return _dbInst
}

func newDb() *DbManager {
	ctx, cancel := context.WithCancel(context.Background())
	var client *firestore.Client
	var err error
	if __dbConf.LocalRun {
		_ = os.Setenv("FIRESTORE_EMULATOR_HOST", "localhost:8080")
		client, err = firestore.NewClientWithDatabase(ctx, __dbConf.ProjectID, "dessage")
	} else {
		client, err = firestore.NewClientWithDatabase(ctx, __dbConf.ProjectID,
			__dbConf.DatabaseID, option.WithCredentialsFile(__dbConf.KeyFilePath))
	}
	if err != nil {
		panic(err)
	}
	var dbm = &DbManager{
		fileCli: client,
		ctx:     ctx,
		cancel:  cancel,
	}
	return dbm
}

/*******************************************************************************************************
*
* Twitter User Infos
*
 ******************************************************************************************************/

type TwAPIResponse struct {
	TwitterData *TWUserInfo `json:"data"`
	EthAddr     string      `json:"eth_addr"`
	SignUpAt    int64       `json:"sign_up_at"`
}
type Web3Binding struct {
	TwitterID string `json:"twitter_id" firestore:"twitter_id"`
	EthAddr   string `json:"eth_addr" firestore:"eth_addr"`
	SignUpAt  int64  `json:"sign_up_at" firestore:"sign_up_at"`
	Signature string `json:"signature" firestore:"signature"`
}

type TWUserInfo struct {
	ID                   string `json:"id" firestore:"id"`
	Name                 string `json:"name" firestore:"name"`
	ScreenName           string `json:"username" firestore:"username"`
	Description          string `json:"description" firestore:"description"`
	ProfileImageUrlHttps string `json:"profile_image_url" firestore:"profile_image_url"`
	Verified             bool   `json:"verified"  firestore:"verified"`
	VoteCount            int    `json:"vote_count" firestore:"vote_count"`
	TweetCount           int    `json:"tweet_count" firestore:"vote_count"`
}

func (t *TWUserInfo) String() string {
	bts, _ := json.Marshal(t)
	return string(bts)
}

func (t *TWUserInfo) RawData() []byte {
	bts, _ := json.Marshal(t)
	return bts
}

/*******************************************************************************************************
*
* Ninja Protocol User Infos
*
 ******************************************************************************************************/

type NinjaUsrInfo struct {
	Address  string `json:"address" firestore:"address"`
	EthAddr  string `json:"eth_addr" firestore:"eth_addr"`
	CreateAt int64  `json:"create_at" firestore:"create_at"`
	TwID     string `json:"tw_id" firestore:"tw_id"`
	UpdateAt int64
}

func (nu *NinjaUsrInfo) String() string {
	bts, _ := json.Marshal(nu)
	return string(bts)
}

func (nu *NinjaUsrInfo) RawData() []byte {
	bts, _ := json.Marshal(nu)
	return bts
}

func (nu *NinjaUsrInfo) RefreshSession() {
	nu.UpdateAt = time.Now().UnixMilli()
}

func NJUsrInfoMust(data []byte) (*NinjaUsrInfo, error) {
	nu := &NinjaUsrInfo{}
	err := json.Unmarshal(data, nu)
	if err != nil {
		return nil, err
	}
	return nu, err
}

/*******************************************************************************************************
*
* Twitter User Token
*
 ******************************************************************************************************/

type TwUserAccessToken struct {
	OauthToken       string `json:"oauth_token" firestore:"oauth_token"`
	OauthTokenSecret string `json:"oauth_token_secret" firestore:"oauth_token_secret"`
	UserId           string `json:"user_id" firestore:"user_id"`
	ScreenName       string `json:"screen_name" firestore:"screen_name"`
}

type TwUserAccessTokenV2 struct {
	UserId string `json:"user_id" firestore:"user_id"`
	*oauth2.Token
}

func (ut *TwUserAccessToken) GetToken() *oauth1.Token {
	return oauth1.NewToken(ut.OauthToken, ut.OauthTokenSecret)
}

func (ut *TwUserAccessToken) String() string {
	bts, _ := json.Marshal(ut)
	return string(bts)
}

/*******************************************************************************************************
*
* Ninja tweet
*
 ******************************************************************************************************/

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

/*******************************************************************************************************
*
* Ninja Protocol User Infos
*
 ******************************************************************************************************/

func (dm *DbManager) NjUserSignIn(ethAddr string) *NinjaUsrInfo {
	nu := &NinjaUsrInfo{
		EthAddr: ethAddr,
	}
	docRef := dm.fileCli.Collection(DBTableNJUser).Doc(ethAddr)
	opCtx, cancel := context.WithTimeout(dm.ctx, DefaultDBTimeOut)
	defer cancel()
	doc, err := docRef.Get(opCtx)
	if err == nil {
		err = doc.DataTo(nu)
		if err != nil {
			util.LogInst().Err(err).Str("eth-addr", ethAddr).Msg("parse firestore data  to NinjaUsrInfo failed")
			return nil
		}
		util.LogInst().Debug().Str("eth-addr", ethAddr).Msg("firestore load ninja user info success")
		return nu
	}

	if status.Code(err) != codes.NotFound {
		util.LogInst().Err(err).Str("eth-addr", ethAddr).Msg("query firestore failed")
		return nil
	}

	nu = &NinjaUsrInfo{
		EthAddr:  ethAddr,
		CreateAt: time.Now().UnixMilli(),
	}
	_, err = docRef.Set(opCtx, nu)
	if err != nil {
		util.LogInst().Err(err).Str("eth-addr", ethAddr).Msg("Set firestore data as NinjaUsrInfo failed")
		return nil
	}

	util.LogInst().Debug().Str("eth-addr", ethAddr).Msg("firestore create ninja user success")
	return nu
}
func (dm *DbManager) BindingWeb3ID(bindData *Web3Binding, twMeta *TWUserInfo) (*NinjaUsrInfo, error) {
	var ethAddr = bindData.EthAddr
	opCtx, cancel := context.WithTimeout(dm.ctx, DefaultDBTimeOut)
	defer cancel()

	err := dm.fileCli.RunTransaction(opCtx, func(ctx context.Context, tx *firestore.Transaction) error {
		// 检查 NJ 用户基本数据
		njUserDoc := dm.fileCli.Collection(DBTableNJUser).Doc(ethAddr)
		doc, err := tx.Get(njUserDoc)
		if err != nil {
			util.LogInst().Err(err).Str("eth-addr", ethAddr).Msg("no nj user sign in data")
			return err
		}
		nu := &NinjaUsrInfo{}
		err = doc.DataTo(nu)
		if err != nil {
			util.LogInst().Err(err).Str("eth-addr", ethAddr).Msg("parse nj user failed")
			return err
		}
		if len(nu.TwID) > 0 {
			util.LogInst().Err(err).Str("eth-addr", ethAddr).
				Str("twitter-id", nu.TwID).Msg("duplicate web3 binding")
			return fmt.Errorf("duplicate web3 binding")
		}

		// 设置 Twitter 元数据绑定
		twitterDoc := dm.fileCli.Collection(DBTableTWUser).Doc(twMeta.ID)
		err = tx.Set(twitterDoc, twMeta)
		if err != nil {
			util.LogInst().Err(err).Str("twitter-id", twMeta.ID).Msg("update twitter meta failed")
			return err
		}

		// 设置 Web3 绑定
		bindDoc := dm.fileCli.Collection(DBTableWeb3Bindings).Doc(ethAddr)
		err = tx.Set(bindDoc, bindData)
		if err != nil {
			util.LogInst().Err(err).Str("eth-addr", ethAddr).
				Str("twitter-id", twMeta.ID).Msg("update web3 binding failed")
			return err
		}

		// 更新 NJ 用户基本数据
		nu.TwID = bindData.TwitterID
		err = tx.Set(njUserDoc, nu, firestore.Merge([]string{"tw_id"}))
		if err != nil {
			util.LogInst().Err(err).Str("eth-addr", ethAddr).
				Str("twitter-id", nu.TwID).Msg("update nj user failed")
			return err
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	njUserDoc := dm.fileCli.Collection(DBTableNJUser).Doc(ethAddr)
	docSnapshot, err := njUserDoc.Get(opCtx)
	if err != nil {
		util.LogInst().Err(err).Str("eth-addr", ethAddr).Msg("Failed to get updated nj user data")
		return nil, err
	}

	updatedNu := &NinjaUsrInfo{}
	err = docSnapshot.DataTo(updatedNu)
	if err != nil {
		util.LogInst().Err(err).Str("eth-addr", ethAddr).Msg("Failed to decode updated nj user data")
		return nil, err
	}

	return updatedNu, nil
}

//
//func (dm *DbManager) BindingWeb3ID(bindData *Web3Binding, twMeta *TWUserInfo) (*NinjaUsrInfo, error) {
//	var ethAddr = bindData.EthAddr
//	opCtx, cancel := context.WithTimeout(dm.ctx, DefaultDBTimeOut)
//	defer cancel()
//
//	/*check nj user basic data*/
//	njUserDoc := dm.fileCli.Collection(DBTableNJUser).Doc(ethAddr)
//	doc, err := njUserDoc.Get(opCtx)
//	if err != nil {
//		util.LogInst().Err(err).Str("eth-addr", ethAddr).Msg("no nj user sign in data")
//		return nil, err
//	}
//	nu := &NinjaUsrInfo{}
//	err = doc.DataTo(nu)
//	if err != nil {
//		util.LogInst().Err(err).Str("eth-addr", ethAddr).Msg("parse nj user failed")
//		return nil, err
//	}
//	if len(nu.TwID) > 0 {
//		util.LogInst().Err(err).Str("eth-addr", ethAddr).
//			Str("twitter-id", nu.TwID).Msg("duplicate web3 binding")
//		return nil, fmt.Errorf("duplicate web3 binding")
//	}
//
//	/*set meta binding  data*/
//	twitterDoc := dm.fileCli.Collection(DBTableTWUser).Doc(twMeta.ID)
//	_, err = twitterDoc.Set(opCtx, twMeta)
//	if err != nil {
//		util.LogInst().Err(err).Str("twitter-id", twMeta.ID).Msg("update twitter meta failed")
//		return nil, err
//	}
//
//	bindDoc := dm.fileCli.Collection(DBTableWeb3Bindings).Doc(ethAddr)
//	_, err = bindDoc.Set(opCtx, bindData)
//	if err != nil {
//		util.LogInst().Err(err).Str("eth-addr", ethAddr).
//			Str("twitter-id", twMeta.ID).Msg("update web3 binding failed")
//		return nil, err
//	}
//
//	/*update nj user basic data*/
//	nu.TwID = bindData.TwitterID
//	_, err = njUserDoc.Set(opCtx, nu, firestore.Merge([]string{"tw_id"}))
//	if err != nil {
//		util.LogInst().Err(err).Str("eth-addr", ethAddr).
//			Str("twitter-id", nu.TwID).Msg("update nj user failed")
//		return nil, err
//	}
//
//	return nu, nil
//}

func (dm *DbManager) TwitterBasicInfo(TID string) (*TWUserInfo, error) {
	opCtx, cancel := context.WithTimeout(dm.ctx, DefaultDBTimeOut)
	defer cancel()
	twitterDoc := dm.fileCli.Collection(DBTableTWUser).Doc(TID)
	doc, err := twitterDoc.Get(opCtx)
	if err != nil {
		util.LogInst().Err(err).Str("twitter-id", TID).Msg("twitterDoc get failed")
		return nil, err
	}
	tu := &TWUserInfo{}
	err = doc.DataTo(tu)
	if err != nil {
		util.LogInst().Err(err).Str("twitter-id", TID).Msg("twitter Doc to TWUserInfo failed")
		return nil, err
	}
	return tu, nil
}

func (dm *DbManager) UpdateBasicInfo(twData *TWUserInfo) error {
	opCtx, cancel := context.WithTimeout(dm.ctx, DefaultDBTimeOut)
	defer cancel()
	twitterDoc := dm.fileCli.Collection(DBTableTWUser).Doc(twData.ID)
	_, err := twitterDoc.Set(opCtx, twData)
	if err != nil {
		util.LogInst().Err(err).Str("twitter-id", twData.ID).Msg("twitterDoc get failed")
		return err
	}
	util.LogInst().Debug().Str("twitter-id", twData.ID).Msg("update twitter user data success")
	return nil
}

func (dm *DbManager) SaveTwAccessToken(token *TwUserAccessToken) error {
	opCtx, cancel := context.WithTimeout(dm.ctx, DefaultDBTimeOut)
	defer cancel()
	tokenDoc := dm.fileCli.Collection(DBTableTWUserAccToken).Doc(token.UserId)
	_, err := tokenDoc.Set(opCtx, token)
	return err
}
func (dm *DbManager) GetTwAccessToken(twitterId string) (*TwUserAccessToken, error) {
	opCtx, cancel := context.WithTimeout(dm.ctx, DefaultDBTimeOut)
	defer cancel()
	tokenDoc := dm.fileCli.Collection(DBTableTWUserAccToken).Doc(twitterId)
	doc, err := tokenDoc.Get(opCtx)
	if err != nil {
		util.LogInst().Err(err).Str("twitter-id", twitterId).Msg("load twitter access token failed")
		return nil, err
	}
	var token TwUserAccessToken
	err = doc.DataTo(&token)
	if err != nil {
		util.LogInst().Err(err).Str("twitter-id", twitterId).Msg("parse twitter access token failed")
		return nil, err
	}
	return &token, nil
}

func (dm *DbManager) SaveTwAccessTokenV2(token *TwUserAccessTokenV2) error {
	opCtx, cancel := context.WithTimeout(dm.ctx, DefaultDBTimeOut)
	defer cancel()
	tokenDoc := dm.fileCli.Collection(DBTableTWUserAccTokenV2).Doc(token.UserId)
	_, err := tokenDoc.Set(opCtx, token)
	return err
}

func (dm *DbManager) GetTwAccessTokenV2(twitterId string) (*TwUserAccessTokenV2, error) {
	opCtx, cancel := context.WithTimeout(dm.ctx, DefaultDBTimeOut)
	defer cancel()
	tokenDoc := dm.fileCli.Collection(DBTableTWUserAccTokenV2).Doc(twitterId)
	doc, err := tokenDoc.Get(opCtx)
	if err != nil {
		util.LogInst().Err(err).Str("twitter-id", twitterId).Msg("load twitter access token failed")
		return nil, err
	}
	var token = &TwUserAccessTokenV2{
		Token: &oauth2.Token{},
	}
	err = doc.DataTo(token)
	if err != nil {
		util.LogInst().Err(err).Str("twitter-id", twitterId).Msg("parse twitter access token failed")
		return nil, err
	}
	return token, nil
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

func (dm *DbManager) QueryGlobalLatestTweets(pageSize int, id int64, readNewest bool, callback func(tweet *NinjaTweet)) error {
	opCtx, cancel := context.WithTimeout(dm.ctx, DefaultDBTimeOut)
	defer cancel()

	var doc = dm.fileCli.Collection(DBTableTweetsPosted)
	var iter *firestore.DocumentIterator
	if readNewest {
		iter = doc.
			Where("create_time", ">", id).
			OrderBy("create_time", firestore.Asc).Limit(pageSize).Documents(opCtx)
	} else {
		iter = doc.
			Where("create_time", "<", id).
			OrderBy("create_time", firestore.Desc).Limit(pageSize).Documents(opCtx)
	}

	defer iter.Stop()
	for {
		doc, err := iter.Next()
		if errors.Is(err, iterator.Done) {
			return nil
		}
		if err != nil {
			util.LogInst().Err(err).Msgf("Failed to iterate: %v", err)
			return err
		}

		var tweet NinjaTweet
		err = doc.DataTo(&tweet)
		if err != nil {
			util.LogInst().Err(err).Msgf("Failed to convert document to NinjaUsrInfo: %v", err)
			return err
		}
		callback(&tweet)
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
		result[ts.CreateTime] = &ts // 假设 NjTweetStatus 结构体中有 CreateTime 字段
	}
}
