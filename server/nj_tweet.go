package server

import (
	"encoding/json"
	"github.com/ninjahome/web-bridge/database"
	"github.com/ninjahome/web-bridge/util"
	"net/http"
	"strconv"
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

	util.LogInst().Debug().Str("param", para.String()).
		Str("eth-addr", nu.EthAddr).
		Int("size", len(tweets)).Msg("global tweets query success")
}

type TweetPaymentStatus struct {
	CreateTime int64             `json:"create_time"`
	Status     database.TxStatus `json:"status,omitempty"`
}

func updateTweetTxStatus(w http.ResponseWriter, r *http.Request, _ *database.NinjaUsrInfo) {
	status := &TweetPaymentStatus{}
	var err = util.ReadRequest(r, status)
	if err != nil {
		util.LogInst().Err(err).Msg("parsing payment status param failed ")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if status.CreateTime == 0 {
		util.LogInst().Warn().Int64("create_time", status.CreateTime).Msg("invalid tweet create time")
		http.Error(w, "invalid tweet create time", http.StatusBadRequest)
		return
	}

	err = database.DbInst().UpdateTweetPaymentStatus(status.CreateTime, status.Status)
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

	util.LogInst().Debug().Int64("create_time", status.CreateTime).
		Str("status", status.Status.String()).
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

func updateTweetVoteStatus(w http.ResponseWriter, r *http.Request, nu *database.NinjaUsrInfo) {
	vote := &database.TweetVoteAction{}
	var err = util.ReadRequest(r, vote)
	if err != nil {
		util.LogInst().Err(err).Msg("parsing payment status param failed ")
		http.Error(w, err.Error(), http.StatusBadRequest)
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

	util.LogInst().Debug().Int64("create_time", vote.CreateTime).
		Int("vote_count", vote.VoteCount).
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

	var status TweetPaymentStatus
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
	w.Write([]byte("success"))

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
