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
	"time"
)

const (
	sesKeyForStateV2        = "twitter-signup-state-key"
	sesKeyForVerifierCodeV2 = "code_verifier"
	sesKeyForAccessTokenV2  = "twitter-access-key-v2"

	authorizeURLV2   = "https://twitter.com/i/oauth2/authorize"
	accessTokenURLV2 = "https://api.twitter.com/2/oauth2/token"
	accessUserURLV2  = "https://api.twitter.com/2/users/me?user.fields=profile_image_url,description"
)

type stateParam struct {
	ethAddr string
	stateNo string
}

func (sp *stateParam) String() string {
	return sp.ethAddr + "_" + sp.stateNo
}

func parseStateParam(str string) *stateParam {
	var strArr = strings.Split(str, "_")
	if len(strArr) != 2 {
		return nil
	}
	return &stateParam{ethAddr: strArr[0], stateNo: strArr[1]}
}

func signUpByTwitterV2(w http.ResponseWriter, r *http.Request) {

	ethAddr := r.URL.Query().Get("eth_addr")
	if ethAddr == "" {
		http.Error(w, "eth_addr parameter is required", http.StatusBadRequest)
		return
	}
	codeVerifier := util.RandomBytesInHex(32)

	sha2 := sha256.New()
	_, _ = io.WriteString(sha2, codeVerifier)
	codeChallenge := base64.RawURLEncoding.EncodeToString(sha2.Sum(nil))

	state := stateParam{ethAddr: ethAddr, stateNo: util.RandomBytesInHex(24)}
	var stateStr = state.String()
	var err1, err2 = SMInst().Set(r, w, sesKeyForStateV2, stateStr), SMInst().Set(r, w, sesKeyForVerifierCodeV2, codeVerifier)
	if err1 != nil || err2 != nil {
		util.LogInst().Err(err1).Err(err2).Msg("session set error for twitter signIn")
		http.Error(w, "sign up error for session", http.StatusInternalServerError)
		return
	}

	oauthUrl := _globalCfg.twOauthCfg.AuthCodeURL(stateStr) + "&code_challenge=" + url.QueryEscape(codeChallenge) + "&code_challenge_method=S256"
	http.Redirect(w, r, oauthUrl, http.StatusTemporaryRedirect)
}

func exchangeWithCodeVerifier(conf *oauth2.Config, code string, codeVerifier string) (*oauth2.Token, error) {
	ctx := context.Background()
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

func twitterSignCallBackV2(w http.ResponseWriter, r *http.Request) {
	util.LogInst().Info().Msg("call back from twitter")

	errStr := r.URL.Query().Get("error")
	if len(errStr) > 0 {
		util.LogInst().Warn().Msgf("twitter call back has err")
		http.Error(w, errStr, http.StatusInternalServerError)
		return
	}

	originalState, err := SMInst().Get(sesKeyForStateV2, r)
	if err != nil {
		util.LogInst().Err(err).Msg("failed to get state when twitter user call back")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	state := r.URL.Query().Get("state")
	if state != originalState {
		util.LogInst().Warn().Msg("invalid state in twitter sign up call back")
		http.Error(w, "invalid state in twitter sign up call back", http.StatusInternalServerError)
		return
	}

	code := r.URL.Query().Get("code")

	codeVerifier, err := SMInst().Get(sesKeyForVerifierCodeV2, r)
	if err != nil {
		util.LogInst().Err(err).Msg("get verifier code failed")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer SMInst().Del(sesKeyForVerifierCodeV2, r, w)
	token, errToken := exchangeWithCodeVerifier(_globalCfg.twOauthCfg, code, codeVerifier.(string))
	if errToken != nil {
		util.LogInst().Err(errToken).Msgf("exchange err:%s", errToken)
		http.Error(w, errToken.Error(), http.StatusInternalServerError)
		return
	}

	bts, _ := json.Marshal(token)
	err = SMInst().Set(r, w, sesKeyForAccessTokenV2, bts)
	if err != nil {
		util.LogInst().Err(err).Msgf("save twitter acccess token failed")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/signUpSuccessByTwV2", http.StatusFound)
}

func signUpSuccessByTwV2(w http.ResponseWriter, r *http.Request) {
	stateStr, _ := SMInst().Get(sesKeyForStateV2, r)
	defer SMInst().Del(sesKeyForStateV2, r, w)
	state := parseStateParam(stateStr.(string))
	if state == nil {
		util.LogInst().Warn().Msg("state lost for twitter sign up session")
		http.Error(w, "state lost for twitter sign up session", http.StatusInternalServerError)
		return
	}

	token, err := getAccessTokenFromSessionV2(r)
	if err != nil {
		util.LogInst().Err(err).Msg("no valid access token in current session")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	result := &TwAPIResponse{}
	err = twitterGetWithAccessToken(token, accessUserURLV2, result)
	if err != nil {
		util.LogInst().Err(err).Msgf("get twitter info failed")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	result.EthAddr = state.ethAddr
	result.SignUpAt = time.Now().UnixMilli()
	ut := &TwUserAccessTokenV2{
		UserId: result.TwitterData.ID,
		Token:  token,
	}
	err = DbInst().SaveTwAccessTokenV2(ut)
	if err != nil {
		util.LogInst().Err(err).Msg("save user access token failed")
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	err = htmlTemplateManager.ExecuteTemplate(w, "signUpSuccess.html", result)
	if err != nil {
		util.LogInst().Err(err).Msg("show sign up by twitter page failed")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func refreshAccessToken(refreshToken string) (*oauth2.Token, error) {
	ctx := context.Background()
	tokenSource := _globalCfg.twOauthCfg.TokenSource(ctx, &oauth2.Token{RefreshToken: refreshToken})
	newToken, err := tokenSource.Token()
	if err != nil {
		return nil, err
	}
	return newToken, nil
}
