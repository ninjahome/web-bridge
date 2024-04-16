package server

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/golang/freetype"
	"github.com/golang/freetype/truetype"
	"github.com/ninjahome/web-bridge/blockchain"
	"github.com/ninjahome/web-bridge/blockchain/ethapi"
	"github.com/ninjahome/web-bridge/database"
	"github.com/ninjahome/web-bridge/util"
	"golang.org/x/oauth2"
	"html/template"
	"math/big"
	"net/http"
	"os"
)

var (
	cfgActionRouter = map[string]LogicAction{
		"/signUpByTwitter":           {signUpByTwitterV1, true},
		"/tw_callback":               {twitterSignCallBackV1, false},
		"/signUpSuccessByTw":         {signUpSuccessByTw, false},
		"/signInByEth":               {signInByEth, false},
		"/bindWeb3ID":                {bindingWeb3ID, true},
		"/queryTwBasicById":          {queryTwBasicById, true},
		"/queryTwBasicByTweetHash":   {queryTwBasicByTweetHash, true},
		"/queryNjBasicByID":          {queryNjBasicByID, true},
		"/signOut":                   {signOut, false},
		"/main":                      {mainPage, true},
		"/refreshNjUser":             {refreshNjUser, true},
		"/lotteryGame":               {showLotteryMain, true},
		"/kolKey":                    {showKolKeyPage, true},
		"/postTweet":                 {postTweets, true},
		"/updateTweetPaymentStatus":  {updateTweetTxStatus, true},
		"/reloadPaymentDetails":      {queryTweetDetails, false},
		"/updateTweetVoteStatus":     {updateTweetVoteStatus, true},
		"/updatePointsForSingleBets": {updatePointsForSingleBets, true},
		"/shareVoteAction":           {shareVoteAction, true},
		"/buyRights":                 {mainPage, true},
		"/buyFromShare":              {mainPage, true},
		"/tweetQuery":                {globalTweetQuery, true},
		"/votedTweetIds":             {votedTweetsQuery, true},
		"/removeUnpaidTweet":         {removeUnpaidTweet, true},
		"/mostVotedTweet":            {mostVotedTweet, true},
		"/mostVotedKol":              {mostVotedKol, true},
		"/queryTweetByHash":          {queryTweetByHash, true},
		"/queryWinHistory":           {queryWinHistory, true},
		"/searchTwitterUsr":          {searchTwitterUsr, true},
		"/tweetImgRaw":               {tweetImgRaw, false},
	}

	cfgHtmlFileRouter = map[string]string{
		"/signIn": "html/signIn.html",
		"/":       "html/index.html",
		"/app":    "html/signIn.html",
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

type SysConf struct {
	LogLevel string `json:"log_level"`
	UrlHome  string `json:"url_home"`
	HttpPort string `json:"http_port"`
	*HttpConf
	*TwitterConf
	*database.FileStoreConf
	twOauthCfg *oauth2.Config
	*blockchain.BCConf
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
	s += "\n" + c.BCConf.String()
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

	blockchain.InitConf(c.BCConf)
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

const (
	NjTweetID = "NjTID"
	SharedID  = "sharedID"
	SharedUsr = "shareUsr"
)

func (c *SysConf) getContractObj() (*ethapi.TweetLotteryGame, error) {
	cli, err := ethclient.Dial(c.InfuraUrl)
	if err != nil {
		util.LogInst().Err(err).Msg("dial eth failed")
		return nil, err
	}

	defer cli.Close()

	contractAddress := common.HexToAddress(c.GameContract)
	game, err := ethapi.NewTweetLotteryGame(contractAddress, cli)
	if err != nil {
		util.LogInst().Err(err).Str("contract-address", c.GameContract).Msg("failed create game obj")
		return nil, err
	}
	return game, nil
}

func (c *SysConf) getHistoryBonus() (*big.Float, error) {
	game, err := c.getContractObj()
	if err != nil {
		util.LogInst().Err(err).Msg("dial up to  block chain failed")
		return nil, err
	}

	result, err := game.SystemSettings(nil)

	if err != nil {
		util.LogInst().Err(err).Msg("query game system setting failed")
		return nil, err
	}

	weiToEth := new(big.Float).SetInt(big.NewInt(1e18))
	bonusEth := new(big.Float).Quo(new(big.Float).SetInt(result.TBonus), weiToEth)

	return bonusEth, nil
}

func (c *SysConf) GetNjVoteAd(NjTwID int64, web3Id, slogan string) string {
	return fmt.Sprintf("\r\n"+slogan+c.UrlHome+"/buyFromShare?"+SharedID+"=%d&&"+SharedUsr+"=%s", NjTwID, web3Id)
}

func (c *SysConf) GetNjProtocolAd(NjTwID int64, slogan string) string {
	return fmt.Sprintf("\r\n"+slogan+c.UrlHome+"/buyRights?"+NjTweetID+"=%d", NjTwID)
}
