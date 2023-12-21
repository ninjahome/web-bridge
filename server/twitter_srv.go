package server

import (
	"encoding/json"
	"github.com/dghubble/oauth1"
	"github.com/ninjahome/web-bridge/util"
	"net/http"
)

func pullTwitterTimeline(w http.ResponseWriter, r *http.Request) {
	var twitterUid = r.URL.Query().Get("twitter_uid")
	if len(twitterUid) == 0 {
		util.LogInst().Warn().Msg("missing twitter_uid when pull twitter timeline")
		http.Error(w, "missing twitter_uid when pull twitter timeline", http.StatusInternalServerError)
		return
	}

	var ut, err = getAccessTokenFromSession(r)
	if err != nil {
		ut, err = DbInst().GetTwAccessToken(twitterUid)
		if err != nil {
			util.LogInst().Err(err).Str("twitter-id", twitterUid).Msg("access token not in db")
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	config := oauth1.NewConfig(_globalCfg.ConsumerKey, _globalCfg.ConsumerSecret)
	httpClient := config.Client(oauth1.NoContext, ut.GetToken())

	timelineURL := "https://api.twitter.com/1.1/statuses/user_timeline.json?exclude_replies=true&include_rts=false&count=10"

	resp, err := httpClient.Get(timelineURL)
	if err != nil {
		util.LogInst().Err(err).Str("twitter-id", twitterUid).Msg("http pull time failed")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		util.LogInst().Warn().Str("twitter-id", twitterUid).Msg("pull time status:" + resp.Status)
		http.Error(w, "pull time status:"+resp.Status, http.StatusInternalServerError)
		return
	}

	var tweets []Tweet
	if err := json.NewDecoder(resp.Body).Decode(&tweets); err != nil {
		util.LogInst().Err(err).Str("twitter-id", twitterUid).Msg("parse pull time failed")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	bts, _ := json.Marshal(tweets)
	w.Write(bts)
}

type Tweet struct {
	CreatedAt        string `json:"created_at"`        // 推文的时间
	IDStr            string `json:"id_str"`            // 推文的唯一ID字符串
	Text             string `json:"text"`              // 推文的内容
	User             User   `json:"user"`              // 发送推文的用户信息
	ExtendedEntities Media  `json:"extended_entities"` // 推文中的媒体实体（图片和视频）
}

type User struct {
	IDStr                string `json:"id_str"`                  // 用户的唯一ID字符串
	ScreenName           string `json:"screen_name"`             // 用户的屏幕名（用户名）
	ProfileImageUrlHttps string `json:"profile_image_url_https"` // 用户头像的URL（HTTPS）
}

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
