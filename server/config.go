package server

import (
	"html/template"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

var simpleRouterMap = map[string]string{
	"/":         "html/index.html",
	"/index":    "html/index.html",
	"/register": "html/create_wallet.html",
}

type LogicAction func(ts *TwitterSrv, w http.ResponseWriter, r *http.Request)

var logicRouter = map[string]LogicAction{
	"/signInByTwitter": signInByTwitter,
	"/tw_callback":     twitterSignCallBack,
}

type SrvConf struct {
	DebugMode   bool   `json:"debug_mode"`
	UseHttps    bool   `json:"use_https"`
	SSLCertFile string `json:"ssl_cert_file"`
	SSLKeyFile  string `json:"ssl_key_file"`
}
type TwitterConf struct {
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
}

type Conf struct {
	*SrvConf
	*TwitterConf
}

var templates = parseTemplates("assets/html")

func parseTemplates(path string) *template.Template {
	fs, err := os.ReadDir(path)
	if err != nil {
		panic(err)
	}

	var files []string
	for _, f := range fs {
		if strings.HasSuffix(f.Name(), ".html") {
			files = append(files, filepath.Join(path, f.Name()))
		}
	}

	return template.Must(template.ParseFiles(files...))
}
