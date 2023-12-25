package server

import (
	"fmt"
	"github.com/golang/freetype"
	"github.com/golang/freetype/truetype"
	"github.com/ninjahome/web-bridge/util"
	"golang.org/x/oauth2"
	"html/template"
	"net/http"
	"os"
)

var (
	cfgActionRouter = map[string]LogicAction{
		"/signUpByTwitter":    {signUpByTwitterV1, true},
		"/tw_callback":        {twitterSignCallBackV1, false},
		"/signUpSuccessByTw":  {signUpSuccessByTw, false},
		"/signInByEth":        {signInByEth, false},
		"/bindWeb3ID":         {bindingWeb3ID, true},
		"/queryTwBasicById":   {queryTwBasicById, true},
		"/signOut":            {signOut, false},
		"/main":               {mainPage, true},
		"/postTweet":          {postTweets, true},
		"/buyRights":          {mainPage, true},
		"/globalLatestTweets": {globalTweetQuery, true},
	}

	cfgHtmlFileRouter = map[string]string{
		"/signIn": "html/signIn.html",
		"/":       "html/signIn.html",
	}

	_globalCfg *SysConf
)

type LogicAction struct {
	Action    func(w http.ResponseWriter, r *http.Request, token *NinjaUsrInfo)
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
	ClientID       string  `json:"client_id"`
	ClientSecret   string  `json:"client_secret"`
	ConsumerKey    string  `json:"consumer_key"`
	ConsumerSecret string  `json:"consumer_secret"`
}

func (c *TwitterConf) String() string {
	s := "\n------twitter config------"
	s += "\nclient id:" + c.ClientID
	s += "\n--------------------------"
	return s
}

type FileStoreConf struct {
	ProjectID      string `json:"project_id"`
	KeyFilePath    string `json:"key_file_path"`
	TweetsPageSize int    `json:"tweets_page_size"`
}

func (c *FileStoreConf) String() string {
	s := "\n------twitter config------"
	s += "\nproject id:" + c.ProjectID
	s += "\nkey path :" + c.KeyFilePath
	s += "\n--------------------------"
	return s
}

type SysConf struct {
	Log      string `json:"log"`
	LocalRun bool   `json:"local_run"`
	UrlHome  string `json:"url_home"`
	*HttpConf
	*TwitterConf
	*FileStoreConf
	twOauthCfg *oauth2.Config
}

func (c *SysConf) String() any {
	var s = "\n=======================system config==========================="
	s += "\nlog level:" + c.Log
	s += "\nlocal mode:" + fmt.Sprintf("%t", c.LocalRun)
	s += "\nhome:" + c.UrlHome
	s += "\n" + c.HttpConf.String()
	s += "\n" + c.TwitterConf.String()
	s += "\n" + c.FileStoreConf.String()
	s += "\n=============================================================="
	return s
}

var (
	twitterSignUpCallbackURL = ""
)

func InitConf(c *SysConf) {
	_globalCfg = c
	util.SetLogLevel(c.Log)
	fmt.Println(c.String())

	_ = DbInst()

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
	return fmt.Sprintf("\nBuy Rights:%s/buyRights?id=%d", c.UrlHome, NjTwID)
}
