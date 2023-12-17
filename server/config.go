package server

import (
	"html/template"
	"net/http"
	"os"
	"path/filepath"
	"strings"
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
		"/register": "html/create_wallet.html",
	}

	htmlTemplateManager *template.Template
)

type LogicAction func(ts *TwitterSrv, w http.ResponseWriter, r *http.Request)

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

type FileStoreConf struct {
	ProjectID   string `json:"project_id"`
	KeyFilePath string `json:"key_file_path"`
}

type Conf struct {
	Log string `json:"log"`
	*SrvConf
	*TwitterConf
	*FileStoreConf
}

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
