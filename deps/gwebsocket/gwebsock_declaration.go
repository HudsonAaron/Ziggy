package gwebsocket

import (
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

var (
	Version = "1.0.1"
	// 创建WebSocket Upgrader对象，用于升级HTTP连接为WebSocket连接
	upgrader = websocket.Upgrader{
		// 允许所有CORS请求
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
	// 创建WebSocket服务端对象
	server = NewServer()
	// 最大客户端连接数
	maxConnCount = 2000
	wss          *WSServer
	stateLock    sync.Mutex
	started      bool
)

const (
	ActionClose     = 0
	ActionBroadcast = 1
)

// WebSocket客户端结构体
type GWClient struct {
	Conn *websocket.Conn
}

// WebSocket服务端结构体
type GWServer struct {
	// 客户端连接
	clients map[*GWClient]bool
	// 消息广播通道
	broadcast chan int
	// 读写锁
	lock sync.RWMutex
}

// WebSocket服务结构体
type WSServer struct {
	addr   string
	handle http.Handler
	server *http.Server
}

// WebSocket路由结构体
type WSRouter struct {
	Path   string                                                    // 路由路径
	Handle func(http.ResponseWriter, *http.Request, *websocket.Conn) // 路由处理函数
}
