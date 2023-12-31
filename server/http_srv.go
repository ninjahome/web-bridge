package server

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/ninjahome/web-bridge/util"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"runtime/debug"
	"strings"
	"time"
)

const (
	staticFileDir = "assets"
)

type MainService struct {
	router *mux.Router
}

func httpRecover(url string) {
	if r := recover(); r != nil {
		stack := debug.Stack()
		util.LogInst().Error().Str("method", url).
			Str("panic", fmt.Sprintf("%v", r)).
			Str("stack", string(stack)).
			Msg("api service failed")
	}
}

func NewMainService() *MainService {
	r := mux.NewRouter()
	bh := &MainService{
		router: r,
	}

	r.PathPrefix("/" + staticFileDir + "/").HandlerFunc(bh.assetsRouter)
	r.PathPrefix("/" + staticFileDir + "/").Handler(http.StripPrefix("/"+staticFileDir+"/", http.FileServer(http.Dir(staticFileDir))))

	for route, fileName := range cfgHtmlFileRouter {
		r.HandleFunc(route, func(w http.ResponseWriter, r *http.Request) {
			bh.assetsStaticFile(w, r, fileName)
		})
	}
	r.HandleFunc("/user_profile/{web3-id}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		web3ID := vars["web3-id"]
		userProfile(w, r, web3ID)
	})

	for route, twService := range cfgActionRouter {
		var url, action = route, twService
		r.HandleFunc(route, func(writer http.ResponseWriter, request *http.Request) {
			defer httpRecover(url)
			util.LogInst().Debug().Str("url", url).Send()

			if request.ContentLength > util.MaxReqContentLen {
				err := fmt.Errorf("content length too large (%d>%d)", request.ContentLength, util.MaxReqContentLen)
				http.Error(writer, err.Error(), http.StatusRequestEntityTooLarge)
				return
			}
			if !action.NeedToken {
				action.Action(writer, request, nil)
				return
			}

			var token = validateUsrRights(request)
			if token == nil {
				http.Redirect(writer, request, "/signIn", http.StatusFound)
				return
			}
			action.Action(writer, request, token)
		}).Methods("GET", "POST")
	}

	return bh
}

func (bh *MainService) Start() {
	cfg := _globalCfg
	if cfg.UseHttps {
		if cfg.SSLCertFile == "" || cfg.SSLKeyFile == "" {
			panic("HTTPS needs ssl key and cert files")
		}
		fmt.Print("HTTPS Mode")
		panic(http.ListenAndServeTLS(":443", cfg.SSLCertFile, cfg.SSLKeyFile, bh.router))
	} else {
		fmt.Print("Simple HTTP")
		panic(http.ListenAndServe(":"+cfg.HttpPort, bh.router))
	}
}

func (bh *MainService) assetsRouter(writer http.ResponseWriter, request *http.Request) {
	realUrlPath := request.URL.Path
	if strings.HasSuffix(realUrlPath, ".map") {
		realUrlPath = strings.TrimSuffix(realUrlPath, ".map")
	}
	realUrlPath = strings.TrimPrefix(realUrlPath, "/"+staticFileDir+"/")
	bh.assetsStaticFile(writer, request, realUrlPath)
}

func (bh *MainService) assetsStaticFile(writer http.ResponseWriter, request *http.Request, fileName string) {
	fPath := filepath.Join(staticFileDir, fileName)

	fileInfo, err := os.Stat(fPath)
	if err != nil {
		if os.IsNotExist(err) {
			http.Error(writer, "File not found", http.StatusNotFound)
		} else {
			http.Error(writer, "Internal Server Error", http.StatusInternalServerError)
		}
		log.Print(err)
		return
	}

	modTime := fileInfo.ModTime()
	if _globalCfg.RefreshContent {
		writer.Header().Set("Cache-Control", "max-age=0, no-store, must-revalidate")
	} else {
		writer.Header().Set("Cache-Control", "public, max-age=1036000") // 缓存10天
	}

	writer.Header().Set("Last-Modified", modTime.UTC().Format(http.TimeFormat))

	if t, err := time.Parse(http.TimeFormat, request.Header.Get("If-Modified-Since")); err == nil && modTime.Before(t.Add(1*time.Second)) {
		writer.WriteHeader(http.StatusNotModified)
		return
	}

	http.ServeFile(writer, request, fPath)
}
