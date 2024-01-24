package server

import (
	"encoding/json"
	database2 "github.com/ninjahome/web-bridge/database"
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

func signInByEth(w http.ResponseWriter, r *http.Request, _ *database2.NinjaUsrInfo) {
	param := &SignDataByEth{}
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
	_, err = util.Verify(obj.EthAddr, param.Message, param.Signature)
	if err != nil {
		util.LogInst().Err(err).Msg("sign in verify signature failed")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	util.LogInst().Info().Str("eth-addr", obj.EthAddr).Int64("sign-time", obj.SignTim).Msg("sign in success")
	nu := database2.DbInst().NjUserSignIn(obj.EthAddr)
	if nu == nil {
		util.LogInst().Warn().Str("eth-addr", obj.EthAddr).Msgf("no user found")
		http.Error(w, "database error", http.StatusNotFound)
		return
	}

	err = SMInst().Set(r, w, sesKeyForRightCheck, nu.RawData())
	if err != nil {
		util.LogInst().Err(err).Msg("save sign in param to session failed")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	nu.RefreshSession()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(nu.RawData())
	util.LogInst().Debug().Str("eth-addr", obj.EthAddr).Msg("sign in by eth success")
}
