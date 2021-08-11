/*
Mux是一个简单高效的路由分配器，支持标准库的net/http的接口。路由数据存储在前缀树preTree,由 https://github.com/obity/pretree 提供支持。

目前版本只支持变量路由和简单路由，例如/user/:id或者/user/login。
根据http请求方法GET、HEAD、POST、PUT、PATCH、DELETE、CONNECT、OPTIONS、TRACE分组单独处理，不同分组之间路由相互隔离。
路由变量的书写格式同时支持:id和{id},但实际存储的是:id。

Mux is a simple and efficient route distributor that supports the net/http interface of the standard library.Routing data is stored in the prefix tree preTree, supported by https://github.com/obity/pretree.

The current version only supports variable routes and simple routes such as /user/:id or /user/login.
According to the http request method GET, HEAD, POST, PUT, PATCH, DELETE, CONNECT, OPTIONS, TRACE grouping is handled separately, and the routes between different groups are isolated from each other.
Routing variable writing format supports both :id and {id},but the actual storage is :id.
*/
package mux
