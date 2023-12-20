package server

import (
	"encoding/json"
	"fmt"
	"github.com/dghubble/oauth1"
	"github.com/ninjahome/web-bridge/util"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const (
	sesKeyForNjUserId      = "twitter-signup-ninja-user-id"
	sesKeyForAccessToken   = "twitter-signup-access-token"
	callbackURL            = "https://bridge.simplenets.org/tw_callback"
	accessUserProUrl       = "https://api.twitter.com/1.1/account/update_profile.json"
	sesKeyForRequestSecret = "ses-key-for-request-secret"
	accessReqTokenURL      = "https://api.twitter.com/oauth/request_token"
	accessOauthTokenURL    = "https://api.twitter.com/oauth/authorize?oauth_token=%s"
	accessAccessTokenURL   = "https://api.twitter.com/oauth/access_token"
	accessUserInfoURL      = "https://api.twitter.com/2/users/me?user.fields=id,name,username,profile_image_url,description"
)

func signUpByTwitter(w http.ResponseWriter, r *http.Request) {
	ethAddr := r.URL.Query().Get("eth_addr")
	if ethAddr == "" {
		http.Error(w, "eth_addr parameter is required", http.StatusBadRequest)
		return
	}

	oauth1Config := oauth1.NewConfig(_globalCfg.ConsumerKey, _globalCfg.ConsumerSecret)
	oauth1Token := oauth1.NewToken("", "")
	httpClient := oauth1Config.Client(oauth1.NoContext, oauth1Token)

	callbackURL := url.QueryEscape(callbackURL)
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
	values, err := url.ParseQuery(bodyString)
	if err != nil {
		util.LogInst().Err(err).Msg("Failed to parse query from response body")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	requestToken := values.Get("oauth_token")
	requestSecret := values.Get("oauth_token_secret")

	_ = SMInst().Set(r, w, sesKeyForRequestSecret, requestSecret)
	_ = SMInst().Set(r, w, sesKeyForNjUserId, ethAddr)

	authorizeURL := fmt.Sprintf(accessOauthTokenURL, requestToken)
	http.Redirect(w, r, authorizeURL, http.StatusTemporaryRedirect)
}

func twitterSignCallBack(w http.ResponseWriter, r *http.Request) {
	oauth1Config := oauth1.NewConfig(_globalCfg.ConsumerKey, _globalCfg.ConsumerSecret)
	requestSecret, _ := SMInst().Get(sesKeyForRequestSecret, r)

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
	values, err := url.ParseQuery(bodyString)
	if err != nil {
		util.LogInst().Err(err).Msg("Failed to parse response body")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	accessToken := values.Get("oauth_token")
	accessSecret := values.Get("oauth_token_secret")
	token := oauth1.NewToken(accessToken, accessSecret)
	bts, _ := json.Marshal(token)
	_ = SMInst().Set(r, w, sesKeyForAccessToken, bts)
	http.Redirect(w, r, "/signUpSuccessByTw", http.StatusFound)
}

func getAccessTokenFromSession(r *http.Request) (*oauth1.Token, error) {
	bts, err := SMInst().Get(sesKeyForAccessToken, r)
	if err != nil {
		return nil, err
	}
	token := &oauth1.Token{}
	err = json.Unmarshal(bts.([]byte), token)
	return token, err
}

func signUpSuccessByTw(w http.ResponseWriter, r *http.Request) {
	ethAddr, err := SMInst().Get(sesKeyForNjUserId, r)
	if err != nil {
		util.LogInst().Err(err).Msg("no valid ninja id found")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	token, err := getAccessTokenFromSession(r)
	if err != nil {
		util.LogInst().Err(err).Msg("no user access token found")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	result, err := fetchTwitterUserInfo(token)
	if err != nil {
		util.LogInst().Err(err).Msg("get user basic info failed")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	result.EthAddr = ethAddr.(string)
	result.SignUpAt = time.Now().UnixMilli()
	err = htmlTemplateManager.ExecuteTemplate(w, "signUpSuccess.html", result)
	if err != nil {
		util.LogInst().Err(err).Msg("show sign up by twitter page failed")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	util.LogInst().Debug().Str("tw-id", result.TwitterData.ID).Str("ninja-id", result.EthAddr).Msg("twitter user sign up success")
}

func updateTwitterBio(token *oauth1.Token, newBio string) error {
	config := oauth1.NewConfig(_globalCfg.ConsumerKey, _globalCfg.ConsumerSecret)
	httpClient := config.Client(oauth1.NoContext, token)

	values := url.Values{}
	values.Add("description", newBio)
	req, err := http.NewRequest("POST", accessUserProUrl, strings.NewReader(values.Encode()))
	if err != nil {
		util.LogInst().Err(err).Msg("updateTwitterBio NewRequest failed")
		return err
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, err := httpClient.Do(req)
	if err != nil {
		util.LogInst().Err(err).Msg("updateTwitterBio httpClient Do failed")
		return err
	}
	defer resp.Body.Close()
	return nil
}

func fetchTwitterUserInfo(token *oauth1.Token) (*TwitterAPIResponse, error) {
	oauth1Config := oauth1.NewConfig("consumerKey", "consumerSecret")
	httpClient := oauth1Config.Client(oauth1.NoContext, token)

	resp, err := httpClient.Get(accessUserInfoURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("twitter API responded with status: %s", resp.Status)
	}

	result := &TwitterAPIResponse{}
	err = json.NewDecoder(resp.Body).Decode(result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
