package ghttp

import (
	"context"
	"fmt"
	"main/deps/glog"
	"main/deps/gsafego"
	"math"
	"net"
	"net/http"
	"time"
)

// 创建http监听和接收
func Start(httpConf map[string]any, hr []HRouter) error {
	serverMu.Lock()
	defer serverMu.Unlock()
	addr, err := GetHttpConf(httpConf)
	// glog.Info("Http Service addr:%s", addr)
	if err != nil {
		return err
	}
	var hrouter = []HRouter{}
	if hr != nil {
		hrouter = hr
	} else {
		hrouter = DefaultHRouter()
	}
	handle := RunMuxRouter(hrouter)
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}
	// 创建HTTP服务器
	hs := &HServer{
		addr:     addr,
		handle:   handle,
		listener: listener,
		server: &http.Server{
			Addr:    addr,
			Handler: handle,
		},
	}
	serverInstance = hs
	gsafego.SafeGoRebootless(hs.startServe)
	// go hs.startServe()
	glog.Info("Http Service started successfully!")
	return nil
}

// 运行http服务
func (hs *HServer) startServe() {
	// 创建http监听
	err := hs.server.Serve(hs.listener)
	if err != nil && err != http.ErrServerClosed {
		glog.Error("http server start error:%s", err.Error())
	}
}

// 关闭http服务
func Stop() error {
	serverMu.Lock()
	instance := serverInstance
	serverMu.Unlock()
	if instance != nil {
		if err := instance.Close(); err != nil {
			return err
		}
		serverMu.Lock()
		serverInstance = nil
		serverMu.Unlock()
		return nil
	}
	glog.Info("Http Service already stopped")
	return nil
}

func (hs *HServer) Close() error {
	if hs.server != nil {
		// 创建超时上下文
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		// 优雅关闭HTTP服务器
		err := hs.server.Shutdown(ctx)
		if err != nil {
			glog.Error("http server shutdown error:%s", err.Error())
			return err
		}
		glog.Info("Http Service stopped")
	}
	return nil
}

// 获取配置
func GetHttpConf(httpConf map[string]any) (string, error) {
	var domain = ""
	var ip string = "127.0.0.1"
	if httpConf["ip"] != nil {
		if _ip, ok := httpConf["ip"].(string); ok {
			ip = _ip
		} else {
			return "", fmt.Errorf("http ip type error")
		}
	}
	var port int = 0
	if httpConf["port"] != nil {
		if _port, ok := httpConf["port"].(int); ok {
			port = _port
		} else if _port, ok := httpConf["port"].(float64); ok {
			if math.Trunc(_port) != _port {
				return "", fmt.Errorf("http port must be an integer")
			}
			port = int(_port)
		} else {
			return "", fmt.Errorf("http port type error")
		}
	}
	if port <= 0 || port > 65535 {
		return "", fmt.Errorf("http port value out of range")
	}
	domain = fmt.Sprintf("%s:%d", ip, port)
	return domain, nil
}
