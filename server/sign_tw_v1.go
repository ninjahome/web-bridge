package server

import (
	"encoding/json"
	"fmt"
	"github.com/dghubble/oauth1"
	database2 "github.com/ninjahome/web-bridge/database"
	"github.com/ninjahome/web-bridge/util"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const (
	Web3IDProfile          = "Dessage-Web3-ID:"
	sesKeyForNjUserId      = "twitter-signup-ninja-user-id"
	sesKeyForAccessTokenV1 = "twitter-access-key-v1"
	sesKeyForRequestSecret = "ses-key-for-request-secret"
	accessUserProUrl       = "https://api.twitter.com/1.1/account/update_profile.json"
	accessReqTokenURL      = "https://api.twitter.com/oauth/request_token"
	accessOauthTokenURL    = "https://api.twitter.com/oauth/authorize?oauth_token=%s"
	accessAccessTokenURL   = "https://api.twitter.com/oauth/access_token"
	verifyCredentialsURL   = "https://api.twitter.com/1.1/account/verify_credentials.json?skip_status=true"
)

func parseUserToken(values url.Values) *database2.TwUserAccessToken {
	accessToken := values.Get("oauth_token")
	accessSecret := values.Get("oauth_token_secret")
	userID := values.Get("user_id")
	screenName := values.Get("screen_name")
	return &database2.TwUserAccessToken{
		OauthToken:       accessToken,
		OauthTokenSecret: accessSecret,
		UserId:           userID,
		ScreenName:       screenName,
	}
}

func getAccessTokenFromSession(r *http.Request) (*database2.TwUserAccessToken, error) {
	bts, err := SMInst().Get(sesKeyForAccessTokenV1, r)
	if err != nil {
		return nil, err
	}
	var token database2.TwUserAccessToken
	err = json.Unmarshal([]byte(bts.(string)), &token)
	if err != nil {
		return nil, err
	}
	return &token, nil
}

func signUpByTwitterV1(w http.ResponseWriter, r *http.Request, nu *database2.NinjaUsrInfo) {

	oauth1Config := oauth1.NewConfig(_globalCfg.ConsumerKey, _globalCfg.ConsumerSecret)
	oauth1Token := oauth1.NewToken("", "")
	httpClient := oauth1Config.Client(oauth1.NoContext, oauth1Token)

	callbackURL := url.QueryEscape(twitterSignUpCallbackURL)
	response, err := httpClient.PostForm(accessReqTokenURL, url.Values{"oauth_callback": {callbackURL}})
	if err != nil {
		util.LogInst().Err(err).Msg("Failed to get request token")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer response.Body.Close()
	bodyBytes, err := io.ReadAll(response.Body)
	if err != nil {
		util.LogInst().Err(err).Msg("Failed to read response body")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	bodyString := string(bodyBytes)

	if response.StatusCode != http.StatusOK {
		http.Error(w, bodyString, http.StatusInternalServerError)
		util.LogInst().Warn().Msg(bodyString)
		return
	}

	values, err := url.ParseQuery(bodyString)
	if err != nil {
		util.LogInst().Err(err).Msg("Failed to parse query from response body")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	requestToken := values.Get("oauth_token")
	requestSecret := values.Get("oauth_token_secret")

	var err1, err2 = SMInst().Set(r, w, sesKeyForRequestSecret, requestSecret), SMInst().Set(r, w, sesKeyForNjUserId, nu.EthAddr)
	if err1 != nil || err2 != nil {
		util.LogInst().Err(err1).Err(err2).Msg("save secret or eth id to session failed")
		http.Error(w, "session save failed", http.StatusInternalServerError)
		return
	}
	authorizeURL := fmt.Sprintf(accessOauthTokenURL, requestToken)
	http.Redirect(w, r, authorizeURL, http.StatusTemporaryRedirect)
}

func twitterSignCallBackV1(w http.ResponseWriter, r *http.Request, _ *database2.NinjaUsrInfo) {
	oauth1Config := oauth1.NewConfig(_globalCfg.ConsumerKey, _globalCfg.ConsumerSecret)
	requestSecret, err := SMInst().Get(sesKeyForRequestSecret, r)
	if err != nil {
		util.LogInst().Err(err).Msg("get secret from session failed")
		http.Error(w, "get secret from session failed", http.StatusInternalServerError)
		return
	}
	defer SMInst().Del(sesKeyForRequestSecret, r, w)

	requestToken := r.URL.Query().Get("oauth_token")
	verifier := r.URL.Query().Get("oauth_verifier")

	params := url.Values{
		"oauth_token":    {requestToken},
		"oauth_verifier": {verifier},
	}
	httpClient := oauth1Config.Client(oauth1.NoContext, oauth1.NewToken(requestToken, requestSecret.(string)))
	resp, err := httpClient.PostForm(accessAccessTokenURL, params)
	if err != nil {
		util.LogInst().Err(err).Msg("Failed to request access token")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		util.LogInst().Err(err).Msg("Failed to read response body")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	bodyString := string(bodyBytes)
	if resp.StatusCode != http.StatusOK {
		util.LogInst().Warn().Msg(bodyString)
		http.Redirect(w, r, "/main", http.StatusFound)
		return
	}

	values, err := url.ParseQuery(bodyString)
	if err != nil {
		util.LogInst().Err(err).Msg("Failed to parse response body")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	token := parseUserToken(values)
	_ = SMInst().Set(r, w, sesKeyForAccessTokenV1, token.String())
	err = database2.DbInst().SaveTwAccessToken(token)
	if err != nil {
		util.LogInst().Err(err).Msg("save twitter user access token failed")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/signUpSuccessByTw", http.StatusFound)
}

func signUpSuccessByTw(w http.ResponseWriter, r *http.Request, _ *database2.NinjaUsrInfo) {
	ethAddr, err := SMInst().Get(sesKeyForNjUserId, r)
	if err != nil {
		util.LogInst().Err(err).Msg("no valid ninja id found")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer SMInst().Del(sesKeyForNjUserId, r, w)

	token, err := getAccessTokenFromSession(r)
	if err != nil {
		util.LogInst().Err(err).Msg("no user access token found")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	userData, err := verifyTwitterCredentials(token)
	if err != nil {
		util.LogInst().Err(err).Msg("get user basic info failed")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	result := &TwAPIResponse{
		EthAddr:  ethAddr.(string),
		SignUpAt: time.Now().UnixMilli(),
		TwitterData: &database2.TWUserInfo{
			ID:                   userData.IDStr,
			Name:                 userData.Name,
			ScreenName:           userData.ScreenName,
			Description:          userData.Description,
			ProfileImageUrlHttps: userData.ProfileImageUrlHttps,
		},
	}

	err = _globalCfg.htmlTemplateManager.ExecuteTemplate(w, "signUpSuccess.html", result)
	if err != nil {
		util.LogInst().Err(err).Msg("show sign up by twitter page failed")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	util.LogInst().Debug().Str("tw-id", result.TwitterData.ID).Str("ninja-id", result.EthAddr).Msg("twitter user sign up success")
}

func updateTwitterBio(r *http.Request, origDes, web3ID string) error {
	ut, err := getAccessTokenFromSession(r)
	if err != nil {
		util.LogInst().Err(err).Msg("no user access token found")
		return err
	}
	if strings.Contains(origDes, web3ID) {
		util.LogInst().Info().Msg("twitter user has got web3 id")
		return nil
	}
	var des = origDes
	var idx = strings.Index(origDes, Web3IDProfile)
	if idx >= 0 {
		des = origDes[:idx+len(Web3IDProfile)] + web3ID + origDes[idx+len(Web3IDProfile)+42:]
	} else {
		des = origDes + "\n" + Web3IDProfile + web3ID
	}

	values := url.Values{}
	values.Add("description", des)
	return twitterApiPost(accessUserProUrl, ut.GetToken(), strings.NewReader(values.Encode()),
		"application/x-www-form-urlencoded", nil)
}

func verifyTwitterCredentials(ut *database2.TwUserAccessToken) (*VerifiedTwitterUser, error) {
	config := oauth1.NewConfig(_globalCfg.ConsumerKey, _globalCfg.ConsumerSecret)
	httpClient := config.Client(oauth1.NoContext, ut.GetToken())

	resp, err := httpClient.Get(verifyCredentialsURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("twitter API responded with status: %s", resp.Status)
	}

	var verifiedUser VerifiedTwitterUser
	if err := json.NewDecoder(resp.Body).Decode(&verifiedUser); err != nil {
		return nil, err
	}
	//bts, _ := json.Marshal(verifiedUser)
	//util.LogInst().Debug().Msg(string(bts))
	return &verifiedUser, nil
}
