package controllers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/inact25/E-WarungApi/masters/apis/models"
	"github.com/inact25/E-WarungApi/masters/apis/usecases"
	"log"
	"net/http"
)

type UsersHandler struct {
	UserUsecases usecases.UsersUseCases
}

func UsersControll(r *mux.Router, service usecases.UsersUseCases) {
	UsersHandler := UsersHandler{service}
	//r.Use(middlewares.ActivityLogMiddleware)
	r.HandleFunc("/oauth", UsersHandler.OAuth).Methods(http.MethodPost)
	r.HandleFunc("/oauth/", UsersHandler.OAuth).Methods(http.MethodPost)
}

func (s UsersHandler) OAuth(w http.ResponseWriter, r *http.Request) {
	users := &models.SelfUserModels{}
	var resp = Res{}
	w.Header().Set("Content-Type", "application/json")
	err := json.NewDecoder(r.Body).Decode(&users)
	log.Print("c : ", users)
	if err != nil {
		resp = Res{"Decode Failed", nil}
	}
	authData, err := s.UserUsecases.OAuth(users)
	if err != nil {
		resp = Res{err.Error(), nil}
	} else {
		resp = Res{"Login Successfully", authData}
	}
	bytes, _ := json.Marshal(resp)
	w.Write(bytes)
}
