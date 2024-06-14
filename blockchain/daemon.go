package blockchain

import (
	"fmt"
	"github.com/ninjahome/web-bridge/database"
	"github.com/ninjahome/web-bridge/util"
	"sync"
	"time"
)

const (
	DefaultElderChecker  = 2
	DefaultBonusChecker  = 5
	DefaultBonusInterval = 8 * 60
)

type BCConf struct {
	TweetContract           string `json:"tweet_vote_contract_address"`
	GameContract            string `json:"game_plugin_contract_address"`
	KolKeyContractAddress   string `json:"kol_key_contract_address"`
	InfuraUrl               string `json:"infura_url"`
	GameTimeInMinute        int    `json:"game_time_in_minute,omitempty"`
	TxCheckerInSeconds      int    `json:"tx_checker_in_seconds,omitempty"`
	ChainID                 int64  `json:"chain_id,omitempty"`
	ElderCheckTimeInSec     int    `json:"elder_check_time_in_sec"`
	PointBonusCheckInMin    int    `json:"point_bonus_check_in_min"`
	PointBonusIntervalInMin int    `json:"point_bonus_interval_in_min"`
}

func (c *BCConf) String() string {
	s := "\n------block chain config------"

	s += "\ntweet vote:" + c.TweetContract
	s += "\ngame:" + c.GameContract
	s += "\nkol key:" + c.KolKeyContractAddress
	s += "\ninfura url:" + c.InfuraUrl
	s += "\ngame check time(minutes):" + fmt.Sprintf("%d", c.GameTimeInMinute)
	s += "\ntransaction check time(seconds):" + fmt.Sprintf("%d", c.TxCheckerInSeconds)
	s += "\nchain id:" + fmt.Sprintf("%d", c.ChainID)
	s += "\nelder check time(seconds):" + fmt.Sprintf("%d", c.ElderCheckTimeInSec)
	s += "\npoint bonus check time(minutes):" + fmt.Sprintf("%d", c.PointBonusCheckInMin)
	s += "\npoint bonus interval time(minutes):" + fmt.Sprintf("%d", c.PointBonusIntervalInMin)

	s += "\n--------------------------"
	return s
}

type DaemonProc struct {
	elderCheck       *time.Ticker
	pointsBonusCheck *time.Ticker
	nextBonusTime    time.Time
	pointSumSnapshot float64
}

func InitConf(cf *BCConf) {
	__conf = cf
	if __conf.ElderCheckTimeInSec == 0 {
		__conf.ElderCheckTimeInSec = DefaultElderChecker
	}
	if __conf.PointBonusCheckInMin == 0 {
		__conf.PointBonusCheckInMin = DefaultBonusChecker
	}
	if __conf.PointBonusIntervalInMin == 0 {
		__conf.PointBonusIntervalInMin = DefaultBonusInterval
	}
}

var __conf *BCConf

func newDaemon() *DaemonProc {
	var dp = &DaemonProc{
		elderCheck:       time.NewTicker(time.Duration(__conf.ElderCheckTimeInSec) * time.Second),
		pointsBonusCheck: time.NewTicker(time.Duration(__conf.PointBonusCheckInMin) * time.Minute),
		nextBonusTime:    time.Now().Add(time.Duration(__conf.PointBonusIntervalInMin) * time.Minute),
	}
	return dp
}

var _dSyncOnce sync.Once
var _dInst *DaemonProc

func DaemonInst() *DaemonProc {
	_dSyncOnce.Do(func() {
		_dInst = newDaemon()
	})
	return _dInst
}

func (dp *DaemonProc) Monitor() {

	dp.pointSumSnapshot = database.DbInst().PointsAtSnapshot()
	for {
		select {
		case <-dp.elderCheck.C:
			go database.DbInst().CheckKolElder()
			break
		case <-dp.pointsBonusCheck.C:
			go dp.checkPointBonus()
			break
		}
	}
}

func (dp *DaemonProc) checkPointBonus() {
	util.LogInst().Debug().Msg("start to check point bonus")
	now := time.Now()
	if now.Before(dp.nextBonusTime) {
		util.LogInst().Debug().Msg("Points reward time has not arrived")
		return
	}

	newTotal := database.DbInst().RewardForOneRound(dp.pointSumSnapshot)
	if newTotal <= 0 {
		util.LogInst().Error().Msg("calculate new points sum failed")
		return
	}
	dp.pointSumSnapshot = newTotal
}

func (dp *DaemonProc) PointSumAtCurrentRound() float64 {
	return dp.pointSumSnapshot
}
