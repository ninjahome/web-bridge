package server

import (
	"encoding/json"
	"fmt"
	"github.com/dghubble/oauth1"
	"github.com/ninjahome/web-bridge/util"
	"io"
	"net/http"
	"net/url"
)

type Media struct {
	Media []MediaInfo `json:"media"` // 媒体信息数组
}

type MediaInfo struct {
	MediaURLHttps string `json:"media_url_https"`      // 媒体内容的HTTPS URL
	Type          string `json:"type"`                 // 媒体类型（例如："photo" 或 "video"）
	VideoInfo     Video  `json:"video_info,omitempty"` // 视频信息，仅在类型为视频时出现
}

type Video struct {
	Variants []VideoVariant `json:"variants"` // 视频变体信息，例如不同的分辨率
}

type VideoVariant struct {
	ContentType string `json:"content_type"` // 内容类型，例如 "video/mp4"
	URL         string `json:"url"`          // 视频的URL
}

type TweetContent struct {
	TweetContent string `json:"tweet_content"`
}

type Tweet struct {
	CreatedAt        string `json:"created_at"`
	IDStr            string `json:"id_str"`
	Text             string `json:"text"`
	User             User   `json:"user"`
	ExtendedEntities Media  `json:"extended_entities"`
}

type User struct {
	ID                   int64  `json:"id"`
	IDStr                string `json:"id_str"`
	Name                 string `json:"name"`
	ScreenName           string `json:"screen_name"`
	ProfileImageUrlHttps string `json:"profile_image_url_https"`
}

func checkTwitterRights(w http.ResponseWriter, r *http.Request) (*TwUserAccessToken, error) {
	var ninjaUsr, err = validateUsrRights(w, r)
	if err != nil {
		http.Redirect(w, r, "/signIn", http.StatusFound)
		return nil, err
	}

	var twitterUid = ninjaUsr.TwID
	if len(twitterUid) == 0 {
		util.LogInst().Warn().Msg("no twitter id for ninja user:" + ninjaUsr.EthAddr)
		return nil, fmt.Errorf("bind twitter first")
	}
	var ut, errToken = getAccessTokenFromSession(r)
	if errToken == nil {
		return ut, nil
	}
	ut, err = DbInst().GetTwAccessToken(twitterUid)
	if err != nil {
		util.LogInst().Err(err).Str("twitter-id", twitterUid).Msg("access token not in db")
		return nil, err
	}
	return ut, nil
}

func pullTwitterTimeline(w http.ResponseWriter, r *http.Request) {

	var ut, errToken = checkTwitterRights(w, r)
	if errToken != nil {
		util.LogInst().Err(errToken).Msg("load access  token failed")
		http.Error(w, errToken.Error(), http.StatusInternalServerError)
		return
	}

	config := oauth1.NewConfig(_globalCfg.ConsumerKey, _globalCfg.ConsumerSecret)
	httpClient := config.Client(oauth1.NoContext, ut.GetToken())

	timelineURL := "https://api.twitter.com/1.1/statuses/user_timeline.json?exclude_replies=true&include_rts=false&count=10"

	resp, errResp := httpClient.Get(timelineURL)
	if errResp != nil {
		util.LogInst().Err(errResp).Str("twitter-id", ut.UserId).Msg("http pull time failed")
		http.Error(w, errResp.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bts, _ := io.ReadAll(resp.Body)
		util.LogInst().Warn().Str("twitter-id", ut.UserId).Msg("pull time status:" + resp.Status + string(bts))
		http.Error(w, "pull time status:"+resp.Status, http.StatusInternalServerError)
		return
	}

	var tweets []Tweet
	if err := json.NewDecoder(resp.Body).Decode(&tweets); err != nil {
		util.LogInst().Err(err).Str("twitter-id", ut.UserId).Msg("parse pull time failed")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	bts, _ := json.Marshal(tweets)
	w.Write(bts)
}

// postTweets 发布推文到Twitter
func postTweets(w http.ResponseWriter, r *http.Request) {
	var ut, errToken = checkTwitterRights(w, r)
	if errToken != nil {
		util.LogInst().Err(errToken).Msg("load access  token failed")
		http.Error(w, errToken.Error(), http.StatusInternalServerError)
		return
	}
	var tweetContent TweetContent
	err := util.ReadRequest(r, &tweetContent)
	if err != nil {
		util.LogInst().Err(err).Msg("Error parsing request body")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if tweetContent.TweetContent == "" {
		util.LogInst().Warn().Msg("Tweet text cannot be empty")
		http.Error(w, "Tweet text cannot be empty", http.StatusBadRequest)
		return
	}

	config := oauth1.NewConfig(_globalCfg.ConsumerKey, _globalCfg.ConsumerSecret)
	token := oauth1.NewToken(ut.OauthToken, ut.OauthTokenSecret)
	httpClient := config.Client(oauth1.NoContext, token)

	updateURL := "https://api.twitter.com/1.1/statuses/update.json"

	// 构建请求参数
	params := url.Values{}
	params.Set("status", tweetContent.TweetContent)

	// 发送POST请求到Twitter API
	resp, err := httpClient.PostForm(updateURL, params)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bts, _ := io.ReadAll(resp.Body)
		util.LogInst().Warn().Msg("post tweet failed" + string(bts))
		http.Error(w, fmt.Sprintf("Twitter API responded with status: %s", resp.Status), http.StatusInternalServerError)
		return
	}

	var tweetResponse Tweet
	if err := json.NewDecoder(resp.Body).Decode(&tweetResponse); err != nil {
		http.Error(w, "Error decoding response from Twitter API", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(tweetResponse)
	util.LogInst().Debug().Msg("Tweet posted successfully")
}

func pullTwitterHomeTimeline(w http.ResponseWriter, r *http.Request) {
	// 检查Twitter权限
	var ut, errToken = checkTwitterRights(w, r)
	if errToken != nil {
		util.LogInst().Err(errToken).Msg("load access token failed")
		http.Error(w, errToken.Error(), http.StatusInternalServerError)
		return
	}

	// 设置Twitter API配置
	config := oauth1.NewConfig(_globalCfg.ConsumerKey, _globalCfg.ConsumerSecret)
	httpClient := config.Client(oauth1.NoContext, ut.GetToken())

	// 构建请求的URL
	timelineURL := "https://api.twitter.com/1.1/statuses/home_timeline.json?exclude_replies=true&include_rts=false&count=10"

	// 发送请求
	resp, errResp := httpClient.Get(timelineURL)
	if errResp != nil {
		util.LogInst().Err(errResp).Str("twitter-id", ut.UserId).Msg("http pull timeline failed")
		http.Error(w, errResp.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	// 检查响应状态码
	if resp.StatusCode != http.StatusOK {
		bts, _ := io.ReadAll(resp.Body)
		util.LogInst().Warn().Str("twitter-id", ut.UserId).Msg("pull timeline status:" + resp.Status + string(bts))
		http.Error(w, "pull timeline status:"+resp.Status, http.StatusInternalServerError)
		return
	}

	// 解析响应数据
	var tweets []Tweet
	if err := json.NewDecoder(resp.Body).Decode(&tweets); err != nil {
		util.LogInst().Err(err).Str("twitter-id", ut.UserId).Msg("parse pull timeline failed")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// 发送响应
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	bts, _ := json.Marshal(tweets)
	w.Write(bts)
}
