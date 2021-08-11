package mux

import "net/http"

type Route struct {
	handler http.Handler
}

func NewRoute() *Route {
	return &Route{}
}
