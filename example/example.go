package main

import (
	"net/http"

	"github.com/obity/mux"
)

func main() {

	m := mux.NewMux()

	m.GET("/pet/findByStatus", FindByStatus)
	m.GET("/test/{id}", TestHandler)
	m.POST("/user/createWithList", CreateWithList)
	m.DELETE("/user/:username", UserInfo)
	m.Start(":8001")

}

func TestHandler(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("TestHandler matched:" + r.RequestURI))
	return
}

func FindByStatus(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("FindByStatus matched:" + r.RequestURI))
	return
}

func CreateWithList(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("CreateWithList matched:" + r.RequestURI))
	return

}

func UserInfo(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("UserInfo matched:" + r.RequestURI))
	return
}
