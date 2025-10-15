package ghttp

import (
	"fmt"
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
func Start(httpConf map[string]any, hr []HRouter) error {
	addr, err := GetHttpConf(httpConf)
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

// 获取配置
func GetHttpConf(httpConf map[string]any) (string, error) {
	var domain = ""
	if httpConf["domain"] != "" { // domain字段存在
		if _domain, ok := httpConf["domain"].(string); ok {
			domain = _domain
			return domain, nil
		} else {
			return "", fmt.Errorf("http domain type error")
		}
	} else { // domain字段不存在，只存在ip和port字段
		var ip string = ""
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
				port = int(_port)
			} else {
				return "", fmt.Errorf("http port type error")
			}
		}
		domain = fmt.Sprintf("%s:%d", ip, port)
		return domain, nil
	}
}
