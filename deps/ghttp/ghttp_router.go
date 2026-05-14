package ghttp

import (
	"net/http"
)

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
