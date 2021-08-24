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
