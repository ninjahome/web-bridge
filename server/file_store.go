package server

import (
	"cloud.google.com/go/firestore"
	"context"
	"encoding/json"
	"fmt"
	"github.com/ninjahome/web-bridge/util"
	"google.golang.org/api/option"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"os"
	"sync"
	"time"
)

const (
	DefaultTwitterProjectID = "dessage"
	DefaultDBTimeOut        = 10 * time.Second
	DBTableNJUser           = "ninja-user"
	DBTableTWUser           = "twitter-user"
	DBTableWeb3Bindings     = "twitter-eth-binding"
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
	if _globalCfg.LocalRun {
		_ = os.Setenv("FIRESTORE_EMULATOR_HOST", "localhost:8080")
		client, err = firestore.NewClient(ctx, _globalCfg.ProjectID)
	} else {
		client, err = firestore.NewClient(ctx, _globalCfg.ProjectID, option.WithCredentialsFile(_globalCfg.KeyFilePath))
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

type TwAPIResV2 struct {
	TwitterData *TWUserInfoV2 `json:"data"`
	EthAddr     string        `json:"eth_addr"`
	SignUpAt    int64         `json:"sign_up_at"`
}

type TWUserInfoV2 struct {
	ID              string `json:"id"`
	Name            string `json:"name"`
	Username        string `json:"username"`
	ProfileImageURL string `json:"profile_image_url"`
	Description     string `json:"description"`
}

func (t *TWUserInfoV2) String() string {
	bts, _ := json.Marshal(t)
	return string(bts)
}

func TWUsrInfoMustV2(str string) *TWUserInfoV2 {
	t := &TWUserInfoV2{}
	err := json.Unmarshal([]byte(str), t)
	if err != nil {
		return t
	}
	return t
}

type TwAPIResponse struct {
	TwitterData *TWUserInfo `json:"data"`
	EthAddr     string      `json:"eth_addr"`
	SignUpAt    int64       `json:"sign_up_at"`
}
type Web3Binding struct {
	TwitterID string `json:"twitter_id"`
	EthAddr   string `json:"eth_addr"`
	SignUpAt  int64  `json:"sign_up_at"`
	Signature string `json:"signature"`
}

type TWUserInfo struct {
	ID                   string `json:"id_str"`
	Name                 string `json:"name"`
	ScreenName           string `json:"screen_name"`
	Description          string `json:"description"`
	ProfileImageUrlHttps string `json:"profile_image_url_https"`
}

func (t *TWUserInfo) String() string {
	bts, _ := json.Marshal(t)
	return string(bts)
}

func (t *TWUserInfo) RawData() []byte {
	bts, _ := json.Marshal(t)
	return bts
}

func TWUsrInfoMust(str string) *TWUserInfo {
	t := &TWUserInfo{}
	err := json.Unmarshal([]byte(str), t)
	if err != nil {
		return t
	}
	return t
}

/*******************************************************************************************************
*
* Ninja Protocol User Infos
*
 ******************************************************************************************************/

type NinjaUsrInfo struct {
	Address  string `json:"address"`
	EthAddr  string `json:"eth_addr"`
	CreateAt int64  `json:"create_at"`
	TwID     string `json:"tw_id"`
}

func (nu *NinjaUsrInfo) String() string {
	bts, _ := json.Marshal(nu)
	return string(bts)
}

func (nu *NinjaUsrInfo) RawData() []byte {
	bts, _ := json.Marshal(nu)
	return bts
}

func NJUsrInfoMust(str string) *NinjaUsrInfo {
	nu := &NinjaUsrInfo{}
	err := json.Unmarshal([]byte(str), nu)
	if err != nil {
		return nu
	}
	return nu
}

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

	/*check nj user basic data*/
	njUserDoc := dm.fileCli.Collection(DBTableNJUser).Doc(ethAddr)
	doc, err := njUserDoc.Get(opCtx)
	if err != nil {
		util.LogInst().Err(err).Str("eth-addr", ethAddr).Msg("no nj user sign in data")
		return nil, err
	}
	nu := &NinjaUsrInfo{}
	err = doc.DataTo(nu)
	if err != nil {
		util.LogInst().Err(err).Str("eth-addr", ethAddr).Msg("parse nj user failed")
		return nil, err
	}
	if len(nu.TwID) > 0 {
		util.LogInst().Err(err).Str("eth-addr", ethAddr).
			Str("twitter-id", nu.TwID).Msg("duplicate web3 binding")
		return nil, fmt.Errorf("duplicate web3 binding")
	}

	/*set meta binding  data*/
	twitterDoc := dm.fileCli.Collection(DBTableTWUser).Doc(twMeta.ID)
	_, err = twitterDoc.Set(opCtx, twMeta)
	if err != nil {
		util.LogInst().Err(err).Str("twitter-id", twMeta.ID).Msg("update twitter meta failed")
		return nil, err
	}

	bindDoc := dm.fileCli.Collection(DBTableWeb3Bindings).Doc(ethAddr)
	_, err = bindDoc.Set(opCtx, bindData)
	if err != nil {
		util.LogInst().Err(err).Str("eth-addr", ethAddr).
			Str("twitter-id", twMeta.ID).Msg("update web3 binding failed")
		return nil, err
	}

	/*update nj user basic data*/
	nu.TwID = bindData.TwitterID
	_, err = njUserDoc.Set(opCtx, nu)
	if err != nil {
		util.LogInst().Err(err).Str("eth-addr", ethAddr).
			Str("twitter-id", nu.TwID).Msg("update nj user failed")
		return nil, err
	}

	return nu, nil
}

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
