package server

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
)

const staticFileDir = "assets" // 静态文件目录

type MainService struct {
	cfg   *Conf
	twSrv *TwitterSrv
}

func NewMainService(cfg *Conf) *MainService {
	bh := &MainService{cfg: cfg}
	http.Handle("/"+staticFileDir+"/", http.StripPrefix("/"+staticFileDir+"/", http.HandlerFunc(bh.assetsRouter)))
	for route, fileName := range simpleRouterMap {
		http.HandleFunc(route, bh.simpleRouter(fileName))
	}

	twSrv := NewTwitterSrv(cfg.TwitterConf)
	for route, twService := range logicRouter {
		log.Println(route, twService)
		var r, s = route, twService
		http.HandleFunc(r, func(writer http.ResponseWriter, request *http.Request) {
			log.Println(r, s, request.URL.Path)
			s(twSrv, writer, request)
		})
	}

	return bh
}

func (bh *MainService) Start() {
	if bh.cfg.UseHttps {
		if bh.cfg.SSLCertFile == "" || bh.cfg.SSLKeyFile == "" {
			panic("HTTPS 服务器需要指定证书文件和私钥文件")
		}
		fmt.Print("HTTPS模式")
		panic(http.ListenAndServeTLS(":443", bh.cfg.SSLCertFile, bh.cfg.SSLKeyFile, nil))
	} else {
		fmt.Print("简单模式")
		panic(http.ListenAndServe(":80", nil))
	}
}

func (bh *MainService) simpleRouter(fileName string) func(http.ResponseWriter, *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		bh.assetsStaticFile(writer, request, fileName)
	}
}

func (bh *MainService) assetsRouter(writer http.ResponseWriter, request *http.Request) {
	// 获取文件路径
	realUrlPath := request.URL.Path
	if strings.HasSuffix(realUrlPath, ".map") {
		realUrlPath = strings.TrimSuffix(realUrlPath, ".map")
	}
	bh.assetsStaticFile(writer, request, realUrlPath)
}

func (bh *MainService) assetsStaticFile(writer http.ResponseWriter, request *http.Request, fileName string) {
	fPath := filepath.Join(staticFileDir, fileName)

	// 获取文件状态
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
	if bh.cfg.DebugMode {
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
