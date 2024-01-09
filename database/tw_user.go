package database

import (
	"context"
	"encoding/json"
	"github.com/dghubble/oauth1"
	"github.com/ninjahome/web-bridge/util"
	"golang.org/x/oauth2"
)

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
}

func (t *TWUserInfo) String() string {
	bts, _ := json.Marshal(t)
	return string(bts)
}

func (t *TWUserInfo) RawData() []byte {
	bts, _ := json.Marshal(t)
	return bts
}

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
