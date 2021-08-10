package main

import (
	"net/http"

	"github.com/obity/mux"
)

func main() {

	r := mux.NewMux()

	r.GET("/test", TestHandler)
	http.ListenAndServe("127.0.0.1:8000", r)
}

func TestHandler(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("hello, test"))
}
