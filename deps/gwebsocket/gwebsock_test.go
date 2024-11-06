package gwebsocket

import (
	"main/deps/glog"
	"net/http"

	"github.com/gorilla/websocket"
)

func _TestGWS() {
	_ = DoStart("127.0.0.1:8080/", _GetWSRouter())
}

func _GetWSRouter() []WSRouter {
	return []WSRouter{
		{"/", _HandleConn},
	}
}

// 处理WebSocket连接请求
func _HandleConn(w http.ResponseWriter, r *http.Request, conn *websocket.Conn) {
	// 设置返回状态码
	for {
		msgType, bytes, err := conn.ReadMessage()
		if err != nil {
			glog.Error("ReadMessage error:%v", err)
			break
		}
		switch msgType {
		case websocket.TextMessage:
			glog.Info("msgType:%d, string(bytes):%s", msgType, string(bytes))
			gwsm := GWSMsg{Field: "ok", Value: true}
			conn.WriteJSON(gwsm)
		case websocket.BinaryMessage:
			glog.Info("msgType:%d, string(bytes):%s", msgType, string(bytes))
			gwsm := GWSMsg{Field: "ok", Value: true}
			conn.WriteJSON(gwsm)
		case websocket.CloseMessage:
			glog.Info("msgType:%d, string(bytes):%s", msgType, string(bytes))
			gwsm := GWSMsg{Field: "ok", Value: true}
			conn.WriteJSON(gwsm)
		case websocket.PingMessage:
			glog.Info("msgType:%d, string(bytes):%s", msgType, string(bytes))
			gwsm := GWSMsg{Field: "ok", Value: true}
			conn.WriteJSON(gwsm)
		case websocket.PongMessage:
			glog.Info("msgType:%d, string(bytes):%s", msgType, string(bytes))
			gwsm := GWSMsg{Field: "ok", Value: true}
			conn.WriteJSON(gwsm)
		default:
			glog.Info("msgType:%d, string(bytes):%s", msgType, string(bytes))
			gwsm := GWSMsg{Field: "ok", Value: true}
			conn.WriteJSON(gwsm)
		}
	}
}
