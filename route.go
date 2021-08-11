package mux

import "net/http"

// 路由对应处理函数
type route struct {
	handler http.Handler
}

// 新建空路由工厂方法
func newRoute() *route {
	return &route{}
}
