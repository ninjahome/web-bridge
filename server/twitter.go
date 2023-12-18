package server

import (
	"context"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/ninjahome/web-bridge/util"
	"golang.org/x/oauth2"
	"io"
	"net/http"
	"net/url"
	"strings"
)

const (
	callbackURL    = "https://bridge.simplenets.org/tw_callback"
	authorizeURL   = "https://twitter.com/i/oauth2/authorize"
	accessTokenURL = "https://api.twitter.com/2/oauth2/token"
	accessUserURL  = "https://api.twitter.com/2/users/me"

	sessionKeyForUser = "twitter-user-info"
	verifierCodeKey   = "code_verifier"
)

type TwitterSrv struct {
	oauth2Config *oauth2.Config
}

func NewTwitterSrv() *TwitterSrv {
	conf := _globalCfg.TwitterConf
	var oauth2Config = &oauth2.Config{
		RedirectURL:  callbackURL,
		ClientID:     conf.ClientID,
		ClientSecret: conf.ClientSecret,
		Scopes:       []string{"tweet.read", "tweet.write", "follows.read", "follows.write", "users.read", "offline.access"},
		Endpoint: oauth2.Endpoint{
			AuthURL:  authorizeURL,
			TokenURL: accessTokenURL,
		},
	}
	htmlTemplateManager = util.ParseTemplates("assets/html")

	return &TwitterSrv{oauth2Config: oauth2Config}
}

func signInByTwitter(ts *TwitterSrv, w http.ResponseWriter, r *http.Request) {
	codeVerifier, verifierErr := util.RandomBytesInHex(32) // 64 character string here
	if verifierErr != nil {
		return
	}
	sha2 := sha256.New()
	_, err := io.WriteString(sha2, codeVerifier)
	if err != nil {
		util.LogInst().Err(err).Msg("creating verifier code failed")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = SMInst().Set(r, w, verifierCodeKey, codeVerifier)
	if err != nil {
		util.LogInst().Err(err).Msg("save verifier code failed")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	util.LogInst().Debug().Str("codeVerifier", codeVerifier).Send()
	codeChallenge := base64.RawURLEncoding.EncodeToString(sha2.Sum(nil))
	state, _ := util.RandomBytesInHex(24)
	oauthUrl := ts.oauth2Config.AuthCodeURL(state) + "&code_challenge=" + url.QueryEscape(codeChallenge) + "&code_challenge_method=S256"
	http.Redirect(w, r, oauthUrl, http.StatusTemporaryRedirect)
}

func exchangeWithCodeVerifier(ctx context.Context, conf *oauth2.Config, code string, codeVerifier string) (*oauth2.Token, error) {
	values := url.Values{}
	values.Add("client_id", conf.ClientID)
	values.Add("grant_type", "authorization_code")
	values.Add("code", code)
	values.Add("redirect_uri", conf.RedirectURL)
	values.Add("code_verifier", codeVerifier)
	queryStr := strings.NewReader(values.Encode())
	req, err := http.NewRequestWithContext(ctx, "POST", conf.Endpoint.TokenURL, queryStr)
	if err != nil {
		util.LogInst().Err(err).Msg("NewRequestWithContext failed")
		return nil, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	auth := conf.ClientID + ":" + conf.ClientSecret
	basicAuth := base64.StdEncoding.EncodeToString([]byte(auth))
	req.Header.Set("Authorization", "Basic "+basicAuth)

	httpClient := &http.Client{}
	resp, err := httpClient.Do(req)
	if err != nil {
		util.LogInst().Err(err).Msg("http do failed")
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %v", resp.Status)
	}
	var token oauth2.Token
	if err := json.NewDecoder(resp.Body).Decode(&token); err != nil {
		util.LogInst().Err(err).Msg("decode exchange body failed")
		return nil, err
	}

	return &token, nil
}

func twitterSignCallBack(ts *TwitterSrv, w http.ResponseWriter, r *http.Request) {
	util.LogInst().Info().Msg("call back from twitter")

	errStr := r.URL.Query().Get("error")
	if len(errStr) > 0 {
		util.LogInst().Warn().Msgf("twitter call back has err:%s", errStr)
		http.Error(w, errStr, http.StatusInternalServerError)
		return
	}
	code := r.URL.Query().Get("code")
	state := r.URL.Query().Get("state")
	ctx := context.Background()

	codeVerifier, err := SMInst().Get(verifierCodeKey, r)
	if err != nil {
		util.LogInst().Err(err).Msg("get verifier code failed")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	token, errToken := exchangeWithCodeVerifier(ctx, ts.oauth2Config, code, codeVerifier.(string))
	if errToken != nil {
		util.LogInst().Err(errToken).Msgf("exchange err:%s", errToken)
		http.Error(w, errToken.Error(), http.StatusInternalServerError)
		return
	}

	ts.saveRefreshToken(token.RefreshToken, state)

	client := ts.oauth2Config.Client(context.Background(), token)
	response, err3 := client.Get(accessUserURL)
	if err3 != nil {
		util.LogInst().Err(err).Msgf("create client err:%s", err3)
		http.Error(w, err3.Error(), http.StatusInternalServerError)
		return
	}
	defer response.Body.Close()

	var apiResponse TwitterAPIResponse
	if err := json.NewDecoder(response.Body).Decode(&apiResponse); err != nil {
		util.LogInst().Err(err).Msgf("parse twitter call back data  err:%s", err)
		http.Error(w, err3.Error(), http.StatusInternalServerError)
		return
	}

	result := apiResponse.Data
	err = SMInst().Set(r, w, sessionKeyForUser, result.String())
	if err != nil {
		util.LogInst().Err(err).Msgf("save twitter info failed:%v", result)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/main", http.StatusFound)
}

func showMainPage(ts *TwitterSrv, w http.ResponseWriter, r *http.Request) {
	resultStr, err := SMInst().Get(sessionKeyForUser, r)
	if err != nil {
		util.LogInst().Err(err).Msg("no twitter user info found")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	result := TWUsrInfoMust(resultStr.(string))
	err = htmlTemplateManager.ExecuteTemplate(w, "main.html", result)
	if err != nil {
		util.LogInst().Err(err).Msg("show main page failed")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (ts *TwitterSrv) saveRefreshToken(refreshToken, state string) {
}

func refreshAccessToken(ts *TwitterSrv, refreshToken string) (*oauth2.Token, error) {
	ctx := context.Background()
	tokenSource := ts.oauth2Config.TokenSource(ctx, &oauth2.Token{RefreshToken: refreshToken})
	newToken, err := tokenSource.Token()
	if err != nil {
		return nil, err
	}
	return newToken, nil
}
