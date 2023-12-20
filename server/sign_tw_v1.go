package server

import (
	"fmt"
	"github.com/dghubble/oauth1"
	"github.com/ninjahome/web-bridge/util"
	"io"
	"net/http"
	"net/url"
	"strings"
)

const (
	callbackURL            = "https://bridge.simplenets.org/tw_callback"
	accessUserProUrl       = "https://api.twitter.com/1.1/account/update_profile.json"
	sesKeyForRequestSecret = "ses-key-for-request-secret"
)

func signUpByTwitter(w http.ResponseWriter, r *http.Request) {
	oauth1Config := oauth1.NewConfig(_globalCfg.ConsumerKey, _globalCfg.ConsumerSecret)
	oauth1Token := oauth1.NewToken("", "")
	httpClient := oauth1Config.Client(oauth1.NoContext, oauth1Token)

	requestTokenURL := "https://api.twitter.com/oauth/request_token"
	callbackURL := url.QueryEscape(callbackURL)
	response, err := httpClient.PostForm(requestTokenURL, url.Values{"oauth_callback": {callbackURL}})
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
	authorizeURL := fmt.Sprintf("https://api.twitter.com/oauth/authorize?oauth_token=%s", requestToken)
	http.Redirect(w, r, authorizeURL, http.StatusTemporaryRedirect)
}

func twitterSignCallBack(w http.ResponseWriter, r *http.Request) {
	oauth1Config := oauth1.NewConfig(_globalCfg.ConsumerKey, _globalCfg.ConsumerSecret)
	requestSecret, _ := SMInst().Get(sesKeyForRequestSecret, r)

	requestToken := r.URL.Query().Get("oauth_token")
	verifier := r.URL.Query().Get("oauth_verifier")
	util.LogInst().Debug().Str("requestToken", requestToken).Str("verifier", verifier).
		Str("requestSecret", requestSecret.(string)).Send()

	accessTokenURL := "https://api.twitter.com/oauth/access_token"
	params := url.Values{
		"oauth_token":    {requestToken},
		"oauth_verifier": {verifier},
	}
	httpClient := oauth1Config.Client(oauth1.NoContext, oauth1.NewToken(requestToken, requestSecret.(string)))
	resp, err := httpClient.PostForm(accessTokenURL, params)
	if err != nil {
		util.LogInst().Err(err).Msg("Failed to request access token")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()
	// 解析响应体以获取访问令牌和访问令牌秘密
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
	util.LogInst().Debug().Str("accessToken", accessToken).Str("accessSecret", accessSecret).Send()
	err = updateTwitterBio(token, "web player")
	if err != nil {
		util.LogInst().Err(err).Msg("updateTwitterBio failed")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
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
