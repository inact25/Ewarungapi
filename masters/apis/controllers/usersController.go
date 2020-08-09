package controllers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/inact25/E-WarungApi/masters/apis/models"
	"github.com/inact25/E-WarungApi/masters/apis/usecases"
	"net/http"
)

type UsersHandler struct {
	UserUsecases usecases.UsersUseCases
}

func UsersControll(r *mux.Router, service usecases.UsersUseCases) {
	UsersHandler := UsersHandler{service}
	r.HandleFunc("/oauth", UsersHandler.OAuth).Methods(http.MethodPost)
	r.HandleFunc("/oauth/", UsersHandler.OAuth).Methods(http.MethodPost)

	r.HandleFunc("/users/{id}", UsersHandler.GetUser).Methods(http.MethodPost)
	r.HandleFunc("/users/{id}/", UsersHandler.GetUser).Methods(http.MethodPost)

	r.HandleFunc("/users", UsersHandler.GetAllUser).Methods(http.MethodGet)
	r.HandleFunc("/users/", UsersHandler.GetAllUser).Methods(http.MethodGet)
}

func (u UsersHandler) OAuth(w http.ResponseWriter, r *http.Request) {
	users := &models.UserModels{}
	var resp = Res{}
	w.Header().Set("Content-Type", "application/json")
	err := json.NewDecoder(r.Body).Decode(&users)
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

func (u UsersHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	users := &models.UserModels{}
	users.UserID = vars["id"]
	var resp = Res{}
	data, err := u.UserUsecases.GetUser(users)
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		resp = Res{Msg: "Data not found", Data: nil}
	}
	resp = Res{Msg: "getUser", Data: data}
	byteOfData, err := json.Marshal(resp)
	if err != nil {
		w.Write([]byte("Something when Wrong"))
	}
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
