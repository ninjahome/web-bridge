package server

import (
	"fmt"
	"github.com/dghubble/oauth1"
	"io"
	"net/http"
	"net/url"
	"strings"
)

const (
	callbackURL = "https://bridge.simplenets.org/tw_callback"
)

func signUpByTwitter(w http.ResponseWriter, r *http.Request) {
	config := oauth1.NewConfig(_globalCfg.ConsumerKey, _globalCfg.ConsumerSecret)
	requestTokenURL := "https://api.twitter.com/oauth/request_token"
	authorizeURL := "https://api.twitter.com/oauth/authorize"

	// 创建一个OAuth1的HTTP客户端
	httpClient := config.Client(oauth1.NoContext, nil)

	// Step 1: 获取请求令牌
	resp, err := httpClient.PostForm(requestTokenURL, url.Values{
		"oauth_callback": {callbackURL},
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	// 解析响应以获取请求令牌
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	bodyString := string(bodyBytes)
	values, err := url.ParseQuery(bodyString)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	requestToken := values.Get("oauth_token")

	// Step 2: 重定向用户到Twitter进行授权
	authURL := fmt.Sprintf("%s?oauth_token=%s", authorizeURL, requestToken)
	http.Redirect(w, r, authURL, http.StatusTemporaryRedirect)
}

func twitterSignCallBack(w http.ResponseWriter, r *http.Request) {
	config := oauth1.NewConfig("consumerKey", "consumerSecret")
	accessTokenURL := "https://api.twitter.com/oauth/access_token"

	requestToken := r.URL.Query().Get("oauth_token")
	verifier := r.URL.Query().Get("oauth_verifier")

	// Step 3: Exchange request token and verifier for an access token
	accessToken, accessSecret, err := config.AccessToken(accessTokenURL, requestToken, verifier)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	token := oauth1.NewToken(accessToken, accessSecret)
	err = updateTwitterBio(token, "web player")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func updateTwitterBio(token *oauth1.Token, newBio string) error {
	config := oauth1.NewConfig(_globalCfg.ConsumerKey, _globalCfg.ConsumerSecret)
	httpClient := config.Client(oauth1.NoContext, token)

	values := url.Values{}
	values.Add("description", newBio)
	req, err := http.NewRequest("POST", accessUserProUrlV2, strings.NewReader(values.Encode()))
	if err != nil {
		return err
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, err := httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return nil
}
