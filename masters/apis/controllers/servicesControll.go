package controllers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/inact25/E-WarungApi/masters/apis/models"
	"github.com/inact25/E-WarungApi/masters/apis/usecases"
	"github.com/inact25/E-WarungApi/utils"
	"log"
	"net/http"
	"time"
)

type ServicesHandler struct {
	ServicesUseCase usecases.ServiceUseCases
}

func ServicesControll(r *mux.Router, service usecases.ServiceUseCases) {
	ServicesHandler := ServicesHandler{service}

	r.HandleFunc("/services", ServicesHandler.GetAllServices).Methods(http.MethodGet)
	r.HandleFunc("/services/", ServicesHandler.GetAllServices).Methods(http.MethodGet)

	r.HandleFunc("/services/{status}", ServicesHandler.GetAllServicesByStatus).Methods(http.MethodGet)
	r.HandleFunc("/services/{status}/", ServicesHandler.GetAllServicesByStatus).Methods(http.MethodGet)

	r.HandleFunc("/services", ServicesHandler.AddNewServices).Methods(http.MethodPost)
	r.HandleFunc("/services/", ServicesHandler.AddNewServices).Methods(http.MethodPost)

}

func (s ServicesHandler) GetAllServices(w http.ResponseWriter, r *http.Request) {

	menu, err := s.ServicesUseCase.GetAllServices()
	log.Println("c : ", menu)
	if err != nil {
		w.Write([]byte("Data Not Found"))
	}
	var resp = Res{Msg: "getAllMenu", Data: menu}
	byteOfMenu, err := json.Marshal(resp)
	if err != nil {
		w.Write([]byte("Something when Wrong"))
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(byteOfMenu)
}

func (s ServicesHandler) GetAllServicesByStatus(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	status := vars["status"]
	log.Println("C : ", status)
	menu, err := s.ServicesUseCase.GetAllServicesByStatus(status)
	log.Println("c : ", menu)
	if err != nil {
		w.Write([]byte("Data Not Found"))
	}
	var resp = Res{Msg: "getAllMenu", Data: menu}
	byteOfMenu, err := json.Marshal(resp)
	if err != nil {
		w.Write([]byte("Something when Wrong"))
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(byteOfMenu)
}

func (s ServicesHandler) AddNewServices(w http.ResponseWriter, r *http.Request) {
	dt := time.Now()
	day := dt.Format("2006.01.02 15:04:05")
	services := &models.ServicesModels{}
	getJsonDataCheck := json.NewDecoder(r.Body).Decode(&services)
	log.Println(getJsonDataCheck)
	log.Println(services)
	utils.ErrorCheck(getJsonDataCheck, "Fatal")
	_, err := s.ServicesUseCase.AddNewServices(day, services)
	utils.ErrorCheck(err, "Fatal")
	w.Write([]byte("Category Succesfully Added"))

}
