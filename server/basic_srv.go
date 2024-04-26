package server

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/csrf"
	"github.com/ninjahome/web-bridge/database"
	"github.com/ninjahome/web-bridge/util"
	"html/template"
	"net/http"
	"strconv"
)

const (
	sesKeyForRightCheck = "session-key-right-checking"
	BuyRightsUrlKey     = "tweet-info-from-outer-link"
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

type TweetImgData struct {
	RawData   string `json:"raw_data"`
	Hash      string `json:"hash"`
	ThumbNail string `json:"thumb_nail"`
}

func (sp *SignDataByEth) ParseNinjaTweet() (*database.NinjaTweet, error) {
	var tweetContent database.NinjaTweet
	var err = json.Unmarshal([]byte(sp.Message), &tweetContent)
	if err != nil {
		util.LogInst().Err(err).Msg("Error parsing tweet ")
		return nil, err
	}

	if !tweetContent.IsValid() {
		util.LogInst().Warn().Msg("invalid tweet content:" + tweetContent.String())
		return nil, fmt.Errorf("invalid tweet content")
	}

	prefixedHash, err := util.Verify(tweetContent.Web3ID, sp.Message, sp.Signature)
	if err != nil {
		util.LogInst().Err(err).Msg("tweet signature verify failed")
		return nil, err
	}
	tweetContent.Signature = sp.Signature
	tweetContent.PrefixedHash = prefixedHash
	tweetContent.PaymentStatus = database.TxStNotPay
	var payloadStr, ok = sp.PayLoad.(string)
	if ok {
		var imagesI []TweetImgData
		if err := json.Unmarshal([]byte(payloadStr), &imagesI); err != nil {
			util.LogInst().Err(err).Msg("parse tweet img failed")
			return &tweetContent, nil
		}

		for _, imgData := range imagesI {
			tweetContent.Images = append(tweetContent.Images, imgData.ThumbNail)
			tweetContent.ImageHash = append(tweetContent.ImageHash, imgData.Hash)
			tweetContent.ImageRaw = append(tweetContent.ImageRaw, imgData.RawData)
		}
	}

	return &tweetContent, nil
}

func queryTwBasicById(w http.ResponseWriter, r *http.Request, nu *database.NinjaUsrInfo) {
	var needSyncFromTwitter = r.URL.Query().Get("forceSync")
	forceSync, err := strconv.ParseBool(needSyncFromTwitter)
	if err != nil {
		forceSync = false
	}

	var twitterID = r.URL.Query().Get("twitterID")
	if len(twitterID) == 0 {
		util.LogInst().Warn().Str("twitter-id", twitterID).
			Str("eth-addr", nu.EthAddr).Msg("invalid twitter id param")
		http.Error(w, "twitter id invalid", http.StatusBadRequest)
		return
	}

	var userdata *database.TWUserInfo
	if forceSync {
		if twitterID != nu.TwID && forceSync {
			util.LogInst().Warn().Str("twitter-id-query", twitterID).
				Str("real-twitter-id", nu.TwID).
				Str("eth-addr", nu.EthAddr).
				Msg("check twitter access token failed")
			http.Error(w, "no rights", http.StatusBadRequest)
			return
		}

		ut, err := checkTwitterRights(nu.TwID, r)
		if err != nil {
			util.LogInst().Warn().Str("twitter-id", twitterID).
				Str("eth-addr", nu.EthAddr).Msg("check twitter access token failed")
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		twitterUser, err := verifyTwitterCredentials(ut)
		if err != nil {
			util.LogInst().Warn().Str("twitter-id", twitterID).
				Str("eth-addr", nu.EthAddr).Msg("sync twitter server failed")
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		userdata = &database.TWUserInfo{
			ID:                   twitterUser.IDStr,
			Name:                 twitterUser.Name,
			ScreenName:           twitterUser.ScreenName,
			Description:          twitterUser.Description,
			ProfileImageUrlHttps: twitterUser.ProfileImageUrlHttps,
		}
		err = database.DbInst().UpdateBasicInfo(userdata)
		if err != nil {
			util.LogInst().Warn().Str("twitter-id", twitterID).
				Str("eth-addr", nu.EthAddr).Msg("update twitter new user data failed")
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	} else {
		userdata, err = database.DbInst().TwitterBasicInfo(twitterID)
		if err != nil {
			util.LogInst().Err(err).Str("twitter-id", twitterID).
				Str("eth-addr", nu.EthAddr).Msg("query twitter data failed")
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(userdata.RawData())
	util.LogInst().Debug().Str("twitter-id", twitterID).
		Str("eth-addr", nu.EthAddr).Msg("query twitter basic info success")
}

func showKolKeyPage(w http.ResponseWriter, r *http.Request, nu *database.NinjaUsrInfo) {
	var err = _globalCfg.htmlTemplateManager.ExecuteTemplate(w, "kol_key.html", nu)
	if err != nil {
		util.LogInst().Err(err).Msg("main html failed")
		http.Redirect(w, r, "/signIn", http.StatusFound)
		return
	}
}

type OuterLinkParam struct {
	TweetID  string `json:"tweet_id,omitempty"`
	ShareID  string `json:"share_id,omitempty"`
	ShareUsr string `json:"share_usr,omitempty"`
}

func (p *OuterLinkParam) Data() []byte {
	bts, _ := json.Marshal(p)
	return bts
}
func (p *OuterLinkParam) GetValidId() string {
	if len(p.TweetID) > 0 {
		return p.TweetID
	}
	return p.ShareID
}

func refreshNjUser(w http.ResponseWriter, r *http.Request, nu *database.NinjaUsrInfo) {
	newNu, err := database.DbInst().QueryNjUsrById(nu.EthAddr)
	if err != nil {
		util.LogInst().Err(err).Msg("load nj user failed")
		http.Error(w, "load nj user failed", http.StatusBadRequest)
		return
	}

	err = SMInst().Set(r, w, sesKeyForRightCheck, newNu.RawData())
	if err != nil {
		util.LogInst().Err(err).Msg("set session for nj user failed")
		http.Error(w, "set session for nj user failed", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(newNu.RawData())
}
func mainPage(w http.ResponseWriter, r *http.Request, nu *database.NinjaUsrInfo) {

	var param OuterLinkParam
	sData, err := SMInst().Get(BuyRightsUrlKey, r)
	if sData != nil {
		util.LogInst().Debug().Msg(string(sData.([]byte)))
		err = json.Unmarshal(sData.([]byte), &param)
		if err != nil {
			util.LogInst().Err(err).Msg("parse outer link param failed")
		}
		SMInst().Del(BuyRightsUrlKey, r, w)
	}

	var tweetId = param.GetValidId()
	var tweet *database.NinjaTweet
	if len(tweetId) > 0 {
		createAt, err := strconv.ParseInt(tweetId, 10, 60)
		if err != nil {
			util.LogInst().Err(err).Str("tweet-id", tweetId).Msg("invalid tweet id")
		} else {
			tweet, err = database.DbInst().NjTweetDetails(createAt)
			if err != nil {
				util.LogInst().Err(err).Str("tweet-id", tweetId).Msg("failed to load the tweet")
			}
		}
	}

	data := struct {
		NinjaUsrInfoJson template.JS
		TargetTweet      template.JS
		CSRFToken        string
	}{
		NinjaUsrInfoJson: template.JS(nu.RawData()),
		CSRFToken:        csrf.Token(r),
	}
	if tweet != nil {
		data.TargetTweet = template.JS(tweet.String())
	} else {
		data.TargetTweet = template.JS("{}")
	}

	err = _globalCfg.htmlTemplateManager.ExecuteTemplate(w, "main.html", data)
	if err != nil {
		util.LogInst().Err(err).Msg("main html failed")
		http.Redirect(w, r, "/signIn", http.StatusFound)
		return
	}
}

func showLotteryMain(w http.ResponseWriter, r *http.Request, nu *database.NinjaUsrInfo) {

	data := struct {
		NinjaUsrInfoJson template.JS
		CSRFToken        string
	}{
		CSRFToken:        csrf.Token(r),
		NinjaUsrInfoJson: template.JS(nu.RawData()),
	}
	var err = _globalCfg.htmlTemplateManager.ExecuteTemplate(w, "lottery_game.html", data)
	if err != nil {
		util.LogInst().Err(err).Msg("main html failed")
		http.Redirect(w, r, "/signIn", http.StatusFound)
		return
	}
}

func signOut(w http.ResponseWriter, r *http.Request, _ *database.NinjaUsrInfo) {
	_ = SMInst().Del(sesKeyForRightCheck, r, w)
	http.Redirect(w, r, "/signIn", http.StatusFound)
}

func validateUsrRights(r *http.Request) *database.NinjaUsrInfo {
	var data, err = SMInst().Get(sesKeyForRightCheck, r)
	if err != nil {
		util.LogInst().Warn().Msgf("%s", err.Error())
		return nil
	}

	var njUser, errNu = database.NJUsrInfoMust(data.([]byte))
	if errNu != nil {
		util.LogInst().Warn().Msgf("ninja user not found")
		return nil
	}
	return njUser
}

func bindingWeb3ID(w http.ResponseWriter, r *http.Request, origNu *database.NinjaUsrInfo) {
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
	twUsrData := &database.TWUserInfo{}
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
	_, err = util.Verify(data.EthAddr, param.Message, param.Signature)
	if err != nil {
		util.LogInst().Err(err).Msg("binding data verify signature failed")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	util.LogInst().Info().Str("eth-addr", data.EthAddr).
		Str("twitter-id", data.TwID).
		Int64("bind-time", data.BindTime).Msg("sign data success")

	bindDataToStore := &database.Web3Binding{
		TwitterID: data.TwID,
		EthAddr:   data.EthAddr,
		SignUpAt:  data.BindTime,
		Signature: param.Signature,
	}

	newNu, err := database.DbInst().BindingWeb3ID(bindDataToStore, twUsrData)
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

	//_ = updateTwitterBio(r, twUsrData.Description, data.EthAddr)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(newNu.RawData())

	util.LogInst().Info().Str("tweet-id", data.TwID).Str("web3-id", data.EthAddr).
		Int64("bind-time", data.BindTime).Msg("bind web3 and social id success")
}

func queryTwBasicByTweetHash(w http.ResponseWriter, r *http.Request, _ *database.NinjaUsrInfo) {
	var tweetHash = r.URL.Query().Get("tweet_hash")
	if len(tweetHash) == 0 {
		util.LogInst().Warn().Msg("invalid tweet hash")
		http.Error(w, "invalid tweet hash", http.StatusBadRequest)
		return
	}

	var twObj, err = database.DbInst().QueryTwUserByTweetHash(tweetHash)
	if err != nil {
		util.LogInst().Err(err).Str("tweet-hash", tweetHash).
			Msg("failed to query twitter info by hash")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(twObj.RawData())
}

func queryWinHistory(w http.ResponseWriter, r *http.Request, nu *database.NinjaUsrInfo) {

	var data, err = database.DbInst().QueryGameWinner(nu.EthAddr)
	if err != nil {
		util.LogInst().Err(err).Str("web3-id", nu.EthAddr).
			Msg("failed to query game winner")

		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	bts, _ := json.Marshal(data)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(bts)

	util.LogInst().Debug().Int("len", len(data)).Msg("query winner history success")
}
