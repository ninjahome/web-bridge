package server

import "github.com/ninjahome/web-bridge/server/database"

type TwAPIResponse struct {
	TwitterData *database.TWUserInfo `json:"data"`
	EthAddr     string               `json:"eth_addr"`
	SignUpAt    int64                `json:"sign_up_at"`
}

type TweetResponse struct {
	Data TweetPostResult `json:"data"`
}

type TweetPostResult struct {
	ID   string `json:"id"`
	Text string `json:"text"`
}

type TweetRequest struct {
	Text                  string `json:"text,omitempty"`
	DirectMessageDeepLink string `json:"direct_message_deep_link,omitempty"`
	ForSuperFollowersOnly bool   `json:"for_super_followers_only,omitempty"`
	Geo                   *Geo   `json:"geo,omitempty"`
	Media                 *Media `json:"media,omitempty"`
	Poll                  *Poll  `json:"poll,omitempty"`
	QuoteTweetID          string `json:"quote_tweet_id,omitempty"`
	Reply                 *Reply `json:"reply,omitempty"`
	ReplySettings         string `json:"reply_settings,omitempty"`
}

type Geo struct {
	PlaceID string `json:"place_id,omitempty"`
}

type Media struct {
	MediaIDs      []string `json:"media_ids,omitempty"`
	TaggedUserIDs []string `json:"tagged_user_ids,omitempty"`
}

type Poll struct {
	DurationMinutes int      `json:"duration_minutes,omitempty"`
	Options         []string `json:"options,omitempty"`
}

type Reply struct {
	ExcludeReplyUserIDs []string `json:"exclude_reply_user_ids,omitempty"`
	InReplyToTweetID    string   `json:"in_reply_to_tweet_id,omitempty"`
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
