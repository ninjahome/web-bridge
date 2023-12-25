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
	Web3IDProfile          = "Ninja Protocol Web3 ID:"
	sesKeyForNjUserId      = "twitter-signup-ninja-user-id"
	sesKeyForAccessTokenV1 = "twitter-access-key-v1"
	sesKeyForRequestSecret = "ses-key-for-request-secret"
	accessUserProUrl       = "https://api.twitter.com/1.1/account/update_profile.json"
	accessReqTokenURL      = "https://api.twitter.com/oauth/request_token"
	accessOauthTokenURL    = "https://api.twitter.com/oauth/authorize?oauth_token=%s"
	accessAccessTokenURL   = "https://api.twitter.com/oauth/access_token"
	verifyCredentialsURL   = "https://api.twitter.com/1.1/account/verify_credentials.json?skip_status=true"
)

func parseUserToken(values url.Values) *TwUserAccessToken {
	accessToken := values.Get("oauth_token")
	accessSecret := values.Get("oauth_token_secret")
	userID := values.Get("user_id")
	screenName := values.Get("screen_name")
	return &TwUserAccessToken{
		OauthToken:       accessToken,
		OauthTokenSecret: accessSecret,
		UserId:           userID,
		ScreenName:       screenName,
	}
}

func (ut *TwUserAccessToken) GetToken() *oauth1.Token {
	return &oauth1.Token{
		Token:       ut.OauthToken,
		TokenSecret: ut.OauthTokenSecret,
	}
}

func (ut *TwUserAccessToken) String() string {
	bts, _ := json.Marshal(ut)
	return string(bts)
}

func getAccessTokenFromSession(r *http.Request) (*TwUserAccessToken, error) {
	bts, err := SMInst().Get(sesKeyForAccessTokenV1, r)
	if err != nil {
		return nil, err
	}
	var token TwUserAccessToken
	err = json.Unmarshal([]byte(bts.(string)), &token)
	if err != nil {
		return nil, err
	}
	return &token, nil
}

func signUpByTwitterV1(w http.ResponseWriter, r *http.Request) {
	ethAddr := r.URL.Query().Get("eth_addr")
	if ethAddr == "" {
		http.Error(w, "eth_addr parameter is required", http.StatusBadRequest)
		return
	}

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
	values, err := url.ParseQuery(bodyString)
	if err != nil {
		util.LogInst().Err(err).Msg("Failed to parse query from response body")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	requestToken := values.Get("oauth_token")
	requestSecret := values.Get("oauth_token_secret")

	var err1, err2 = SMInst().Set(r, w, sesKeyForRequestSecret, requestSecret), SMInst().Set(r, w, sesKeyForNjUserId, ethAddr)
	if err1 != nil || err2 != nil {
		util.LogInst().Err(err1).Err(err2).Msg("save secret or eth id to session failed")
		http.Error(w, "session save failed", http.StatusInternalServerError)
		return
	}
	authorizeURL := fmt.Sprintf(accessOauthTokenURL, requestToken)
	http.Redirect(w, r, authorizeURL, http.StatusTemporaryRedirect)
}

func twitterSignCallBackV1(w http.ResponseWriter, r *http.Request) {
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
	values, err := url.ParseQuery(bodyString)
	if err != nil {
		util.LogInst().Err(err).Msg("Failed to parse response body")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	token := parseUserToken(values)
	_ = SMInst().Set(r, w, sesKeyForAccessTokenV1, token.String())
	err = DbInst().SaveTwAccessToken(token)
	if err != nil {
		util.LogInst().Err(err).Msg("save twitter user access token failed")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/signUpSuccessByTw", http.StatusFound)
}

func signUpSuccessByTw(w http.ResponseWriter, r *http.Request) {
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
		TwitterData: &TWUserInfo{
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

func updateTwitterBio(w http.ResponseWriter, r *http.Request) error {

	ut, err := getAccessTokenFromSession(r)
	if err != nil {
		util.LogInst().Err(err).Msg("no user access token found")
	}
	config := oauth1.NewConfig(_globalCfg.ConsumerKey, _globalCfg.ConsumerSecret)
	httpClient := config.Client(oauth1.NoContext, ut.GetToken())

	values := url.Values{}
	values.Add("description", Web3IDProfile)
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

func fetchTwitterUserInfo(ut *TwUserAccessToken) (*TwAPIResponse, error) {
	config := oauth1.NewConfig(_globalCfg.ConsumerKey, _globalCfg.ConsumerSecret)

	util.LogInst().Debug().Msg(ut.String())

	httpClient := config.Client(oauth1.NoContext, ut.GetToken())

	//userInfoURL := fmt.Sprintf("https://api.twitter.com/1.1/users/show.json?screen_name=%s", ut.ScreenName)
	userInfoURL := fmt.Sprintf("https://api.twitter.com/1.1/users/show.json?user_id=%s", ut.UserId)

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

func verifyTwitterCredentials(ut *TwUserAccessToken) (*VerifiedTwitterUser, error) {
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

	return &verifiedUser, nil
}

type VerifiedTwitterUser struct {
	ContributorsEnabled            bool   `json:"contributors_enabled"`
	CreatedAt                      string `json:"created_at"`
	DefaultProfile                 bool   `json:"default_profile"`
	DefaultProfileImage            bool   `json:"default_profile_image"`
	Description                    string `json:"description"`
	FavouritesCount                int    `json:"favourites_count"`
	FollowersCount                 int    `json:"followers_count"`
	FriendsCount                   int    `json:"friends_count"`
	GeoEnabled                     bool   `json:"geo_enabled"`
	ID                             int64  `json:"id"`
	IDStr                          string `json:"id_str"`
	IsTranslator                   bool   `json:"is_translator"`
	Lang                           string `json:"lang"`
	ListedCount                    int    `json:"listed_count"`
	Location                       string `json:"location"`
	Name                           string `json:"name"`
	ProfileBackgroundColor         string `json:"profile_background_color"`
	ProfileBackgroundImageUrl      string `json:"profile_background_image_url"`
	ProfileBackgroundImageUrlHttps string `json:"profile_background_image_url_https"`
	ProfileBackgroundTile          bool   `json:"profile_background_tile"`
	ProfileImageUrl                string `json:"profile_image_url"`
	ProfileImageUrlHttps           string `json:"profile_image_url_https"`
	ProfileLinkColor               string `json:"profile_link_color"`
	ProfileSidebarBorderColor      string `json:"profile_sidebar_border_color"`
	ProfileSidebarFillColor        string `json:"profile_sidebar_fill_color"`
	ProfileTextColor               string `json:"profile_text_color"`
	ProfileUseBackgroundImage      bool   `json:"profile_use_background_image"`
	Protected                      bool   `json:"protected"`
	ScreenName                     string `json:"screen_name"`
	ShowAllInlineMedia             bool   `json:"show_all_inline_media"`
	StatusesCount                  int    `json:"statuses_count"`
	TimeZone                       string `json:"time_zone"`
	URL                            string `json:"url"`
	UtcOffset                      int    `json:"utc_offset"`
	Verified                       bool   `json:"verified"`
}

type Web3BindingData struct {
	EthAddr  string `json:"eth_addr"`
	TwID     string `json:"tw_id"`
	BindTime int64  `json:"bind_time"`
}
