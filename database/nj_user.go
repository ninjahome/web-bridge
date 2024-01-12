package database

import (
	"context"
	"encoding/json"
	"github.com/ninjahome/web-bridge/util"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"time"
)

type NinjaUsrInfo struct {
	Address      string `json:"address" firestore:"address"`
	EthAddr      string `json:"eth_addr" firestore:"eth_addr"`
	CreateAt     int64  `json:"create_at" firestore:"create_at"`
	TwID         string `json:"tw_id" firestore:"tw_id"`
	UpdateAt     int64  `json:"update_at"`
	TweetCount   int    `json:"tweet_count" firestore:"tweet_count"`
	VoteCount    int    `json:"vote_count" firestore:"vote_count"`
	BeVotedCount int    `json:"be_voted_count" firestore:"be_voted_count"`
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

func (dm *DbManager) QueryNjUsrById(web3ID string) (*NinjaUsrInfo, error) {
	opCtx, cancel := context.WithTimeout(dm.ctx, DefaultDBTimeOut)
	defer cancel()

	docRef := dm.fileCli.Collection(DBTableNJUser).Doc(web3ID)
	doc, err := docRef.Get(opCtx)
	if err != nil {
		util.LogInst().Err(err).Str("web3-id", web3ID).Msg("query nj user data err")
		return nil, err
	}
	var nu NinjaUsrInfo
	err = doc.DataTo(&nu)
	if err != nil {
		util.LogInst().Err(err).Str("web3-id", web3ID).Msg("data to nj user err")
		return nil, err
	}
	return &nu, nil

}
