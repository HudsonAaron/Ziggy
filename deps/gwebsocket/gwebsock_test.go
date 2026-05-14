package gwebsocket

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"main/deps/glog"
	"net/http"
	"net/url"
	"strconv"
	"testing"
	"time"

	"github.com/gorilla/websocket"
)

func _TestGWS() {
	_ = doStart("127.0.0.1:8080", _GetWSRouter())
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

// / 客户端连接
var addr = flag.String("addr", "127.0.0.1:9003", "http service address")

func WSClient(n int) {
	flag.Parse()
	log.SetFlags(0)

	u := url.URL{Scheme: "ws", Host: *addr, Path: "/ws"}

	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Println("dial:", err)
		return
	}
	defer c.Close()

	fmt.Printf("%d connected to %s\n", n, u.String())

	for {
		msg := map[string]any{
			"hello": "i'm " + strconv.Itoa(n),
			"ok":    true,
		}
		message, _ := json.Marshal(msg)
		err = c.WriteMessage(websocket.TextMessage, message)
		if err != nil {
			log.Fatal("write:", err)
		}
		_, _, err := c.ReadMessage()
		if err != nil {
			log.Fatal("read:", err)
		}
		// fmt.Printf("recv: %s", message)
		randNum := 10
		time.Sleep(time.Second * time.Duration(randNum))
	}
}

func TestStartStopLifecycle(t *testing.T) {
	err := doStart("127.0.0.1:0", _GetWSRouter())
	if err != nil {
		t.Fatalf("start failed: %v", err)
	}

	err = doStart("127.0.0.1:0", _GetWSRouter())
	if err == nil {
		t.Fatalf("expected repeated start to fail")
	}

	err = Stop()
	if err != nil {
		t.Fatalf("stop failed: %v", err)
	}

	err = Stop()
	if err == nil {
		t.Fatalf("expected repeated stop to fail")
	}
}

func TestBroadcastCloseClearsClients(t *testing.T) {
	gws := NewServer()
	gws.AddClient(&GWClient{})
	gws.AddClient(&GWClient{})
	if got := gws.ConnCount(); got != 2 {
		t.Fatalf("unexpected initial conn count: %d", got)
	}

	go gws.broadcastMsg()
	gws.broadcast <- ActionClose
	time.Sleep(50 * time.Millisecond)

	if got := gws.ConnCount(); got != 0 {
		t.Fatalf("expected clients to be cleared, got: %d", got)
	}
}
