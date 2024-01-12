package server

import (
	"github.com/ninjahome/web-bridge/database"
	"github.com/ninjahome/web-bridge/util"
	"net/http"
)

func userProfile(w http.ResponseWriter, r *http.Request, web3ID string) {

}

func queryNjBasicByID(w http.ResponseWriter, r *http.Request, nu *database.NinjaUsrInfo) {
	var web3ID = r.URL.Query().Get("web3_id")
	if len(web3ID) == 0 {
		util.LogInst().Warn().Str("web3-id", web3ID).
			Str("eth-addr", nu.EthAddr).Msg("invalid web3 id param")
		http.Error(w, "web3 id invalid", http.StatusBadRequest)
		return
	}

	obj, err := database.DbInst().QueryNjUsrById(web3ID)
	if err != nil {
		util.LogInst().Err(err).Msg("query ninja user data  failed")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(obj.RawData())
}
