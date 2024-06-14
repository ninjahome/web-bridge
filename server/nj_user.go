package server

import (
	"encoding/json"
	"github.com/ninjahome/web-bridge/blockchain"
	"github.com/ninjahome/web-bridge/database"
	"github.com/ninjahome/web-bridge/util"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net/http"
)

func queryNjBasicByReferrer(w http.ResponseWriter, r *http.Request, nu *database.NinjaUsrInfo) {
	var refCode = r.URL.Query().Get("referrer_code")
	if len(refCode) == 0 {
		util.LogInst().Warn().Str("referrer_code", refCode).
			Str("eth-addr", nu.EthAddr).Msg("invalid web3 id param")
		http.Error(w, "web3 id invalid", http.StatusBadRequest)
		return
	}
	obj, err := database.DbInst().QueryNjUsrByReferrer(refCode)
	if err != nil && status.Code(err) != codes.NotFound {
		util.LogInst().Err(err).Msg("query ninja user data  failed")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if obj != nil {
		json.NewEncoder(w).Encode(obj)
	}
}

func queryNjBasicByID(w http.ResponseWriter, r *http.Request, _ *database.NinjaUsrInfo) {
	var web3ID = r.URL.Query().Get("web3_id")
	if len(web3ID) == 0 {
		util.LogInst().Warn().Str("web3-id", web3ID).Msg("invalid web3 id param")
		http.Error(w, "web3 id invalid", http.StatusBadRequest)
		return
	}

	obj, err := database.DbInst().QueryNjUsrById(web3ID)
	if err != nil && status.Code(err) != codes.NotFound {
		util.LogInst().Err(err).Msg("query ninja user data  failed")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if obj != nil {
		json.NewEncoder(w).Encode(obj)
	}
}

func mostVotedKol(w http.ResponseWriter, r *http.Request, nu *database.NinjaUsrInfo) {
	var para database.TweetQueryParam
	var err = util.ReadRequest(r, &para)
	if err != nil {
		util.LogInst().Err(err).Str("param", para.String()).
			Msg("invalid query parameter")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	obj, err := database.DbInst().MostVotedKol(_globalCfg.TweetsPageSize, para.StartID, len(para.VotedIDs) > 0)

	if err != nil {
		util.LogInst().Err(err).Str("user-web3-id", nu.EthAddr).
			Msg("failed to query most voted kol")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	bts, _ := json.Marshal(obj)
	w.Write(bts)

	util.LogInst().Debug().Int("id-len", len(obj)).Str("param", para.String()).
		Msg(" query most voted kol success")
}

func pointsForNJUsr(w http.ResponseWriter, r *http.Request, _ *database.NinjaUsrInfo) {
	var web3ID = r.URL.Query().Get("web3_id")
	if len(web3ID) == 0 {
		util.LogInst().Warn().Str("web3-id", web3ID).Msg("invalid web3 id param")
		http.Error(w, "web3 id invalid", http.StatusBadRequest)
		return
	}

	obj, err := database.DbInst().QuerySystemPoints(web3ID)
	if err != nil && status.Code(err) != codes.NotFound {
		util.LogInst().Err(err).Msg("query ninja user data  failed")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	obj.CurTotalPoints = blockchain.DaemonInst().PointSumAtCurrentRound()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if obj != nil {
		json.NewEncoder(w).Encode(obj)
	}
}
