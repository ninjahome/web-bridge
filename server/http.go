package server

import (
	"fmt"
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
	bh := &MainService{}

	http.Handle("/"+staticFileDir+"/", http.StripPrefix("/"+staticFileDir+"/", http.HandlerFunc(bh.assetsRouter)))

	for route, fileName := range cfgHtmlFileRouter {
		http.HandleFunc(route, bh.simpleRouter(fileName))
	}

	for route, twService := range cfgActionRouter {
		var url, action = route, twService
		http.HandleFunc(url, func(writer http.ResponseWriter, request *http.Request) {
			defer httpRecover(url)
			util.LogInst().Debug().Str("url", url).Send()
			if request.Method != http.MethodPost && request.Method != http.MethodGet {
				http.Error(writer, "", http.StatusMethodNotAllowed)
				return
			}

			if request.ContentLength > util.MaxReqContentLen {
				err := fmt.Errorf("content length too large (%d>%d)", request.ContentLength, util.MaxReqContentLen)
				http.Error(writer, err.Error(), http.StatusRequestEntityTooLarge)
				return
			}

			action(writer, request)
		})
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
		panic(http.ListenAndServeTLS(":443", cfg.SSLCertFile, cfg.SSLKeyFile, nil))
	} else {
		fmt.Print("Simple HTTP")
		panic(http.ListenAndServe(":80", nil))
	}
}

func (bh *MainService) simpleRouter(fileName string) func(http.ResponseWriter, *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		bh.assetsStaticFile(writer, request, fileName)
	}
}

func (bh *MainService) assetsRouter(writer http.ResponseWriter, request *http.Request) {
	realUrlPath := request.URL.Path
	if strings.HasSuffix(realUrlPath, ".map") {
		realUrlPath = strings.TrimSuffix(realUrlPath, ".map")
	}
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
