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
	"context"
	"net/http"
)

// 路由对应处理函数
type route struct {
	handler http.Handler
}

// 新建路由工厂方法
//
// New routing factory method
func newRoute() *route {
	return &route{}
}

type ContextType int

// 路由变量存储在上下文中的key
//
// The key of the routing variable stored in the context
const VarsKey ContextType = iota

// 把路由变量放到http请求中传递
//
// Route variables are passed in HTTP requests
func SetVars(r *http.Request, vars map[string]string) *http.Request {
	return r.WithContext(context.WithValue(r.Context(), VarsKey, vars))
}

// 获取路由变量
//
// Get routing variables
func Vars(r *http.Request) map[string]string {
	v := r.Context().Value(VarsKey)
	if nil == v {
		return nil
	}
	return v.(map[string]string)
}
