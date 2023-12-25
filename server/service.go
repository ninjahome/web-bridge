package server

import (
	"encoding/json"
	"fmt"
	"github.com/ninjahome/web-bridge/util"
	"html/template"
	"net/http"
)

const (
	sesKeyForRightCheck = "session-key-right-checking"
	BuyRightsUrlKey     = "twOwner"
)

type SignDataByEth struct {
	Message   string `json:"message"`
	Signature string `json:"signature"`
	PayLoad   any    `json:"pay_load,omitempty"`
}
type Web3BindingData struct {
	EthAddr  string `json:"eth_addr"`
	TwID     string `json:"tw_id"`
	BindTime int64  `json:"bind_time"`
}

func (sp *SignDataByEth) RawData() string {
	bts, _ := json.Marshal(sp)
	return string(bts)
}

func (sp *SignDataByEth) ParseNinjaTweet() (*NinjaTweet, error) {
	var tweetContent NinjaTweet
	var err = json.Unmarshal([]byte(sp.Message), &tweetContent)
	if err != nil {
		util.LogInst().Err(err).Msg("Error parsing tweet ")
		return nil, err
	}

	if !tweetContent.IsValid() {
		util.LogInst().Warn().Msg("invalid tweet content:" + tweetContent.String())
		return nil, fmt.Errorf("invalid tweet content")
	}

	err = util.Verify(tweetContent.Web3ID, sp.Message, sp.Signature)
	if err != nil {
		util.LogInst().Err(err).Msg("tweet signature verify failed")
		return nil, err
	}
	tweetContent.Signature = sp.Signature

	return &tweetContent, nil
}

func queryTwBasicById(w http.ResponseWriter, _ *http.Request, ninjaUser *NinjaUsrInfo) {

	var twitterID = ninjaUser.TwID
	if len(twitterID) == 0 {
		util.LogInst().Warn().Msg("invalid twitter id param")
		http.Error(w, "twitter id invalid", http.StatusBadRequest)
		return
	}
	var userdata, err = DbInst().TwitterBasicInfo(twitterID)
	if err != nil {
		util.LogInst().Err(err).Msg("query twitter data failed")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(userdata.RawData())
}

func mainPage(w http.ResponseWriter, r *http.Request, nu *NinjaUsrInfo) {

	data := struct {
		NinjaUsrInfoJson template.JS
	}{
		NinjaUsrInfoJson: template.JS(nu.RawData()),
	}
	var err = _globalCfg.htmlTemplateManager.ExecuteTemplate(w, "main.html", data)
	if err != nil {
		util.LogInst().Err(err).Msg("main html failed")
		http.Redirect(w, r, "/signIn", http.StatusFound)
		return
	}
}

func signOut(w http.ResponseWriter, r *http.Request, _ *NinjaUsrInfo) {
	_ = SMInst().Del(sesKeyForRightCheck, r, w)
	http.Redirect(w, r, "/signIn", http.StatusFound)
}

func validateUsrRights(r *http.Request) *NinjaUsrInfo {
	var data, err = SMInst().Get(sesKeyForRightCheck, r)
	if err != nil {
		util.LogInst().Warn().Msgf("%s", err.Error())
		return nil
	}

	var njUser, errNu = NJUsrInfoMust(data.([]byte))
	if errNu != nil {
		util.LogInst().Warn().Msgf("ninja user not found")
		return nil
	}
	return njUser
}

func bindingWeb3ID(w http.ResponseWriter, r *http.Request, origNu *NinjaUsrInfo) {
	param := &SignDataByEth{}
	err := util.ReadRequest(r, param)
	if err != nil {
		util.LogInst().Err(err).Msg("no sign data by eth found")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if param.PayLoad == nil {
		util.LogInst().Err(err).Msg("lost payload in sign data")
		http.Error(w, "lost payload in sign data", http.StatusBadRequest)
		return
	}
	twUsrData := &TWUserInfo{}
	err = json.Unmarshal([]byte(param.PayLoad.(string)), twUsrData)
	if err != nil {
		util.LogInst().Err(err).Msg("parse twitter data failed")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	data := &Web3BindingData{}
	err = json.Unmarshal([]byte(param.Message), data)
	if err != nil {
		util.LogInst().Err(err).Msg("parse sign data failed")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if data.EthAddr != origNu.EthAddr {
		util.LogInst().Warn().Msgf("metamask account has changed[%s=>%s]",
			origNu.EthAddr, data.EthAddr)
		http.Redirect(w, r, "/signIn", http.StatusFound)
		return
	}
	err = util.Verify(data.EthAddr, param.Message, param.Signature)
	if err != nil {
		util.LogInst().Err(err).Msg("binding data verify signature failed")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	util.LogInst().Info().Str("eth-addr", data.EthAddr).
		Str("twitter-id", data.TwID).
		Int64("bind-time", data.BindTime).Msg("sign data success")

	bindDataToStore := &Web3Binding{
		TwitterID: data.TwID,
		EthAddr:   data.EthAddr,
		SignUpAt:  data.BindTime,
		Signature: param.Signature,
	}
	newNu, err := DbInst().BindingWeb3ID(bindDataToStore, twUsrData)
	if err != nil {
		util.LogInst().Err(err).Msg("save binding data  failed")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = SMInst().Set(r, w, sesKeyForRightCheck, newNu.RawData())
	if err != nil {
		util.LogInst().Err(err).Msg("setup new ninja user data  failed")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_ = updateTwitterBio(r, twUsrData.Description, data.EthAddr)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(newNu.RawData())
}
