package server

import (
	"github.com/ninjahome/web-bridge/util"
	"net/http"
)

const (
	sesKeyForSignInParam = "session-key-for-sign-in-param"
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

func mainPage(w http.ResponseWriter, r *http.Request) {
	var _, err = SMInst().Get(sesKeyForSignInParam, r)
	if err != nil {
		http.Redirect(w, r, "/signIn", http.StatusFound)
		util.LogInst().Warn().Msgf("%s", err.Error())
		return
	}
	err = htmlTemplateManager.ExecuteTemplate(w, "main.html", nil)
	if err != nil {
		http.Redirect(w, r, "/signIn", http.StatusFound)
		return
	}
}

func signOut(w http.ResponseWriter, r *http.Request) {
	_ = SMInst().Del(sesKeyForSignInParam, r, w)
	http.Redirect(w, r, "/signIn", http.StatusFound)
}
