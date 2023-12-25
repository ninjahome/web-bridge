package server

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
