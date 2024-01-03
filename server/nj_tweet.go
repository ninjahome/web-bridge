package server

import (
	"encoding/json"
	"github.com/ninjahome/web-bridge/util"
	"net/http"
	"strconv"
)

func globalTweetQuery(w http.ResponseWriter, r *http.Request, nu *NinjaUsrInfo) {
	var startIDStr = r.URL.Query().Get("startID")
	var newestStr = r.URL.Query().Get("isRefresh")
	startID, err1 := strconv.ParseInt(startIDStr, 10, 64)

	newest, err2 := strconv.ParseBool(newestStr)
	if err1 != nil || err2 != nil {
		util.LogInst().Err(err1).Err(err2).Str("latest-id", startIDStr).
			Str("eth-addr", nu.EthAddr).Str("newest", newestStr).
			Msg("invalid query parameter")
		http.Error(w, err1.Error(), http.StatusBadRequest)
		return
	}
	var tweets = make([]*NinjaTweet, 0)
	var err = DbInst().QueryGlobalLatestTweets(_globalCfg.TweetsPageSize, startID, newest, func(tweet *NinjaTweet) {
		tweets = append(tweets, tweet)
	})
	if err != nil {
		util.LogInst().Err(err).Str("latest-id", startIDStr).
			Str("eth-addr", nu.EthAddr).Msg("query global tweets failed")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	bts, _ := json.Marshal(tweets)
	w.Write(bts)

	util.LogInst().Debug().Str("latest-id", startIDStr).
		Str("newest", newestStr).
		Str("eth-addr", nu.EthAddr).
		Int("size", len(tweets)).Msg("global tweets query success")
}

type TweetPaymentStatus struct {
	CreateTime int64    `json:"create_time"`
	Status     TxStatus `json:"status"`
	Hash       string   `json:"hash"`
}

func updateTweetTxStatus(w http.ResponseWriter, r *http.Request, _ *NinjaUsrInfo) {
	status := &TweetPaymentStatus{}
	var err = util.ReadRequest(r, status)
	if err != nil {
		util.LogInst().Err(err).Msg("parsing payment status param failed ")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if status.CreateTime == 0 {
		util.LogInst().Warn().Str("hash", status.Hash).Msg("invalid tweet create time")
		http.Error(w, "invalid tweet create time", http.StatusBadRequest)
		return
	}

	err = DbInst().UpdateTweetPaymentStatus(status.CreateTime, status.Status, status.Hash)
	if err != nil {
		util.LogInst().Err(err).Int64("create_time", status.CreateTime).
			Str("status", status.Status.String()).Str("hash", status.Hash).
			Msg("failed to update tweet payment status")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	bts, _ := json.Marshal(status)
	w.Write(bts)

	util.LogInst().Debug().Int64("create_time", status.CreateTime).
		Str("status", status.Status.String()).Str("hash", status.Hash).
		Msg(" update status of tweet payment success")
}
