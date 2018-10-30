package controller

import (
	"encoding/json"
	"estor/core"
	"github.com/julienschmidt/httprouter"
	"github.com/ssrs100/logs"
	"io/ioutil"
	"net/http"
)


var (
	log = logs.GetLogger()
)

func GetUsers(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	users := core.GetUsers()
	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func GetUser(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	w.WriteHeader(http.StatusOK)
}

func CreateUser(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Error("Receive body failed: %v", err.Error())
		DefaultHandler.ServeHTTP(w, req, err, http.StatusBadRequest)
		return
	}
	defer req.Body.Close()

	var userReq = &core.User{}
	err = json.Unmarshal(body, userReq)
	if err != nil {
		log.Error("Invalid body. err:%s", err.Error())
		DefaultHandler.ServeHTTP(w, req, err, http.StatusBadRequest)
		return
	}
	err = core.AddUser(userReq)
	if err != nil {
		log.Error("Create user fail. err:%s", err.Error())
		DefaultHandler.ServeHTTP(w, req, err, http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func DeleteUser(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	userName := ps.ByName("username")
	if err := core.DelUser(userName); err != nil {
		log.Error("Delete user fail. err:%s", err.Error())
		DefaultHandler.ServeHTTP(w, req, err, http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
}
