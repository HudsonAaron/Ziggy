package gwebsocket

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

// WebSocket路由结构体
type WSRouter struct {
	Path   string                                                    // 路由路径
	Handle func(http.ResponseWriter, *http.Request, *websocket.Conn) // 路由处理函数
}

// 获取http路由
func DefaultWSRouter() []WSRouter {
	return []WSRouter{}
}

// 运行多监听http路由
func RunMuxRouter(wsrouter []WSRouter) http.Handler {
	// 创建多监听路由
	mux := http.NewServeMux()
	for _, hr := range wsrouter {
		mux.HandleFunc(hr.Path, hr.ActualHandle)
	}
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		mux.ServeHTTP(w, r)
	})
}

// 实际处理函数
func (hr *WSRouter) ActualHandle(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}
	defer conn.Close()
	// 连接成功
	client := &GWClient{conn: conn}
	server.clients[client] = true
	hr.Handle(w, r, conn)
	fmt.Println("websocket 连接关闭")
	delete(server.clients, client)
}
