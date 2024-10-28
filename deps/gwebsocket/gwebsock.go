package gwebsocket

import (
	"main/deps/glog"
	"net/http"

	"github.com/gorilla/websocket"
)

var (
	// 创建WebSocket Upgrader对象，用于升级HTTP连接为WebSocket连接
	upgrader = websocket.Upgrader{
		// 允许所有CORS请求
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	// 创建WebSocket服务端对象
	server = NewServer()
)

// WebSocket客户端结构体
type GWClient struct {
	conn *websocket.Conn
}

// WebSocket服务端结构体
type GWServer struct {
	// 客户端连接
	clients map[*GWClient]bool
	// 消息广播通道
	broadcast chan []byte
}

// WebSocket服务结构体
type WSServer struct {
	addr   string
	handle http.Handler
}

// 创建WebSocket对象
func NewServer() *GWServer {
	return &GWServer{
		clients:   make(map[*GWClient]bool),
		broadcast: make(chan []byte),
	}
}

// 获取WebSocket服务端对象
func GetServer() *GWServer {
	return server
}

// 创建http监听和接收
func Start(addr string, wsr []WSRouter) error {
	var wsrouter = DefaultWSRouter()
	if wsr != nil {
		wsrouter = wsr
	}
	handle := RunMuxRouter(wsrouter)
	wss := &WSServer{
		addr:   addr,
		handle: handle,
	}
	go wss.StartServe()      // 启动websocket服务
	go server.BroadcastMsg() // 启动广播协程
	glog.Info("WebSocket Server start up")
	return nil
}

// 运行http服务
func (wss *WSServer) StartServe() {
	// 创建http监听
	err := http.ListenAndServe(wss.addr, wss.handle)
	if err != nil {
		glog.Error("http server start error:%s", err.Error())
	}
}

// 获取当前连接数
func GetConnCount() int {
	return len(server.clients)
}
