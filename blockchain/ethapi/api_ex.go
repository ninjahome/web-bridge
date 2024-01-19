package ethapi

import (
	"encoding/hex"
	"encoding/json"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
	"strings"
)

type GamInfoOnChain struct {
	RoundNo      *big.Int `json:"_"  firestore:"_"`
	RandomHash   string   `json:"random_hash"  firestore:"random_hash"`
	DiscoverTime int64    `json:"discover_time"  firestore:"discover_time"`
	Winner       string   `json:"winner"  firestore:"winner"`
	WinTeam      string   `json:"win_team"  firestore:"win_team"`
	WinTicketID  int64    `json:"win_ticket_id"  firestore:"win_ticket_id"`
	Bonus        float64  `json:"bonus"  firestore:"bonus"`
	RandomVal    string   `json:"random_val"  firestore:"random_val"`
}

func (c *GamInfoOnChain) String() string {
	bts, _ := json.Marshal(c)
	return string(bts)
}

func (_TweetLotteryGame *TweetLotteryGameCaller) GameInfoRecordEx(opts *bind.CallOpts, arg0 *big.Int) (*GamInfoOnChain, error) {
	var out []interface{}
	err := _TweetLotteryGame.contract.Call(opts, &out, "gameInfoRecord", arg0)

	construct := new(GamInfoOnChain)
	if err != nil {
		return construct, err
	}

	construct.RandomHash = hex.EncodeToString((*abi.ConvertType(out[0], new([32]byte)).(*[32]byte))[:])
	construct.DiscoverTime = (*abi.ConvertType(out[1], new(*big.Int)).(**big.Int)).Int64()
	construct.Winner = (*abi.ConvertType(out[2], new(common.Address)).(*common.Address)).Hex()
	construct.Winner = strings.ToLower(construct.Winner)

	construct.WinTeam = hex.EncodeToString((*abi.ConvertType(out[3], new([32]byte)).(*[32]byte))[:])
	construct.WinTicketID = (*abi.ConvertType(out[4], new(*big.Int)).(**big.Int)).Int64()
	construct.RandomVal = (*abi.ConvertType(out[6], new(*big.Int)).(**big.Int)).String()

	bonusBigInt := *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	weiToEth := new(big.Float).SetInt(big.NewInt(1e18))
	bonusEth := new(big.Float).Quo(new(big.Float).SetInt(bonusBigInt), weiToEth)
	construct.Bonus, _ = bonusEth.Float64()

	return construct, err
}

func (_TweetLotteryGame *TweetLotteryGameCaller) HistoryRoundInfoEx(opts *bind.CallOpts, from, to *big.Int) ([]*GamInfoOnChain, error) {

	var out []interface{}
	err := _TweetLotteryGame.contract.Call(opts, &out, "historyRoundInfo", from, to)

	if err != nil {
		return nil, err
	}

	out0 := *abi.ConvertType(out[0], new([]TweetLotteryGameGameInfoOneRound)).(*[]TweetLotteryGameGameInfoOneRound)
	if len(out0) == 0 {
		return nil, nil
	}
	var result = make([]*GamInfoOnChain, 0)
	for _, round := range out0 {

		var construct = new(GamInfoOnChain)
		construct.RandomHash = hex.EncodeToString(round.RandomHash[:])
		construct.DiscoverTime = round.DiscoverTime.Int64()
		construct.Winner = strings.ToLower(round.Winner.Hex())
		construct.WinTeam = hex.EncodeToString(round.WinTeam[:])
		construct.WinTicketID = round.WinTicketID.Int64()
		construct.RandomVal = round.RandomVal.String()

		weiToEth := new(big.Float).SetInt(big.NewInt(1e18))
		bonusEth := new(big.Float).Quo(new(big.Float).SetInt(round.Bonus), weiToEth)
		construct.Bonus, _ = bonusEth.Float64()

		result = append(result, construct)
	}

	return result, nil
}
