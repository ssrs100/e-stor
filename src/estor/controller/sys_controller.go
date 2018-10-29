package controller

import "net/http"

func HealthCheck(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("I am OK."))
}
