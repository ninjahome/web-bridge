package server

import (
	"encoding/json"
	"github.com/ninjahome/web-bridge/util"
	"net/http"
)

type SignInObj struct {
	EthAddr string `json:"eth_addr"`
	SignTim int64  `json:"sign_time"`
}

func (so *SignInObj) String() string {
	bts, _ := json.Marshal(so)
	return string(bts)
}

func MustSignInData(str string) *SignInObj {
	so := &SignInObj{}
	_ = json.Unmarshal([]byte(str), so)
	return so
}

type SignByEthParam struct {
	Message   string `json:"message"`
	Signature string `json:"signature"`
}

func signInByEth(w http.ResponseWriter, r *http.Request) {

	param := &SignByEthParam{}
	err := util.ReadRequest(r, param)
	if err != nil {
		util.LogInst().Err(err).Msg("sign in by eth address failed")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	obj := &SignInObj{}
	err = json.Unmarshal([]byte(param.Message), obj)
	if err != nil {
		util.LogInst().Err(err).Msg("parse sign in obj failed")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = util.Verify(obj.EthAddr, param.Message, param.Signature)
	if err != nil {
		util.LogInst().Err(err).Msg("verify signature failed")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	util.LogInst().Info().Str("eth-addr", obj.EthAddr).Int64("sign-time", obj.SignTim).Msg("sign in success")
	//nu := DbInst().NjUserSignIn(obj.EthAddr)
	nu := &NinjaUsrInfo{EthAddr: obj.EthAddr}
	if nu == nil {
		util.LogInst().Warn().Str("eth-addr", obj.EthAddr).Msgf("no user found")
		http.Error(w, "database error", http.StatusNotFound)
		return
	}

	//err = SMInst().Set(r, w, sesKeyForNJUserSignIN, obj.String())
	//if err != nil {
	//	util.LogInst().Err(err).Msgf("save twitter info failed:%v", nu)
	//	http.Error(w, err.Error(), http.StatusInternalServerError)
	//	return
	//}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(nu.RawData())
}

//func showMainPage(w http.ResponseWriter, r *http.Request) {
//	resultStr, err := SMInst().Get(sesKeyForNJUserSignIN, r)
//	if err != nil {
//		util.LogInst().Err(err).Msg("no user info found")
//		http.Error(w, err.Error(), http.StatusInternalServerError)
//		return
//	}
//	if resultStr == nil {
//		util.LogInst().Warn().Msg("direct visit without session")
//		http.Redirect(w, r, "/", http.StatusFound)
//		return
//	}
//
//	obj := MustSignInData(resultStr.(string))
//	err = htmlTemplateManager.ExecuteTemplate(w, "main.html", obj)
//	if err != nil {
//		util.LogInst().Err(err).Msg("show main page failed")
//		http.Error(w, err.Error(), http.StatusInternalServerError)
//		return
//	}
//}

//func logoutHandler(w http.ResponseWriter, r *http.Request) {
//	session, _ := store.Get(r, "session-name")
//	// 删除会话数据
//	session.Values["user"] = nil
//	session.Save(r, w)
//	// 重定向到登录页或主页
//	http.Redirect(w, r, "/login", http.StatusFound)
//}