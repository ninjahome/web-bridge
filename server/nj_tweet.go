package server

import (
	"encoding/json"
	"github.com/ninjahome/web-bridge/util"
	"net/http"
	"strconv"
)

func globalTweetQuery(w http.ResponseWriter, r *http.Request, nu *NinjaUsrInfo) {
	var latestIDStr = r.URL.Query().Get("lastTwID")
	latestID, err := strconv.ParseInt(latestIDStr, 10, 64)
	if err != nil {
		util.LogInst().Err(err).Str("latest-id", latestIDStr).Msg("invalid latest tweet id")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	var tweets = make([]*NinjaTweet, 0)
	err = DbInst().QueryGlobalLatestTweets(_globalCfg.TweetsPageSize, latestID, func(tweet *NinjaTweet) {
		tweets = append(tweets, tweet)
	})
	if err != nil {
		util.LogInst().Err(err).Str("latest-id", latestIDStr).Msg("query global tweets failed")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	bts, _ := json.Marshal(tweets)
	w.Write(bts)
	util.LogInst().Debug().Str("latest-id", latestIDStr).Int("size", len(tweets)).Msg("global tweets query success")
}
