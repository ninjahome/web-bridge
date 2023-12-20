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
	sesKeyForAccessToken   = "twitter-access-key-v1"
	callbackURL            = "https://bridge.simplenets.org/tw_callback"
	accessUserProUrl       = "https://api.twitter.com/1.1/account/update_profile.json"
	sesKeyForRequestSecret = "ses-key-for-request-secret"
	accessReqTokenURL      = "https://api.twitter.com/oauth/request_token"
	accessOauthTokenURL    = "https://api.twitter.com/oauth/authorize?oauth_token=%s"
	accessAccessTokenURL   = "https://api.twitter.com/oauth/access_token"
)

type userAccessToken struct {
	OauthToken       string
	OauthTokenSecret string
	UserId           string
	ScreenName       string
}

func parseUserToken(values url.Values) *userAccessToken {
	accessToken := values.Get("oauth_token")
	accessSecret := values.Get("oauth_token_secret")
	userID := values.Get("user_id")
	screenName := values.Get("screen_name")
	return &userAccessToken{
		OauthToken:       accessToken,
		OauthTokenSecret: accessSecret,
		UserId:           userID,
		ScreenName:       screenName,
	}
}

func (ut *userAccessToken) GetToken() *oauth1.Token {
	return &oauth1.Token{
		Token:       ut.OauthToken,
		TokenSecret: ut.OauthTokenSecret,
	}
}

func (ut *userAccessToken) string() string {
	bts, _ := json.Marshal(ut)
	return string(bts)
}

func getAccessTokenFromSession(r *http.Request) (*userAccessToken, error) {
	bts, err := SMInst().Get(sesKeyForAccessToken, r)
	if err != nil {
		return nil, err
	}
	var token userAccessToken
	err = json.Unmarshal([]byte(bts.(string)), &token)
	if err != nil {
		return nil, err
	}
	return &token, nil
}

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

	token := parseUserToken(values)
	_ = SMInst().Set(r, w, sesKeyForAccessToken, token.string())
	http.Redirect(w, r, "/signUpSuccessByTw", http.StatusFound)
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

	result, err := verifyTwitterCredentials(token) //fetchTwitterUserInfo(token)//verifyTwitterCredentials
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

func updateTwitterBio(ut *userAccessToken, newBio string) error {
	config := oauth1.NewConfig(_globalCfg.ConsumerKey, _globalCfg.ConsumerSecret)
	httpClient := config.Client(oauth1.NoContext, ut.GetToken())

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

func fetchTwitterUserInfo(ut *userAccessToken) (*TwAPIResponse, error) {
	config := oauth1.NewConfig(_globalCfg.ConsumerKey, _globalCfg.ConsumerSecret)

	util.LogInst().Debug().Msg(ut.string())

	httpClient := config.Client(oauth1.NoContext, ut.GetToken())

	userInfoURL := fmt.Sprintf("https://api.twitter.com/1.1/users/show.json?screen_name=%s", ut.ScreenName)
	//userInfoURL := fmt.Sprintf("https://api.twitter.com/1.1/users/show.json?user_id=%s", ut.UserId)

	resp, err := httpClient.Get(userInfoURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("twitter API responded with status: %s", resp.Status)
	}

	var user TwAPIResponse
	if err := json.NewDecoder(resp.Body).Decode(&user); err != nil {
		return nil, err
	}

	return &user, nil
}

func verifyTwitterCredentials(ut *userAccessToken) (*TwAPIResponse, error) {
	config := oauth1.NewConfig(_globalCfg.ConsumerKey, _globalCfg.ConsumerSecret)
	httpClient := config.Client(oauth1.NoContext, ut.GetToken())

	verifyCredentialsURL := "https://api.twitter.com/1.1/account/verify_credentials.json"

	resp, err := httpClient.Get(verifyCredentialsURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("twitter API responded with status: %s", resp.Status)
	}

	var user TwAPIResponse
	if err := json.NewDecoder(resp.Body).Decode(&user); err != nil {
		return nil, err
	}

	return &user, nil
}
