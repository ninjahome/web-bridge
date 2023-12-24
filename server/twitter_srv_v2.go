package server

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/ninjahome/web-bridge/util"
	"golang.org/x/oauth2"
	"io"
	"net/http"
)

const (
	accessPointPostTweets = "https://api.twitter.com/2/tweets"
	tweetTimeFormat       = "01/02/06 15:04:05"
)

func twitterGetWithAccessToken(token *oauth2.Token, accUrl string, result any) error {
	client := _globalCfg.twOauthCfg.Client(context.Background(), token)
	response, err := client.Get(accUrl)
	if err != nil {
		util.LogInst().Err(err).Msgf("create client err")
		return err
	}
	defer response.Body.Close()

	if err := json.NewDecoder(response.Body).Decode(result); err != nil {
		util.LogInst().Err(err).Msgf("parse twitter call back data  err:%s", err)
		return err
	}
	return nil
}

func twitterPostWithAccessToken(token *oauth2.Token, accUrl string, param any, result any) error {
	client := _globalCfg.twOauthCfg.Client(context.Background(), token)

	jsonContent, err := json.Marshal(param)
	if err != nil {
		util.LogInst().Err(err).Msg("Error marshalling tweet content")
		return err
	}

	req, err := http.NewRequest("POST", accUrl, bytes.NewBuffer(jsonContent))
	if err != nil {
		util.LogInst().Err(err).Msg("Error creating POST request")
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		util.LogInst().Err(err).Msg("Error sending POST request")
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		bts, _ := io.ReadAll(resp.Body)
		util.LogInst().Err(err).Msgf("Twitter API responded with status: %s %s", resp.Status, string(bts))
		return fmt.Errorf("twitter API responded with status: %s", resp.Status)
	}

	if err := json.NewDecoder(resp.Body).Decode(result); err != nil {
		util.LogInst().Err(err).Msgf("parse twitter call back data  err:%s", err)
		return err
	}
	return nil
}

type TwitterPostResponse struct {
	Data TweetPostResult `json:"data"`
}

type TweetPostResult struct {
	ID   string `json:"id"`
	Text string `json:"text"`
}
type TweetContent struct {
	Txt string `json:"text"`
}

func postTweetsV2(w http.ResponseWriter, r *http.Request) {
	var ut, errToken = checkTwitterRightsV2(w, r)
	if errToken != nil {
		util.LogInst().Err(errToken).Msg("load access token failed")
		http.Error(w, errToken.Error(), http.StatusInternalServerError)
		return
	}
	param := &SignDataByEth{}
	err := util.ReadRequest(r, param)
	if err != nil {
		util.LogInst().Err(err).Msg("Error parsing sign data ")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var tweetContent NinjaTweet
	err = json.Unmarshal([]byte(param.Message), &tweetContent)
	if err != nil {
		util.LogInst().Err(err).Msg("Error parsing tweet ")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if !tweetContent.IsValid() {
		util.LogInst().Warn().Msg("invalid tweet content:" + tweetContent.String())
		http.Error(w, "invalid tweet content", http.StatusBadRequest)
		return
	}

	err = util.Verify(tweetContent.Web3ID, param.Message, param.Signature)
	if err != nil {
		util.LogInst().Err(err).Msg("tweet signature verify failed")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	tweetContent.Signature = param.Signature

	var tweetResponse = &TwitterPostResponse{}
	var tweetReq = &TweetContent{Txt: tweetContent.ToTweet()}
	err = twitterPostWithAccessToken(ut.Token, accessPointPostTweets, tweetReq, tweetResponse)
	if err != nil {
		util.LogInst().Err(err).Msg("post tweet failed")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	tweetContent.TweetId = tweetResponse.Data.ID

	err = DbInst().SaveTweet(&tweetContent)
	if err != nil {
		util.LogInst().Err(err).Msg("save tweet failed")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	bts, _ := json.Marshal(tweetContent)
	w.Write(bts)
	util.LogInst().Debug().Str("eth-addr", tweetContent.Web3ID).
		Str("twitter-id", tweetContent.TweetId).
		Str("tweet-id", tweetContent.TweetId).Msg("Tweet posted successfully")
}

func checkTwitterRightsV2(w http.ResponseWriter, r *http.Request) (*TwUserAccessTokenV2, error) {
	var ninjaUsr, err = validateUsrRights(r)
	if err != nil {
		http.Redirect(w, r, "/signIn", http.StatusFound)
		return nil, err
	}

	var twitterUid = ninjaUsr.TwID
	if len(twitterUid) == 0 {
		util.LogInst().Warn().Msg("no twitter id for ninja user:" + ninjaUsr.EthAddr)
		return nil, fmt.Errorf("bind twitter first")
	}
	var token, errToken = getAccessTokenFromSessionV2(r)
	if errToken == nil {
		return &TwUserAccessTokenV2{
			twitterUid, token,
		}, nil
	}
	ut, err := DbInst().GetTwAccessTokenV2(twitterUid)
	if err != nil {
		util.LogInst().Err(err).Str("twitter-id", twitterUid).Msg("access token not in db")
		return nil, err
	}
	if false == isOAuth2TokenValid(ut.Token) {
		ut.Token, err = refreshAccessToken(ut.RefreshToken)
		if err != nil {
			util.LogInst().Err(err).Str("twitter-id", twitterUid).Msg("refresh token failed")
			return nil, err
		}
		err = DbInst().SaveTwAccessTokenV2(ut)
		if err != nil {
			util.LogInst().Err(err).Str("twitter-id", twitterUid).Msg("save refreshed token failed")
			return nil, err
		}
	}
	return ut, nil
}

func getAccessTokenFromSessionV2(r *http.Request) (*oauth2.Token, error) {
	bts, err := SMInst().Get(sesKeyForAccessTokenV2, r)
	if err != nil {
		return nil, err
	}
	var token oauth2.Token
	err = json.Unmarshal(bts.([]byte), &token)
	if err != nil {
		return nil, err
	}
	return &token, nil
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

func isOAuth2TokenValid(token *oauth2.Token) bool {
	client := _globalCfg.twOauthCfg.Client(context.Background(), token)

	resp, err := client.Get(accessUserURLV2)
	if err != nil {
		return false
	}
	defer resp.Body.Close()

	return resp.StatusCode == http.StatusOK
}
