package server

import (
	"fmt"
	"github.com/ninjahome/web-bridge/util"
	"html/template"
	"net/http"
)

var (
	logicRouter = map[string]LogicAction{
		"/signInByTwitter": signInByTwitter,
		"/tw_callback":     twitterSignCallBack,
		"/main":            showMainPage,
	}
	simpleRouterMap = map[string]string{
		"/":         "html/index.html",
		"/index":    "html/index.html",
		"/signPage": "html/sign_twitter.html",
		"/register": "html/create_wallet.html",
	}

	htmlTemplateManager *template.Template //TODO::refactor to a struct
	_globalCfg          *SysConf
)

type LogicAction func(ts *TwitterSrv, w http.ResponseWriter, r *http.Request)

type SrvConf struct {
	DebugMode   bool   `json:"debug_mode"`
	UseHttps    bool   `json:"use_https"`
	SSLCertFile string `json:"ssl_cert_file"`
	SSLKeyFile  string `json:"ssl_key_file"`
	SessionKey  string `json:"session_key"`
}

func (c *SrvConf) String() string {
	s := "\n------server config------"
	s += "\ndebug mode:" + fmt.Sprintf("%t", c.DebugMode)
	s += "\nuse https:" + fmt.Sprintf("%t", c.UseHttps)
	s += "\nssl cert file:" + c.SSLCertFile
	s += "\nssl key file:" + c.SSLKeyFile
	s += "\n-------------------------"
	return s
}

type TwitterConf struct {
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
}

func (c *TwitterConf) String() string {
	s := "\n------twitter config------"
	s += "\nclient id:" + c.ClientID
	s += "\n--------------------------"
	return s
}

type FileStoreConf struct {
	ProjectID   string `json:"project_id"`
	KeyFilePath string `json:"key_file_path"`
}

func (c *FileStoreConf) String() string {
	s := "\n------twitter config------"
	s += "\nproject id:" + c.ProjectID
	s += "\nkey path :" + c.KeyFilePath
	s += "\n--------------------------"
	return s
}

type SysConf struct {
	Log string `json:"log"`
	*SrvConf
	*TwitterConf
	*FileStoreConf
}

func (c *SysConf) String() any {
	var s = "\n=======================system config==========================="
	s += "\nlog level:" + c.Log
	s += "\n" + c.SrvConf.String()
	s += "\n" + c.TwitterConf.String()
	s += "\n" + c.FileStoreConf.String()
	s += "\n=============================================================="
	return s
}

func InitConf(c *SysConf) {
	_globalCfg = c
	util.SetLogLevel(c.Log)
	fmt.Println(c.String())
}
