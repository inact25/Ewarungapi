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
	r.HandleFunc("/data/", UsersHandler.OAuth).Methods(http.MethodPost)

	r.HandleFunc("/users/{id}", UsersHandler.GetSelfUser).Methods(http.MethodPost)
	r.HandleFunc("/users/{id}/", UsersHandler.GetSelfUser).Methods(http.MethodPost)

	r.HandleFunc("/users", UsersHandler.GetAllUser).Methods(http.MethodGet)
	r.HandleFunc("/users/", UsersHandler.GetAllUser).Methods(http.MethodGet)
}

func (u UsersHandler) OAuth(w http.ResponseWriter, r *http.Request) {
	users := &models.SelfUserModels{}
	var resp = Res{}
	w.Header().Set("Content-Type", "application/json")
	err := json.NewDecoder(r.Body).Decode(&users)
	log.Print("c : ", users)
	if err != nil {
		resp = Res{"Decode Failed", nil}
	}
	authData, err := u.UserUsecases.OAuth(users)
	if err != nil {
		resp = Res{err.Error(), nil}
	} else {
		resp = Res{"Login Successfully", authData}
	}
	bytes, _ := json.Marshal(resp)
	w.Write(bytes)
}

func (u UsersHandler) GetSelfUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	users := &models.SelfUserModels{}
	users.UserID = vars["id"]
	log.Print("c :", users)
	data, err := u.UserUsecases.GetSelfUser(users)
	if err != nil {
		w.Write([]byte("Data Not Found"))
	}
	var resp = Res{Msg: "getSelfUser", Data: data}
	byteOfData, err := json.Marshal(resp)
	if err != nil {
		w.Write([]byte("Something when Wrong"))
	}
	log.Print("c :", data)
	w.Header().Set("Content-Type", "application/json")
	w.Write(byteOfData)
}

func (u UsersHandler) GetAllUser(w http.ResponseWriter, r *http.Request) {
	data, err := u.UserUsecases.GetAllUsers()
	if err != nil {
		w.Write([]byte("Data Not Found"))
	}
	var resp = Res{Msg: "getAllUsers", Data: data}
	byteOfData, err := json.Marshal(resp)
	if err != nil {
		w.Write([]byte("Something when Wrong"))
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(byteOfData)
}
