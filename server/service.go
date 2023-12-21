package server

import (
	"github.com/ninjahome/web-bridge/util"
	"net/http"
)

func queryTwBasicById(w http.ResponseWriter, r *http.Request) {
	var twitterID = ""
	var err = util.ReadRequest(r, &twitterID)
	if err != nil {
		util.LogInst().Err(err).Msg("twitter id not in param for query basic info")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if len(twitterID) == 0 {
		util.LogInst().Warn().Msg("invalid twitter id param")
		http.Error(w, "twitter id invalid", http.StatusBadRequest)
		return
	}
	var userdata *TWUserInfo = nil
	userdata, err = DbInst().TwitterBasicInfo(twitterID)

	if err != nil {
		util.LogInst().Err(err).Msg("query twitter data failed")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(userdata.RawData())
}
