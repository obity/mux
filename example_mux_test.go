package mux_test

import (
	"net/http"

	"github.com/obity/mux"
)

func Example() {
	m := mux.NewMux()
	m.GET("/pet/findByStatus", FindbystatusHandler)
	m.GET("/pet/{id}", PetHandler)
	m.POST("/user/createWithList", CreatewithlistHandler)
	m.DELETE("/user/:username", UserinfoHandler)
	m.Start(":8001")
}
func ExampleNewMux() {
	m := mux.NewMux()
	m.GET("/pet/findByStatus", FindbystatusHandler)
	m.GET("/pet/{id}", PetHandler)
	m.POST("/user/createWithList", CreatewithlistHandler)
	m.DELETE("/user/:username", UserinfoHandler)
	m.Start(":8001")
}

func ExampleMux_GET_1() {
	m := mux.NewMux()
	m.GET("/user/{id}", UserinfoHandler)
}

func ExampleMux_GET_2() {
	m := mux.NewMux()
	m.GET("/user/:id", UserinfoHandler)
}

func ExampleMux_POST_1() {
	m := mux.NewMux()
	m.POST("/account", UserinfoHandler)
	m.POST("/account/:id/address", UserAddressHandler)
	m.POST("/sendemail", SendEmailandler)
}

func ExampleMux_POST_2() {
	m := mux.NewMux()
	m.POST("/account", UserinfoHandler)
	m.POST("/account/{id}/address", UserAddressHandler)
	m.POST("/sendemail", SendEmailandler)
}

func PetHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("TestHandler matched:" + r.RequestURI))
	return
}

func FindbystatusHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("FindByStatus matched:" + r.RequestURI))
	return
}

func CreatewithlistHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("CreateWithList matched:" + r.RequestURI))
	return
}

func UserinfoHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("UserInfo matched:" + r.RequestURI))
	return
}

func UserHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("UserInfo matched:" + r.RequestURI))
	return
}

func UserAddressHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("UserInfo matched:" + r.RequestURI))
	return
}

func SendEmailandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("UserInfo matched:" + r.RequestURI))
	return
}
