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
	Status     database.TxStatus `json:"status"`
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

type TweetVoteAction struct {
	CreateTime int64 `json:"create_time"`
	VoteCount  int   `json:"vote_count"`
}

func updateTweetVoteStatus(w http.ResponseWriter, r *http.Request, nu *database.NinjaUsrInfo) {
	vote := &TweetVoteAction{}
	var err = util.ReadRequest(r, vote)
	if err != nil {
		util.LogInst().Err(err).Msg("parsing payment status param failed ")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if vote.CreateTime == 0 {
		util.LogInst().Warn().Int64("create_time", vote.CreateTime).Msg("invalid tweet create time")
		http.Error(w, "invalid tweet create time", http.StatusBadRequest)
		return
	}

	newVal, err := database.DbInst().UpdateTweetVoteStatic(vote.CreateTime, vote.VoteCount)
	if err != nil {
		util.LogInst().Err(err).Int64("create_time", vote.CreateTime).
			Int("vote_count", vote.VoteCount).
			Msg("failed to update tweet vote ")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	vote.VoteCount = newVal
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	bts, _ := json.Marshal(vote)
	w.Write(bts)

	util.LogInst().Debug().Int64("create_time", vote.CreateTime).
		Int("vote_count", vote.VoteCount).
		Msg(" update vote count of tweet success")
}
