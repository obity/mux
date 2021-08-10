package mux

import (
	"crypto/sha1"
	"net/http"

	"github.com/obity/pretree"
)

type Mux struct {
	RouteGroup map[string]*Route
}

func (m *Mux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path

	ok, tree := pretree.Query(r.Method, path)
	if !ok {
		http.NotFoundHandler().ServeHTTP(w, r)
	}
	rule := tree.Rule()
	key := ShortPath(rule)
	route := m.RouteGroup[key]
	route.handler.ServeHTTP(w, r)
}

func NewMux() *Mux {
	return &Mux{RouteGroup: make(map[string]*Route)}
}

func (m *Mux) GET(path string, f func(http.ResponseWriter, *http.Request)) {
	m.AddRoute(http.MethodGet, path, f)
}

func (m *Mux) HEAD(path string, f func(http.ResponseWriter, *http.Request)) {
	m.AddRoute(http.MethodHead, path, f)
}

func (m *Mux) POST(path string, f func(http.ResponseWriter, *http.Request)) {
	m.AddRoute(http.MethodPost, path, f)
}

func (m *Mux) PUT(path string, f func(http.ResponseWriter, *http.Request)) {
	m.AddRoute(http.MethodPut, path, f)
}

func (m *Mux) PATCH(path string, f func(http.ResponseWriter, *http.Request)) {
	m.AddRoute(http.MethodPatch, path, f)
}

func (m *Mux) DELETE(path string, f func(http.ResponseWriter, *http.Request)) {
	m.AddRoute(http.MethodDelete, path, f)
}

func (m *Mux) CONNECT(path string, f func(http.ResponseWriter, *http.Request)) {
	m.AddRoute(http.MethodConnect, path, f)
}

func (m *Mux) OPTIONS(path string, f func(http.ResponseWriter, *http.Request)) {
	m.AddRoute(http.MethodOptions, path, f)
}

func (m *Mux) TRACE(path string, f func(http.ResponseWriter, *http.Request)) {
	m.AddRoute(http.MethodTrace, path, f)
}

func (m *Mux) AddRoute(method, path string, f func(http.ResponseWriter, *http.Request)) {
	route := NewRoute()
	route.handler = http.HandlerFunc(f)
	pretree.Store(method, path)
	key := ShortPath(path)
	m.RouteGroup[key] = route
}

func ShortPath(path string) string {
	h := sha1.New()
	b := h.Sum([]byte(path))
	return (string(b))

}
