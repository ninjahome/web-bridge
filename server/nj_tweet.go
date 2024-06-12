package server

import (
	"context"
	"encoding/json"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ninjahome/web-bridge/blockchain/ethapi"
	"github.com/ninjahome/web-bridge/database"
	"github.com/ninjahome/web-bridge/util"
	"net/http"
	"strconv"
	"strings"
	"time"
)

const (
	MaxIntervalForPaymentStatus = 60
)

func globalTweetQuery(w http.ResponseWriter, r *http.Request, nu *database.NinjaUsrInfo) {

	var para database.TweetQueryParm
	var err = util.ReadRequest(r, &para)
	if err != nil {
		util.LogInst().Err(err).Str("param", para.String()).
			Msg("invalid query parameter")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	para.IsOwner = nu.EthAddr == para.Web3ID
	tweets, err := database.DbInst().QueryTweetsByFilter(_globalCfg.TweetsPageSize, &para)
	if err != nil {
		util.LogInst().Err(err).Str("param", para.String()).
			Str("eth-addr", nu.EthAddr).
			Msg("query global tweets failed")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	bts, _ := json.Marshal(tweets)
	w.Write(bts)

	util.LogInst().Info().Str("param", para.String()).
		Str("eth-addr", nu.EthAddr).
		Int("size", len(tweets)).Msg("global tweets query success")
}

func querySimplePaymentTransaction(tx string) bool {
	cli, err := ethclient.Dial(_globalCfg.InfuraUrl)
	if err != nil {
		util.LogInst().Err(err).Msg("dial eth failed")
		return false
	}

	defer cli.Close()
	txHash := common.HexToHash(tx)
	receipt, err := cli.TransactionReceipt(context.Background(), txHash)
	if err != nil {
		util.LogInst().Err(err).Str("tx-hash", tx).Msg("query receipt failed")
		return false
	}
	if receipt.Status != types.ReceiptStatusSuccessful {
		return false
	}

	block, err := util.GetBlockByNumber(_globalCfg.InfuraUrl, receipt.BlockNumber)
	if err != nil {
		util.LogInst().Err(err).Msg("query block failed")
		return false
	}
	util.LogInst().Debug().Int64("now", time.Now().Unix()).Int64("block-time", block.TimeStamp2.Int64()).Msg("check payment status")
	return time.Now().Unix() < block.TimeStamp2.Int64()+MaxIntervalForPaymentStatus
}

func getContractObj() (*ethapi.TweetVote, error) {
	cli, err := ethclient.Dial(_globalCfg.InfuraUrl)
	if err != nil {
		util.LogInst().Err(err).Msg("dial eth failed")
		return nil, err
	}

	defer cli.Close()

	contractAddress := common.HexToAddress(_globalCfg.TweetContract)
	tweetContract, err := ethapi.NewTweetVote(contractAddress, cli)
	if err != nil {
		util.LogInst().Err(err).Str("contract-address", _globalCfg.GameContract).Msg("failed create tweet contract obj")
		return nil, err
	}

	return tweetContract, nil
}

func queryTweetPaymentStatusFromBlockChain(twID int64, realOwner string) bool {
	tweet, err := database.DbInst().NjTweetDetails(twID)
	if err != nil {
		util.LogInst().Err(err).Int64("tweet-id", twID).Msg("check tweet payment status failed")
		return false
	}
	if tweet.PaymentStatus != database.TxStNotPay {
		util.LogInst().Warn().Int64("tweet-id", twID).Msg("duplicate update for tweet status")
		return false
	}

	if len(tweet.PrefixedHash) < 64 {
		util.LogInst().Warn().Int64("tweet-id", twID).Msg("no prefix hash for this tweet")
		return false
	}

	tweetObj, err := getContractObj()
	if err != nil {
		util.LogInst().Err(err).Int64("tweet-id", twID).Msg("create tweet contract obj failed")
		return false
	}

	owner, err := tweetObj.OwnersOfAllTweets(nil, common.HexToHash(tweet.PrefixedHash))
	if err != nil {
		util.LogInst().Err(err).Msg("query tweet owner from contract failed")
	}

	return strings.ToLower(owner.String()) == realOwner
}

func checkStatus(status *database.TweetPaymentStatus, tweetOwner string) bool {
	if status.Status != database.TxStSuccess {
		return true
	}

	if status.CreateTime == 0 {
		util.LogInst().Warn().Int64("create_time", status.CreateTime).Msg("invalid tweet create time")
		return false
	}
	if len(status.TxHash) < 16 {
		return queryTweetPaymentStatusFromBlockChain(status.CreateTime, tweetOwner)
	}
	return querySimplePaymentTransaction(status.TxHash)
}

func updateTweetTxStatus(w http.ResponseWriter, r *http.Request, nu *database.NinjaUsrInfo) {
	status := &database.TweetPaymentStatus{}
	var err = util.ReadRequest(r, status)
	if err != nil {
		util.LogInst().Err(err).Msg("parsing payment status param failed ")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if checkStatus(status, strings.ToLower(nu.EthAddr)) == false {
		http.Error(w, "invalid status update", http.StatusBadRequest)
		return
	}

	err = database.DbInst().UpdateTweetPaymentStatus(status, nu.EthAddr)
	if err != nil {
		util.LogInst().Err(err).Int64("create_time", status.CreateTime).
			Str("status", status.Status.String()).
			Msg("failed to update tweet payment status")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	bts, _ := json.Marshal(status)
	w.Write(bts)

	util.LogInst().Info().Int64("create_time", status.CreateTime).
		Str("status", status.Status.String()).
		Str("web3-id", nu.EthAddr).
		Msg(" update status of tweet payment success")
}

func queryTweetDetails(w http.ResponseWriter, r *http.Request, _ *database.NinjaUsrInfo) {
	var createTimeStr = r.URL.Query().Get("tweetID")
	var createTime, _ = strconv.ParseInt(createTimeStr, 10, 64)
	obj, err := database.DbInst().NjTweetDetails(createTime)
	if err != nil {
		util.LogInst().Err(err).Int64("id", createTime).Msg("query tweet detail failed")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	bts, _ := json.Marshal(obj)
	w.Write(bts)
	util.LogInst().Debug().Int64("id", createTime).Msg("query tweet detail success")
}

func updatePointsForSingleBets(w http.ResponseWriter, r *http.Request, nu *database.NinjaUsrInfo) {
	vote := &database.TweetVoteAction{}
	var err = util.ReadRequest(r, vote)
	if err != nil {
		util.LogInst().Err(err).Msg("parsing vote param failed ")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if querySimplePaymentTransaction(vote.TxHash) == false {
		util.LogInst().Warn().Int64("create_time", vote.CreateTime).Msg("payment status invalid")
		http.Error(w, "payment status invalid", http.StatusBadRequest)
		return
	}

	err = database.DbInst().UpdatePointsForSingleBets(vote, nu.EthAddr)
	if err != nil {
		util.LogInst().Err(err).Int64("create_time", vote.CreateTime).
			Int("vote_count", vote.VoteCount).
			Msg("failed to update points for single bets")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	util.LogInst().Info().Int64("create_time", vote.CreateTime).
		Int("vote_count", vote.VoteCount).
		Str("tx-hash", vote.TxHash).
		Str("web3-id", nu.EthAddr).
		Msg(" update points for single vote success")
}

func updateTweetVoteStatus(w http.ResponseWriter, r *http.Request, nu *database.NinjaUsrInfo) {
	vote := &database.TweetVoteAction{}
	var err = util.ReadRequest(r, vote)
	if err != nil {
		util.LogInst().Err(err).Msg("parsing vote param failed ")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if querySimplePaymentTransaction(vote.TxHash) == false {
		util.LogInst().Warn().Int64("create_time", vote.CreateTime).Msg("payment status invalid")
		http.Error(w, "payment status invalid", http.StatusBadRequest)
		return
	}

	err = database.DbInst().UpdateTweetVoteStatic(vote, nu.EthAddr)
	if err != nil {
		util.LogInst().Err(err).Int64("create_time", vote.CreateTime).
			Int("vote_count", vote.VoteCount).
			Msg("failed to update tweet vote ")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	bts, _ := json.Marshal(vote)
	w.Write(bts)

	util.LogInst().Info().Int64("create_time", vote.CreateTime).
		Int("vote_count", vote.VoteCount).
		Str("web3-id", nu.EthAddr).
		Msg(" update vote count of tweet success")
}

func votedTweetsQuery(w http.ResponseWriter, r *http.Request, nu *database.NinjaUsrInfo) {

	var para database.TweetQueryParm
	var err = util.ReadRequest(r, &para)
	if err != nil {
		util.LogInst().Err(err).Str("param", para.String()).
			Msg("invalid query parameter")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	ids, err := database.DbInst().QueryVotedTweetIDByMe(_globalCfg.TweetsPageSize, para.StartID, para.Web3ID)
	if err != nil {
		util.LogInst().Err(err).Str("user-web3-id", nu.EthAddr).
			Msg("failed to query voted tweets ")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	bts, _ := json.Marshal(ids)
	w.Write(bts)

	util.LogInst().Debug().Int("id-len", len(ids)).Str("param", para.String()).
		Msg(" query voted  tweet success")
}

func removeUnpaidTweet(w http.ResponseWriter, r *http.Request, nu *database.NinjaUsrInfo) {

	var status database.TweetPaymentStatus
	var err = util.ReadRequest(r, &status)

	if err != nil {
		util.LogInst().Err(err).Msg("parsing param failed when delete tweet")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = database.DbInst().DelUnpaidTweet(status.CreateTime, nu.EthAddr)
	if err != nil {
		util.LogInst().Err(err).Int64("create_time", status.CreateTime).
			Str("web3-id", nu.EthAddr).Msg("failed to delete unpaid tweet")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	util.LogInst().Info().Int64("create_time", status.CreateTime).
		Str("web3-id", nu.EthAddr).Msg(" delete unpaid tweet success")

}

func mostVotedTweet(w http.ResponseWriter, r *http.Request, nu *database.NinjaUsrInfo) {
	var para database.TweetQueryParm
	var err = util.ReadRequest(r, &para)
	if err != nil {
		util.LogInst().Err(err).Str("param", para.String()).
			Msg("invalid query parameter")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	tweets, err := database.DbInst().QueryMostVotedTweets(_globalCfg.TweetsPageSize, para.StartID)
	if err != nil {
		util.LogInst().Err(err).Str("param", para.String()).
			Str("eth-addr", nu.EthAddr).
			Msg("query most voted tweets failed")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	bts, _ := json.Marshal(tweets)
	w.Write(bts)

	util.LogInst().Debug().Str("param", para.String()).
		Str("eth-addr", nu.EthAddr).
		Int("size", len(tweets)).Msg("most voted tweets query success")
}

func tweetImgRaw(w http.ResponseWriter, r *http.Request, _ *database.NinjaUsrInfo) {
	var hash = r.URL.Query().Get("img_hash")
	obj, err := database.DbInst().GetRawImg(hash)
	if err != nil {
		util.LogInst().Err(err).Str("img-hash", hash).Msg("query tweet img raw img failed")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	bts, _ := json.Marshal(obj)
	w.Write(bts)
	util.LogInst().Debug().Str("img-hash", hash).Msg("query tweet img raw success")
}

func tweetImgThumb(w http.ResponseWriter, r *http.Request, _ *database.NinjaUsrInfo) {
	var hash = r.URL.Query().Get("img_hash")
	obj, err := database.DbInst().GetThumbImg(hash)
	if err != nil {
		util.LogInst().Err(err).Str("img-hash", hash).Msg("query tweet img raw img failed")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	bts, _ := json.Marshal(obj)
	w.Write(bts)
	util.LogInst().Debug().Str("img-hash", hash).Msg("query tweet img thumb success")
}

func queryTweetByHash(w http.ResponseWriter, r *http.Request, _ *database.NinjaUsrInfo) {
	var tweetHash = r.URL.Query().Get("tweet_hash")

	obj, err := database.DbInst().NjTweetDetailsByHash(tweetHash)
	if err != nil {
		util.LogInst().Err(err).Str("tweet-hash", tweetHash).Msg("query tweet detail failed")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	bts, _ := json.Marshal(obj)
	w.Write(bts)
	util.LogInst().Debug().Str("tweet-hash", tweetHash).Msg("query tweet by hash success")
}
