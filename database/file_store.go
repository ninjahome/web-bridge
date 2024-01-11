package database

import (
	"cloud.google.com/go/firestore"
	"context"
	"fmt"
	"github.com/ninjahome/web-bridge/util"
	"google.golang.org/api/option"
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
	DBTableWeb3Bindings       = "web3-bindings"
	DBTableTweetsPosted       = "tweets-posted"
	DBTableTweetsVoted        = "tweets-voted"
	DBTableTweetsSubStatus    = "vote-status"
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
		client, err = firestore.NewClientWithDatabase(ctx, __dbConf.ProjectID, __dbConf.DatabaseID)
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
