# Mux
Mux is a simple and efficient route distributor that supports the net/http interface of the standard library.
Routing data is stored in the prefix tree preTree, supported by https://github.com/obity/pretree.

The current version only supports variable routes and simple routes such as /user/:id or /user/login.
According to the http request method GET, HEAD, POST, PUT, PATCH, DELETE, CONNECT, OPTIONS, TRACE grouping is handled separately, and the routes between different groups are isolated from each other.
Routing variable writing format supports both :id and {id},but the actual storage is :id.

# Doc

See this document at [GoDoc](https://pkg.go.dev/github.com/obity/mux)

# ToDo

- [ ] 需要获取路由中的{var}变量，handler处理需要这些参数值传输到后端
