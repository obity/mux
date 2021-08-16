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

	"github.com/obity/mux/storage"
	"github.com/obity/pretree"
)

// 路由分配器存储结构
//
// Routing distributor storage structure
type Mux struct {
	basePath      string
	RouteGroup    map[string]*route
	StorageEngine storage.Engine
}

// 实现 http.ServerHTTP 接口函数
//
// Implement the http.ServerHTTP interface function
func (m *Mux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	ok, rule, vars := m.StorageEngine.Query(r.Method, path)
	if !ok {
		http.NotFoundHandler().ServeHTTP(w, r)
		return
	}
	r = SetVars(r, vars)
	key := shortPath(rule)
	route := m.RouteGroup[key]
	route.handler.ServeHTTP(w, r)
}

// 新建路由分配器工厂函数
//
// New route allocator factory function
func NewMux() *Mux {
	m := &Mux{RouteGroup: make(map[string]*route)}
	m.Default()
	return m
}

// 启用默认存储引擎
//
// Enable the default storage algorithm engine
func (m *Mux) Default() {
	m.SetEngine(pretree.NewPreTree())
}

// 修改存储算法引擎
//
// change storage algorithm engine
func (m *Mux) SetEngine(e storage.Engine) {
	m.StorageEngine = e
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

// 设置基础路径，例如API版本"/v1","/v2"
//
// Set the basic path, such as API version "/v1" and "/v2"
func (m *Mux) SetBasePath(basePath string) *Mux {
	m.basePath = basePath
	return m
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
	if len(m.basePath) > 0 {
		path = m.basePath + path
	}
	route := newRoute()
	route.handler = http.HandlerFunc(f)
	m.StorageEngine.Store(method, path)
	key := shortPath(path)
	m.RouteGroup[key] = route
}

// 缩短路径为sha1值
func shortPath(path string) string {
	h := sha1.New()
	b := h.Sum([]byte(path))
	return (string(b))
}
