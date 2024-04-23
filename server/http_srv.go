package server

import (
	"fmt"
	"github.com/gorilla/csrf"
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
		var url, file = route, fileName
		r.HandleFunc(url, func(w http.ResponseWriter, r *http.Request) {
			bh.assetsStaticFile(w, r, file)
		})
	}
	CSRF := csrf.Protect([]byte(_globalCfg.SessionKey), csrf.FieldName("X-CSRF-Token"), csrf.Secure(_globalCfg.UseHttps)) // 在生产中启用 Secure
	securedRouter := r.PathPrefix("/").Subrouter()
	securedRouter.Use(CSRF)

	// Unsecured routes (No CSRF)
	unsecuredRouter := r.PathPrefix("/").Subrouter()
	for route, twService := range cfgActionRouter {
		var url, action = route, twService
		var targetRouter = securedRouter // Default to secured router

		// Determine if the route needs CSRF protection
		if !action.NeedToken {
			targetRouter = unsecuredRouter // No CSRF for callback APIs
		}

		var funcs = func(writer http.ResponseWriter, request *http.Request) {
			defer httpRecover(url)
			util.LogInst().Debug().Str("url", url).Send()

			if request.ContentLength > util.MaxReqContentLen {
				err := fmt.Errorf("content length too large (%d>%d)", request.ContentLength, util.MaxReqContentLen)
				util.LogInst().Err(err).Msg("request invalid")
				http.Error(writer, err.Error(), http.StatusRequestEntityTooLarge)
				return
			}

			if !action.NeedToken {
				action.Action(writer, request, nil)
				return
			}

			if url == "/buyRights" || url == "/buyFromShare" {
				var param = OuterLinkParam{
					TweetID:  request.URL.Query().Get(NjTweetID),
					ShareID:  request.URL.Query().Get(SharedID),
					ShareUsr: request.URL.Query().Get(SharedUsr),
				}
				_ = SMInst().Set(request, writer, BuyRightsUrlKey, param.Data())
			}

			var njUserData = validateUsrRights(request)
			if njUserData == nil {
				http.Redirect(writer, request, "/signIn", http.StatusFound)
				return
			}
			action.Action(writer, request, njUserData)
		}

		targetRouter.HandleFunc(route, funcs)
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
