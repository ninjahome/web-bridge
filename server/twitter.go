package server

import (
	"context"
	"fmt"
	"golang.org/x/oauth2"
	"net/http"
)

const (
	callbackURL    = "https://bridge.simplenets.org:8443/tw_callback"
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
		Scopes:       []string{"tweet.read", "users.read", "offline.access"},
		Endpoint: oauth2.Endpoint{
			AuthURL:  authorizeURL,
			TokenURL: accessTokenURL,
		},
	}

	return &TwitterSrv{oauth2Config: oauth2Config}
}

func signInByTwitter(ts *TwitterSrv, w http.ResponseWriter, r *http.Request) {
	url := ts.oauth2Config.AuthCodeURL("state")
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func twitterSignCallBack(ts *TwitterSrv, w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")
	ctx := context.Background()
	token, err := ts.oauth2Config.Exchange(ctx, code)
	if err != nil {
		return
	}

	client := ts.oauth2Config.Client(context.Background(), token)
	response, err := client.Get("https://api.twitter.com/2/users/me")
	if err != nil {
		return
	}
	defer response.Body.Close()
	fmt.Println(w, response)
}
