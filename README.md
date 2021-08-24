# Mux
Mux is a simple and efficient route distributor that supports the net/http interface of the standard library.
Routing data is stored in the prefix tree preTree, supported by https://github.com/obity/pretree.

The current version only supports variable routes and simple routes such as /user/:id or /user/login.
According to the http request method GET, HEAD, POST, PUT, PATCH, DELETE, CONNECT, OPTIONS, TRACE grouping is handled separately, and the routes between different groups are isolated from each other.
Routing variable writing format supports both :id and {id},but the actual storage is :id.

# Doc

See this document at [GoDoc](https://pkg.go.dev/github.com/obity/mux)

# Install
    
    go get -u github.com/obity/mux@latest

# ToDo

- [x] 需要获取路由中的{var}变量，handler处理需要这些参数值传输到后端
- [x] 抽象算法存储引擎,支持自定义算法，只要实现Engine接口就可以自己启用 
- [x] 增加设置路由基础路径功能，统一设置API的版本如"/v1","/2"默认不启用

# Example
```
package main

import (
	"net/http"

	"github.com/obity/mux"
)

func main() {
	m := mux.NewMux()
	m.GET("/pet/findByStatus", Findbystatus)
	m.GET("/pet/{id}", PetHandler)
	m.POST("/user/createWithList", Createwithlist)
	m.DELETE("/user/:username", Userinfo)
	m.Start(":8001")
}

func PetHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	ID := vars["id"]
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("ID:" + ID))
	return
}

func Findbystatus(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("FindByStatus matched:" + r.RequestURI))
	return
}

func Createwithlist(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("CreateWithList matched:" + r.RequestURI))
	return
}

func Userinfo(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("UserInfo matched:" + r.RequestURI))
	return
}

func UserHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("UserInfo matched:" + r.RequestURI))
	return
}


```
