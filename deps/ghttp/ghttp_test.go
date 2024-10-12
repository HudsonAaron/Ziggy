package ghttp

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func TestGhttp() {
	_ = Start("http://127.0.0.1:8080/", nil)
}

func GetHRouter() []HRouter {
	return []HRouter{
		{"/", OptionsHandler},
		{"/status/", StatusHandler},
		{"/info/", InfoHandler},
	}
}

// http options请求
func OptionsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type,Content-Length,Accept-Encoding,X-CSRF-Token,Authorization")
}

// http status请求
func StatusHandler(w http.ResponseWriter, r *http.Request) {
	// 设置返回状态码
	fmt.Println("真正的中间业务逻辑 StatusHandler")
	fmt.Fprintf(w, "OK")
}

// http info请求
func InfoHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("真正的中间业务逻辑 InfoHandler")
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	log.Println("接收到的数据：", string(body))
	fmt.Fprintf(w, "OK")
}

// http get请求
func GetHandler() {
	// 设置超时5秒
	client := &http.Client{Timeout: 5 * 1000 * 1000 * 1000}
	client.Get("http://localhost:8080/status")
}
