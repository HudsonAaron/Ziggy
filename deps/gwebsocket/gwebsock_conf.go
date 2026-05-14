package gwebsocket

import "fmt"

// 获取配置
func GetWebSockConf(webSockConf map[string]any) (string, error) {
	var domain = ""
	var ip string = ""
	if webSockConf["ip"] != nil {
		if _ip, ok := webSockConf["ip"].(string); ok {
			ip = _ip
		} else {
			return "", fmt.Errorf("websocket ip type error")
		}
	}
	var port int = 0
	if webSockConf["port"] != nil {
		if _port, ok := webSockConf["port"].(int); ok {
			port = _port
		} else if _port, ok := webSockConf["port"].(float64); ok {
			port = int(_port)
		} else {
			return "", fmt.Errorf("websocket port type error")
		}
	}
	domain = fmt.Sprintf("%s:%d", ip, port)
	return domain, nil
}
