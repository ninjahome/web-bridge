package server

import (
	"context"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"golang.org/x/oauth2"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
)

const (
	callbackURL = "https://bridge.simplenets.org/tw_callback"
	//callbackURL    = "http//127.0.0.1/tw_callback"
	authorizeURL   = "https://twitter.com/i/oauth2/authorize"
	accessTokenURL = "https://api.twitter.com/2/oauth2/token"
)

type TwitterSrv struct {
	oauth2Config *oauth2.Config
}

func NewTwitterSrv(conf *TwitterConf) *TwitterSrv {
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

	return &TwitterSrv{oauth2Config: oauth2Config}
}
func randomBytesInHex(count int) (string, error) {
	buf := make([]byte, count)
	_, err := io.ReadFull(rand.Reader, buf)
	if err != nil {
		return "", fmt.Errorf("Could not generate %d random bytes: %v", count, err)
	}

	return hex.EncodeToString(buf), nil
}

var tmpVerifierCode = ""

func signInByTwitter(ts *TwitterSrv, w http.ResponseWriter, r *http.Request) {
	codeVerifier, verifierErr := randomBytesInHex(32) // 64 character string here
	if verifierErr != nil {
		return
	}
	sha2 := sha256.New()
	io.WriteString(sha2, codeVerifier)
	tmpVerifierCode = codeVerifier
	codeChallenge := base64.RawURLEncoding.EncodeToString(sha2.Sum(nil))
	state, _ := randomBytesInHex(24)
	oauthUrl := ts.oauth2Config.AuthCodeURL(state) + "&code_challenge=" + url.QueryEscape(codeChallenge) + "&code_challenge_method=S256"
	//fmt.Printf("Go to the following link in your browser then type the \"code\" parameter here:\n%s\n", url)
	http.Redirect(w, r, oauthUrl, http.StatusTemporaryRedirect)
}

func exchangeWithCodeVerifier(ctx context.Context, conf *oauth2.Config, code string, codeVerifier string) (*oauth2.Token, error) {
	// Prepare the request URL and body
	values := url.Values{}
	values.Add("client_id", conf.ClientID)
	values.Add("grant_type", "authorization_code")
	values.Add("code", code)
	values.Add("redirect_uri", conf.RedirectURL)
	values.Add("code_verifier", codeVerifier)
	queryStr := strings.NewReader(values.Encode())
	fmt.Println("queryStr:", queryStr)
	fmt.Println("TokenURL:", conf.Endpoint.TokenURL)
	req, err := http.NewRequestWithContext(ctx, "POST", conf.Endpoint.TokenURL, queryStr)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	// Send the request
	httpClient := &http.Client{}
	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Parse the response
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %v", resp.Status)
	}
	var token oauth2.Token
	if err := json.NewDecoder(resp.Body).Decode(&token); err != nil {
		return nil, err
	}

	return &token, nil
}

func twitterSignCallBack(ts *TwitterSrv, w http.ResponseWriter, r *http.Request) {
	log.Println("receive call back from twitter")
	code := r.URL.Query().Get("code")
	err2 := r.URL.Query().Get("error")
	state := r.URL.Query().Get("state")
	fmt.Println("code:", code)
	fmt.Println("error")
	fmt.Println(err2, "state", state)
	ctx := context.Background()

	token, err := exchangeWithCodeVerifier(ctx, ts.oauth2Config, code, tmpVerifierCode)
	//ts.oauth2Config.Exchange(ctx, code)
	if err != nil {
		log.Println("exchange err:", err)
		return
	}
	fmt.Println(token, token.RefreshToken)

	if err := ts.saveRefreshToken(token.RefreshToken, state); err != nil {
		return
	}

	client := ts.oauth2Config.Client(context.Background(), token)
	response, err3 := client.Get("https://api.twitter.com/2/users/me")
	if err3 != nil {
		log.Println(" client.Get err:", err3)
		return
	}
	defer response.Body.Close()
	fmt.Println(response)
}
func (ts *TwitterSrv) saveRefreshToken(refreshToken, state string) error {
	return nil
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
