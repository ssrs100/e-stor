package controller

import "net/http"

func GetUsers(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func GetUser(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func CreateUser(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func DeleteUser(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusOK)
}
