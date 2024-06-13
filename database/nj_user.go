package database

import (
	"cloud.google.com/go/firestore"
	"context"
	"encoding/json"
	"errors"
	"github.com/ninjahome/web-bridge/util"
	"google.golang.org/api/iterator"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"strings"
	"time"
)

type NinjaUsrInfo struct {
	Address      string `json:"address" firestore:"address"`
	EthAddr      string `json:"eth_addr" firestore:"eth_addr"`
	CreateAt     int64  `json:"create_at" firestore:"create_at"`
	SignInAt     int64  `json:"signIn_at" firestore:"signIn_at"`
	TwID         string `json:"tw_id" firestore:"tw_id"`
	UpdateAt     int64  `json:"update_at"`
	TweetCount   int    `json:"tweet_count" firestore:"tweet_count"`
	VoteCount    int    `json:"vote_count" firestore:"vote_count"`
	BeVotedCount int    `json:"be_voted_count" firestore:"be_voted_count"`
	Points       int    `json:"points"  firestore:"points"`
	IsElder      bool   `json:"is_elder" firestore:"is_elder"`
	ReferrerCode string `json:"referrer_code" firestore:"referrer_code"`
	SelfRefCode  string `json:"self_ref_code" firestore:"self_ref_code"`
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

func (dm *DbManager) NjUserSignIn(ethAddr, Referer string) *NinjaUsrInfo {
	ethAddr = strings.ToLower(ethAddr)
	signInTime := time.Now().UnixMilli()
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
		updateOps := []firestore.Update{
			{Path: "signIn_at", Value: signInTime},
		}

		if len(nu.SelfRefCode) == 0 {
			updateOps = append(updateOps, firestore.Update{Path: "self_ref_code", Value: ethAddr[len(ethAddr)-6:]})
		}
		if len(Referer) > 0 && len(nu.ReferrerCode) == 0 {
			updateOps = append(updateOps, firestore.Update{Path: "referrer_code", Value: Referer})
		}
		_, _ = docRef.Update(opCtx, updateOps)
		nu.SignInAt = signInTime
		util.LogInst().Debug().Str("eth-addr", ethAddr).Int64("sign-at", nu.SignInAt).Msg("firestore load ninja user info success")
		return nu
	}

	if status.Code(err) != codes.NotFound {
		util.LogInst().Err(err).Str("eth-addr", ethAddr).Msg("query firestore failed")
		return nil
	}

	nu = &NinjaUsrInfo{
		EthAddr:     ethAddr,
		CreateAt:    signInTime,
		SelfRefCode: ethAddr[len(ethAddr)-6:],
	}

	if len(Referer) > 0 {
		nu.ReferrerCode = Referer
	}
	_, err = docRef.Set(opCtx, nu)
	if err != nil {
		util.LogInst().Err(err).Str("eth-addr", ethAddr).Msg("Set firestore data as NinjaUsrInfo failed")
		return nil
	}
	util.LogInst().Debug().Str("eth-addr", ethAddr).Msg("firestore create ninja user success")
	return nu
}

func (dm *DbManager) QueryNjUsrByReferrer(referrer string) (*NinjaUsrInfo, error) {
	opCtx, cancel := context.WithTimeout(dm.ctx, DefaultDBTimeOut)
	defer cancel()
	query := dm.fileCli.Collection(DBTableNJUser).Where("self_ref_code", "==", referrer).Limit(1)
	iter := query.Documents(opCtx)
	defer iter.Stop()

	doc, err := iter.Next()
	if err != nil {
		if errors.Is(err, iterator.Done) {
			return nil, status.Error(codes.NotFound, "no Ninja User found with referrer: "+referrer)
		}
		util.LogInst().Err(err).Str("self_ref_code", referrer).Msg("query ninja user by referrer failed")
		return nil, err
	}

	var nu NinjaUsrInfo
	err = doc.DataTo(&nu)
	if err != nil {
		util.LogInst().Err(err).Str("self_ref_code", referrer).Msg("data to nj user err")
		return nil, err
	}
	return &nu, nil
}

func (dm *DbManager) QueryNjUsrById(web3ID string) (*NinjaUsrInfo, error) {
	opCtx, cancel := context.WithTimeout(dm.ctx, DefaultDBTimeOut)
	defer cancel()
	docRef := dm.fileCli.Collection(DBTableNJUser).Doc(strings.ToLower(web3ID))
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

func (dm *DbManager) MostVotedKol(pageSize int, startID int64, vote bool) ([]*NinjaUsrInfo, error) {
	opCtx, cancel := context.WithTimeout(dm.ctx, DefaultDBTimeOut)
	defer cancel()

	voteRef := dm.fileCli.Collection(DBTableNJUser)
	var query = voteRef.Limit(pageSize)
	var key = "be_voted_count"
	if vote {
		key = "vote_count"
	}
	if startID == 0 {
		query = query.OrderBy(key, firestore.Desc)
	} else {
		query = query.Where(key, "<", startID).OrderBy(key, firestore.Desc)
	}
	query = query.Where(key, ">", 0)
	var iter = query.Documents(opCtx)
	defer iter.Stop()

	var users = make([]*NinjaUsrInfo, 0)

	for {
		doc, err := iter.Next()
		if errors.Is(err, iterator.Done) {
			return users, nil
		}
		if err != nil {
			util.LogInst().Err(err).Msgf("Failed to iterate: %v", err)
			return nil, err
		}

		var usr NinjaUsrInfo
		err = doc.DataTo(&usr)
		if err != nil {
			util.LogInst().Err(err).Msgf("Failed to convert document to NinjaUsrInfo: %v", err)
			return nil, err
		}
		users = append(users, &usr)
	}
}

func (dm *DbManager) CheckKolElder() {
	opCtx, cancel := context.WithTimeout(dm.ctx, DefaultDBTimeOut*10)
	defer cancel()
	util.LogInst().Debug().Msg("start to check elder status")
	err := dm.fileCli.RunTransaction(opCtx, func(ctx context.Context, tx *firestore.Transaction) error {

		randomDoc := dm.fileCli.Collection(DBTableNJUser)
		var query = randomDoc.Where("be_voted_count", ">=", __dbConf.ElderNoFirstGot).
			OrderBy("be_voted_count", firestore.Desc).
			Limit(10)

		iter := query.Documents(opCtx)
		defer iter.Stop()
		var toBeElder = make([]NinjaUsrInfo, 0)

		for {
			doc, err := iter.Next()
			if errors.Is(err, iterator.Done) {
				util.LogInst().Debug().Msg("query kol status success")
				break
			}
			if err != nil {
				util.LogInst().Err(err).Msg("query kol status failed")
				return err
			}

			var njObj NinjaUsrInfo
			err = doc.DataTo(&njObj)
			if err != nil {
				util.LogInst().Err(err).Msg("parse kol failed")
				return err
			}
			if njObj.IsElder == false {
				toBeElder = append(toBeElder, njObj)
			}
		}

		if len(toBeElder) == 0 {
			util.LogInst().Debug().Msg("no need to update elder status")
			return nil
		}
		util.LogInst().Debug().Msgf("elder no:%d to add", len(toBeElder))

		for _, njObj := range toBeElder {
			docRef := dm.fileCli.Collection(DBTableNJUser).Doc(njObj.EthAddr)
			errUpdate := tx.Update(docRef, []firestore.Update{{Path: "is_elder", Value: true}})
			if errUpdate != nil {
				util.LogInst().Err(errUpdate).Str("eth-addr", njObj.EthAddr).Msg("update elder status failed")
				continue
			}
			util.LogInst().Info().Str("eth-addr", njObj.EthAddr).Msg("update elder status success")
		}

		return nil
	})
	if err != nil {
		util.LogInst().Err(err).Msg("update elder transaction failed")
		return
	}
}
