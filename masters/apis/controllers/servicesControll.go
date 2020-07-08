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

	r.HandleFunc("/services", ServicesHandler.UpdateServices).Methods(http.MethodPut)
	r.HandleFunc("/services/", ServicesHandler.UpdateServices).Methods(http.MethodPut)

	r.HandleFunc("/services/prices", ServicesHandler.UpdateServicesPrice).Methods(http.MethodPut)
	r.HandleFunc("/services/prices/", ServicesHandler.UpdateServicesPrice).Methods(http.MethodPut)

	r.HandleFunc("/services/{services_id}", ServicesHandler.DeleteServices).Methods(http.MethodDelete)
	r.HandleFunc("/services/{services_id}/", ServicesHandler.DeleteServices).Methods(http.MethodDelete)

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

func (s ServicesHandler) UpdateServices(w http.ResponseWriter, r *http.Request) {
	services := &models.ServicesModels{}

	getJsonDataCheck := json.NewDecoder(r.Body).Decode(&services)
	log.Println("C :", services)
	utils.ErrorCheck(getJsonDataCheck, "Fatal")
	_, err := s.ServicesUseCase.UpdateServices(services)
	utils.ErrorCheck(err, "Fatal")
	w.Write([]byte("Menu Updated"))

}

func (s ServicesHandler) UpdateServicesPrice(w http.ResponseWriter, r *http.Request) {
	dt := time.Now()
	day := dt.Format("2006.01.02 15:04:05")
	services := &models.ServicesModels{}
	getJsonDataCheck := json.NewDecoder(r.Body).Decode(&services)
	log.Println("C :", services)
	utils.ErrorCheck(getJsonDataCheck, "Fatal")
	_, err := s.ServicesUseCase.UpdateServicesPrice(day, services)
	utils.ErrorCheck(err, "Fatal")
	w.Write([]byte("Menu Updated"))

}

func (s ServicesHandler) DeleteServices(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	_, err := s.ServicesUseCase.DeleteServices(vars["services_id"])
	if err != nil {
		w.Write([]byte("Data Not Found"))
	}
	w.Write([]byte("Data has been Deleted"))
}
