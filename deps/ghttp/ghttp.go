package ghttp

import (
	"main/deps/glog"
	"net/http"
)

// http服务结构体
type HServer struct {
	addr   string
	handle http.Handler
}

var (
	Version = "1.0.0" // 版本号
)

// 创建http监听和接收
func Start(addr string, hr []HRouter) error {
	var hrouter = []HRouter{}
	if hr != nil {
		hrouter = hr
	} else {
		hrouter = DefaultHRouter()
	}
	handle := RunMuxRouter(hrouter)
	hs := &HServer{
		addr:   addr,
		handle: handle,
	}
	go hs.StartServe()
	glog.Info("Http Server start up")
	return nil
}

// 运行http服务
func (hs *HServer) StartServe() {
	// 创建http监听
	err := http.ListenAndServe(hs.addr, hs.handle)
	if err != nil {
		glog.Error("http server start error:%s", err.Error())
	}
}
