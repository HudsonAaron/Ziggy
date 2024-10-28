package ghttp

import (
	"net/http"
)

// http路由结构体
type HRouter struct {
	Path   string                                   // 路由路径
	Handle func(http.ResponseWriter, *http.Request) // 路由处理函数
}

// 获取http路由
func DefaultHRouter() []HRouter {
	return []HRouter{}
}

// 运行多监听http路由
func RunMuxRouter(hrouter []HRouter) http.Handler {
	// 创建多监听路由
	mux := http.NewServeMux()
	for _, hr := range hrouter {
		mux.HandleFunc(hr.Path, hr.Handle)
	}
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		mux.ServeHTTP(w, r)
	})
}
