package mux_test

import (
	"net/http"

	"github.com/obity/mux"
)

func Example_AppendRoute() {
	m := mux.NewMux()
	m.GET("/user/:id", UserHandler)
}

func Example_NewMux() {
	m := mux.NewMux()
	m.GET("/pet/findByStatus", Findbystatus)
	m.GET("/test/{id}", TestHandler)
	m.POST("/user/createWithList", Createwithlist)
	m.DELETE("/user/:username", Userinfo)
	m.Start(":8001")
}

func TestHandler(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("TestHandler matched:" + r.RequestURI))
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
