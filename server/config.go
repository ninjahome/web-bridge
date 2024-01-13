package server

import (
	"fmt"
	"github.com/golang/freetype"
	"github.com/golang/freetype/truetype"
	"github.com/ninjahome/web-bridge/database"
	"github.com/ninjahome/web-bridge/util"
	"golang.org/x/oauth2"
	"html/template"
	"net/http"
	"os"
)

var (
	cfgActionRouter = map[string]LogicAction{
		"/signUpByTwitter":          {signUpByTwitterV1, true},
		"/tw_callback":              {twitterSignCallBackV1, false},
		"/signUpSuccessByTw":        {signUpSuccessByTw, false},
		"/signInByEth":              {signInByEth, false},
		"/bindWeb3ID":               {bindingWeb3ID, true},
		"/queryTwBasicById":         {queryTwBasicById, true},
		"/queryTwBasicByTweetHash":  {queryTwBasicByTweetHash, true},
		"/queryNjBasicByID":         {queryNjBasicByID, true},
		"/signOut":                  {signOut, false},
		"/main":                     {mainPage, true},
		"/lotteryGame":              {showLotteryMain, true},
		"/kolKey":                   {showKolKeyPage, true},
		"/postTweet":                {postTweets, true},
		"/updateTweetPaymentStatus": {updateTweetTxStatus, true},
		"/reloadPaymentDetails":     {queryTweetDetails, false},
		"/updateTweetVoteStatus":    {updateTweetVoteStatus, true},
		"/shareVoteAction":          {shareVoteAction, true},
		"/buyRights":                {mainPage, true},
		"/buyFromShare":             {mainPage, true},
		"/tweetQuery":               {globalTweetQuery, true},
		"/votedTweetIds":            {votedTweetsQuery, true},
		"/removeUnpaidTweet":        {removeUnpaidTweet, true},
		"/mostVotedTweet":           {mostVotedTweet, true},
	}

	cfgHtmlFileRouter = map[string]string{
		"/signIn": "html/signIn.html",
		"/":       "html/signIn.html",
	}

	_globalCfg *SysConf
)

type LogicAction struct {
	Action    func(w http.ResponseWriter, r *http.Request, token *database.NinjaUsrInfo)
	NeedToken bool
}

type HttpConf struct {
	RefreshContent      bool   `json:"refresh_content"`
	UseHttps            bool   `json:"use_https"`
	SSLCertFile         string `json:"ssl_cert_file"`
	SSLKeyFile          string `json:"ssl_key_file"`
	SessionKey          string `json:"session_key"`
	htmlTemplateManager *template.Template
}

func (c *HttpConf) String() string {
	s := "\n------server config------"
	s += "\nrefresh content:" + fmt.Sprintf("%t", c.RefreshContent)
	s += "\nuse https:" + fmt.Sprintf("%t", c.UseHttps)
	s += "\nssl cert file:" + c.SSLCertFile
	s += "\nssl key file:" + c.SSLKeyFile
	s += "\n-------------------------"
	return s
}

type TwitterConf struct {
	imgFont        *truetype.Font
	FontSize       float64 `json:"font_size"`
	FontPath       string  `json:"font_path"`
	MaxTxtPerImg   int     `json:"max_txt_per_img"`
	ClientID       string  `json:"client_id"`
	ClientSecret   string  `json:"client_secret"`
	ConsumerKey    string  `json:"consumer_key"`
	ConsumerSecret string  `json:"consumer_secret"`
}

func (c *TwitterConf) String() string {
	s := "\n------twitter config------"
	s += "\nclient id:" + c.ClientID
	s += "\nfont path:" + c.FontPath
	s += "\nfont size:" + fmt.Sprintf("%.1f", c.FontSize)
	s += "\n--------------------------"
	return s
}

type BlockChainConf struct {
	TweeTVoteContractAddress string `json:"tweet_vote_contract_address"`
	GameContract             string `json:"game_plugin_contract_address"`
	KolKeyContractAddress    string `json:"kol_key_contract_address"`
	InfuraUrl                string `json:"infura_url"`
	GameTimeInMinute         int    `json:"game_time_in_minute,omitempty"`
	TxCheckerInSeconds       int    `json:"tx_checker_in_seconds,omitempty"`
	ChainID                  int64  `json:"chain_id,omitempty"`
}

func (c *BlockChainConf) String() string {
	s := "\n------block chain config------"
	s += "\ntweet vote:" + c.TweeTVoteContractAddress
	s += "\ngame:" + c.GameContract
	s += "\nkol key:" + c.KolKeyContractAddress
	s += "\ninfura url:" + c.InfuraUrl
	s += "\ngame check time(minutes):" + fmt.Sprintf("%d", c.GameTimeInMinute)
	s += "\ntransaction check time(seconds):" + fmt.Sprintf("%d", c.TxCheckerInSeconds)
	s += "\nchain id:" + fmt.Sprintf("%d", c.ChainID)
	s += "\n--------------------------"
	return s
}

type SysConf struct {
	LogLevel string `json:"log_level"`
	UrlHome  string `json:"url_home"`

	Slogan     string `json:"slogan"`
	VoteSlogan string `json:"vote_slogan"`

	HttpPort string `json:"http_port"`
	*HttpConf
	*TwitterConf
	*database.FileStoreConf
	twOauthCfg *oauth2.Config
	*BlockChainConf
}

func (c *SysConf) String() any {
	var s = "\n=======================system config==========================="
	s += "\nlog level:" + c.LogLevel
	s += "\nlocal mode:" + fmt.Sprintf("%t", c.LocalRun)
	s += "\nhome:" + c.UrlHome
	s += "\nhttp port:" + c.HttpPort
	s += "\n" + c.HttpConf.String()
	s += "\n" + c.TwitterConf.String()
	s += "\n" + c.FileStoreConf.String()
	s += "\n" + c.BlockChainConf.String()
	s += "\n=============================================================="
	return s
}

var (
	twitterSignUpCallbackURL = ""
)

func InitConf(c *SysConf) {
	util.SetLogLevel(c.LogLevel)
	if len(c.HttpPort) == 0 {
		c.HttpPort = "80"
	}
	fmt.Println(c.String())

	_globalCfg = c
	database.InitConf(c.FileStoreConf)

	twitterSignUpCallbackURL = _globalCfg.UrlHome + "/tw_callback"
	conf := _globalCfg.TwitterConf
	var oauth2Config = &oauth2.Config{
		RedirectURL:  twitterSignUpCallbackURL,
		ClientID:     conf.ClientID,
		ClientSecret: conf.ClientSecret,
		Scopes:       []string{"tweet.read", "tweet.write", "follows.read", "follows.write", "users.read", "offline.access"},
		Endpoint: oauth2.Endpoint{
			AuthURL:  authorizeURLV2,
			TokenURL: accessTokenURLV2,
		},
	}
	_globalCfg.twOauthCfg = oauth2Config

	_globalCfg.htmlTemplateManager = util.ParseTemplates("assets/html")

	fontBytes, err := os.ReadFile(_globalCfg.FontPath)
	if err != nil {
		panic(err)
	}
	f, err := freetype.ParseFont(fontBytes)
	if err != nil {
		panic(err)
	}

	_globalCfg.imgFont = f
}

func (c *SysConf) GetNjProtocolAd(NjTwID int64) string {
	return fmt.Sprintf("\n%s:%s/buyRights?id=%d", c.Slogan, c.UrlHome, NjTwID)
}

func (c *SysConf) GetNjVoteAd(NjTwID int64, voteCount int, web3Id string) string {
	return fmt.Sprintf("\n%s [%d] bets:%s/buyFromShare?sharedID=%d&&shareUsr=%s", c.VoteSlogan, voteCount, c.UrlHome, NjTwID, web3Id)
}
