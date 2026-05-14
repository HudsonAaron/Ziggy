package gwebsocket

import (
	"errors"
	"fmt"
	"main/deps/glog"
	"net"
	"net/http"
	"time"
)

// 创建WebSocket对象
func NewServer() *GWServer {
	return &GWServer{
		clients:   make(map[*GWClient]bool),
		broadcast: make(chan int),
	}
}

// 获取WebSocket服务端对象
func GetServer() *GWServer {
	return server
}

// 创建http监听和接收
func Start(wsConf map[string]any, wsr []WSRouter) error {
	addr, err := GetWebSockConf(wsConf)
	// glog.Info("WebSocket Service addr:%s", addr)
	if err != nil {
		return err
	}
	return doStart(addr, wsr)
}

// 创建http监听和接收
func doStart(addr string, wsr []WSRouter) error {
	stateLock.Lock()
	defer stateLock.Unlock()
	if started {
		return fmt.Errorf("websocket service already started")
	}

	var wsrouter = DefaultWSRouter()
	if wsr != nil {
		wsrouter = wsr
	}
	handle := RunMuxRouter(wsrouter)
	ln, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}

	server = NewServer()
	wss = &WSServer{
		addr:   addr,
		handle: handle,
		server: &http.Server{
			Addr:    addr,
			Handler: handle,
		},
	}
	go wss.StartServe(ln)    // 启动websocket服务
	go server.broadcastMsg() // 启动广播协程
	started = true
	// go ShowConnCount()       // 启动显示连接数协程
	glog.Info("WebSocket Service started successfully!")
	return nil
}

// 运行http服务
func (wss *WSServer) StartServe(ln net.Listener) {
	err := wss.server.Serve(ln)
	if err != nil && !errors.Is(err, http.ErrServerClosed) {
		glog.Error("http server start error:%s", err.Error())
	}
}

// 关闭WebSocket服务
func Stop() error {
	stateLock.Lock()
	if !started || wss == nil || wss.server == nil {
		stateLock.Unlock()
		return fmt.Errorf("websocket service not started")
	}
	_ws := wss
	_gws := server
	started = false
	stateLock.Unlock()

	// 关闭http服务
	err := _ws.server.Close()
	if err != nil {
		glog.Error("http server close error:%s", err.Error())
	}
	_gws.broadcast <- ActionClose
	glog.Info("WebSocket Service stopped successfully!")
	return nil
}

// 获取当前连接数
func GetConnCount() int {
	return server.ConnCount()
}

func (gws *GWServer) ConnCount() int {
	gws.lock.RLock()
	defer gws.lock.RUnlock()
	return len(gws.clients)
}

func (gws *GWServer) IsConnLimitReached() bool {
	return gws.ConnCount() >= maxConnCount
}

func (gws *GWServer) AddClient(client *GWClient) {
	gws.lock.Lock()
	defer gws.lock.Unlock()
	gws.clients[client] = true
}

func (gws *GWServer) RemoveClient(client *GWClient) {
	gws.lock.Lock()
	defer gws.lock.Unlock()
	delete(gws.clients, client)
}

// TODO: 瞬间连接、瞬间断开 无法承受

// 定时显示连接数
func ShowConnCount() {
	for {
		glog.Info("WebSocket Server current conn count:%d", GetConnCount())
		<-time.After(10 * time.Second)
	}
}
