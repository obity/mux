/*
   Copyright (c) 2021 ffactory.org
   Mux is licensed under Mulan PSL v2.
   You can use this software according to the terms and conditions of the Mulan PSL v2.
   You may obtain a copy of Mulan PSL v2 at:
               http://license.coscl.org.cn/MulanPSL2
   THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND, EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT, MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
   See the Mulan PSL v2 for more details.
*/

package mux

import (
	"crypto/sha1"
	"net/http"

	"github.com/obity/pretree"
)

// 路由分配器存储结构
//
// Routing distributor storage structure
type Mux struct {
	RouteGroup map[string]*route
}

// 实现 http.ServerHTTP 接口函数
//
// Implement the http.ServerHTTP interface function
func (m *Mux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path

	ok, tree := pretree.Query(r.Method, path)
	if !ok {
		http.NotFoundHandler().ServeHTTP(w, r)
		return
	}
	rule := tree.Rule()
	key := shortPath(rule)
	route := m.RouteGroup[key]
	route.handler.ServeHTTP(w, r)
}

// 新建路由分配器工厂函数
//
// New route allocator factory function
func NewMux() *Mux {
	return &Mux{RouteGroup: make(map[string]*route)}
}

// 启动http服务器,请在添加完路由最后在启动
//
// Start the http server, please start it at the end after adding the route.
func (m *Mux) Start(addr string) {
	http.ListenAndServe(addr, m)
}

// 启动带tls证书的https服务器,请在添加完路由后在启动
//
// Start the https server with tls certificate, please start it after adding the route
func (m *Mux) StartTLS(addr, certFile, keyFile string) {
	http.ListenAndServeTLS(addr, certFile, keyFile, m)
}

// 新建GET请求路由
//
// Create a new GET request route
func (m *Mux) GET(path string, f func(http.ResponseWriter, *http.Request)) {
	m.AppendRoute(http.MethodGet, path, f)
}

// 新建HEAD请求路由
//
// Create a new HEAD request route
func (m *Mux) HEAD(path string, f func(http.ResponseWriter, *http.Request)) {
	m.AppendRoute(http.MethodHead, path, f)
}

// 新建POST请求路由
//
// Create a new POST request route
func (m *Mux) POST(path string, f func(http.ResponseWriter, *http.Request)) {
	m.AppendRoute(http.MethodPost, path, f)
}

// 新建PUT请求路由
//
// Create a new PUT request route
func (m *Mux) PUT(path string, f func(http.ResponseWriter, *http.Request)) {
	m.AppendRoute(http.MethodPut, path, f)
}

// 新建PATCH请求路由
//
// Create a new PATCH request route
func (m *Mux) PATCH(path string, f func(http.ResponseWriter, *http.Request)) {
	m.AppendRoute(http.MethodPatch, path, f)
}

// 新建DELETE请求路由
//
// Create a new DELETE request route
func (m *Mux) DELETE(path string, f func(http.ResponseWriter, *http.Request)) {
	m.AppendRoute(http.MethodDelete, path, f)
}

// 新建CONNECT请求路由
//
// Create a new CONNECT request route
func (m *Mux) CONNECT(path string, f func(http.ResponseWriter, *http.Request)) {
	m.AppendRoute(http.MethodConnect, path, f)
}

// 新建OPTIONS请求路由
//
// Create a new OPTIONS request route
func (m *Mux) OPTIONS(path string, f func(http.ResponseWriter, *http.Request)) {
	m.AppendRoute(http.MethodOptions, path, f)
}

// 新建TRACE请求路由
//
// Create a new TRACE request route
func (m *Mux) TRACE(path string, f func(http.ResponseWriter, *http.Request)) {
	m.AppendRoute(http.MethodTrace, path, f)
}

// 通用新建路由函数
//
// Generic new route function
func (m *Mux) AppendRoute(method, path string, f func(http.ResponseWriter, *http.Request)) {
	route := newRoute()
	route.handler = http.HandlerFunc(f)
	pretree.Store(method, path)
	key := shortPath(path)
	m.RouteGroup[key] = route
}

// 缩短路径为sha1值
func shortPath(path string) string {
	h := sha1.New()
	b := h.Sum([]byte(path))
	return (string(b))
}
