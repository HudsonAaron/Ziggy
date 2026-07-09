package ghttp

import (
	"net"
	"net/http"
	"sync"
)

var (
	Version        = "1.0.2" // 版本号
	serverInstance *HServer  // http服务实例
	serverMu       sync.Mutex
)

// http服务结构体
type HServer struct {
	addr     string
	handle   http.Handler
	server   *http.Server
	listener net.Listener
}

// http路由结构体
type HRouter struct {
	Path   string                                   // 路由路径
	Handle func(http.ResponseWriter, *http.Request) // 路由处理函数
}
