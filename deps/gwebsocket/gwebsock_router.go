package gwebsocket

import (
	"main/deps/glog"
	"net/http"
)

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
	if server.IsConnLimitReached() {
		return
	}
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}
	defer conn.Close()
	// 连接成功
	client := &GWClient{Conn: conn}
	server.AddClient(client)
	hr.Handle(w, r, conn)
	glog.Info("websocket 连接关闭")
	server.RemoveClient(client)
}
