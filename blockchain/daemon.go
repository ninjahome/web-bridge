package blockchain

import (
	"fmt"
	"github.com/ninjahome/web-bridge/database"
	"github.com/ninjahome/web-bridge/util"
	"time"
)

type BCConf struct {
	TweeTVoteContractAddress string `json:"tweet_vote_contract_address"`
	GameContract             string `json:"game_plugin_contract_address"`
	KolKeyContractAddress    string `json:"kol_key_contract_address"`
	InfuraUrl                string `json:"infura_url"`
	GameTimeInMinute         int    `json:"game_time_in_minute,omitempty"`
	TxCheckerInSeconds       int    `json:"tx_checker_in_seconds,omitempty"`
	ChainID                  int64  `json:"chain_id,omitempty"`
	CheckTimeInSecond        int    `json:"check_time_in_second"`
}

func (c *BCConf) String() string {
	s := "\n------block chain config------"
	s += "\ntweet vote:" + c.TweeTVoteContractAddress
	s += "\ngame:" + c.GameContract
	s += "\nkol key:" + c.KolKeyContractAddress
	s += "\ninfura url:" + c.InfuraUrl
	s += "\ngame check time(minutes):" + fmt.Sprintf("%d", c.GameTimeInMinute)
	s += "\ntransaction check time(seconds):" + fmt.Sprintf("%d", c.TxCheckerInSeconds)
	s += "\nchain id:" + fmt.Sprintf("%d", c.ChainID)
	s += "\n--------------------------"
	return s
}

type DaemonProc struct {
	checkTicker *time.Ticker
}

func InitConf(cf *BCConf) {
	__conf = cf
	if __conf.CheckTimeInSecond == 0 {
		__conf.CheckTimeInSecond = 2
	}
}

var __conf *BCConf

func NewDaemon() *DaemonProc {
	var dp = &DaemonProc{
		checkTicker: time.NewTicker(time.Duration(__conf.CheckTimeInSecond) * time.Second),
	}
	return dp
}

func (dp *DaemonProc) Monitor() {
	for {
		select {
		case <-dp.checkTicker.C:
			util.LogInst().Debug().Msg("time to check block chain data")
			go database.DbInst().CheckKolElder()
		}
	}
}
