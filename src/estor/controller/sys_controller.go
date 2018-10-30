package controller

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func HealthCheck(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("I am OK."))
}
