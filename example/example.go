package main

import (
	"net/http"

	"github.com/obity/mux"
)

func main() {

	m := mux.NewMux()

	m.GET("/test/{id}", TestHandler)

	http.ListenAndServe("127.0.0.1:8001", m)
}

func TestHandler(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("matched:" + r.RequestURI))
	return
}
