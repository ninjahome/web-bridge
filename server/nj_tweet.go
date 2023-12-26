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
