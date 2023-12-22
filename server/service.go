package server

import (
	"fmt"
	"github.com/ninjahome/web-bridge/util"
	"net/http"
)

const (
	sesKeyForRightCheck = "session-key-right-checking"
)

func queryTwBasicById(w http.ResponseWriter, r *http.Request) {
	var ninjaUser, err = validateUsrRights(w, r)
	if err != nil {
		http.Redirect(w, r, "/signIn", http.StatusFound)
		return
	}
	var twitterID = ninjaUser.TwID
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
	var _, err = validateUsrRights(w, r)
	if err != nil {
		http.Redirect(w, r, "/signIn", http.StatusFound)
		return
	}
	err = htmlTemplateManager.ExecuteTemplate(w, "main.html", nil)
	if err != nil {
		http.Redirect(w, r, "/signIn", http.StatusFound)
		return
	}
}

func signOut(w http.ResponseWriter, r *http.Request) {
	_ = SMInst().Del(sesKeyForRightCheck, r, w)
	http.Redirect(w, r, "/signIn", http.StatusFound)
}

func validateUsrRights(w http.ResponseWriter, r *http.Request) (*NinjaUsrInfo, error) {
	var data, err = SMInst().Get(sesKeyForRightCheck, r)
	if err != nil {
		util.LogInst().Warn().Msgf("%s", err.Error())
		return nil, err
	}

	var njUser, errNu = NJUsrInfoMust(data.([]byte))
	if errNu != nil {
		return nil, fmt.Errorf("not a ninja user struct saved")
	}
	return njUser, nil
}
