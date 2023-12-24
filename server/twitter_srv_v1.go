package server

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/dghubble/oauth1"
	"github.com/ninjahome/web-bridge/util"
	"image"
	"image/jpeg"
	"io"
	"mime/multipart"
	"net/http"
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

func postTweets(w http.ResponseWriter, r *http.Request) {
	var ut, errToken = checkTwitterRights(w, r)
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
	token := oauth1.NewToken(ut.OauthToken, ut.OauthTokenSecret)

	txtImg, err := util.ConvertLongTweetToImg(tweetContent.TweetContent)
	mediaID, err := uploadMedia(token, txtImg)
	config := oauth1.NewConfig(_globalCfg.ConsumerKey, _globalCfg.ConsumerSecret)
	httpClient := config.Client(oauth1.NoContext, token)

	// Twitter API v2 的 URL
	updateURL := "https://api.twitter.com/2/tweets"

	tweetBody, err := json.Marshal(map[string]interface{}{
		//"text": tweetContent.TweetContent,
		"text": "message from dessage web3 ",
		"media": map[string][]string{
			"media_ids": []string{mediaID},
		},
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// 发送POST请求到 Twitter API
	req, err := http.NewRequest("POST", updateURL, bytes.NewBuffer(tweetBody))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := httpClient.Do(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		bts, _ := io.ReadAll(resp.Body)
		util.LogInst().Warn().Int("status", resp.StatusCode).Msg("post tweet failed" + string(bts))
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

func uploadMedia(token *oauth1.Token, img image.Image) (string, error) {
	var buffer bytes.Buffer
	writer := multipart.NewWriter(&buffer)

	// 创建multipart的图片部分
	part, err := writer.CreateFormFile("media", "image.jpg")
	if err != nil {
		return "", err
	}

	err = jpeg.Encode(part, img, &jpeg.Options{Quality: jpeg.DefaultQuality})
	if err != nil {
		return "", err
	}
	writer.Close()

	// 使用 OAuth 1.0a 进行认证
	config := oauth1.NewConfig(_globalCfg.ConsumerKey, _globalCfg.ConsumerSecret)
	httpClient := config.Client(oauth1.NoContext, token)

	req, err := http.NewRequest("POST", "https://upload.twitter.com/1.1/media/upload.json", &buffer)
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())

	resp, err := httpClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	bts, _ := io.ReadAll(resp.Body)
	util.LogInst().Debug().Msg(string(bts))

	var result map[string]interface{}
	json.NewDecoder(bytes.NewBuffer(bts)).Decode(&result) // 使用 bts 构建新的 buffer，因为原 buffer 已被读取

	mediaID, ok := result["media_id_string"].(string)
	if !ok {
		bts, _ := json.Marshal(result)
		util.LogInst().Warn().Msg("upload media failed:" + string(bts))
		return "", fmt.Errorf("error getting media ID")
	}

	return mediaID, nil
}
