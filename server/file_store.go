package server

import (
	"cloud.google.com/go/firestore"
	"context"
	"encoding/json"
	"github.com/ninjahome/web-bridge/util"
	"google.golang.org/api/option"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"sync"
	"time"
)

const (
	DefaultTwitterProjectID = "dessage"
	DefaultDBTimeOut        = 10 * time.Second
	DBTableNJUser           = "ninja-user"
	DBTableUsrLog           = "ninja-user"
	DBTableTWUser           = "twitter-user"
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

	client, err := firestore.NewClient(ctx, _globalCfg.ProjectID, option.WithCredentialsFile(_globalCfg.KeyFilePath))
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

type TWUserInfo struct {
	ID                   string `json:"id_str"`
	Name                 string `json:"name"`
	ScreenName           string `json:"screen_name"`
	Description          string `json:"description"`
	Verified             bool   `json:"verified"`
	FollowersCount       int    `json:"followers_count"`
	FriendsCount         int    `json:"friends_count"`
	CreatedAt            string `json:"created_at"`
	ProfileImageUrlHttps string `json:"profile_image_url_https"`
}

func (t *TWUserInfo) String() string {
	bts, _ := json.Marshal(t)
	return string(bts)
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
