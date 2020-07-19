package controllers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/inact25/E-WarungApi/masters/apis/models"
	"github.com/inact25/E-WarungApi/masters/apis/usecases"
	"github.com/inact25/E-WarungApi/utils"
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

	service, err := s.ServicesUseCase.GetAllServices()
	if err != nil {
		w.Write([]byte("Data Not Found"))
	}
	var resp = Res{Msg: "getAllServices", Data: service}
	byteOfService, err := json.Marshal(resp)
	if err != nil {
		w.Write([]byte("Something when Wrong"))
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(byteOfService)
}

func (s ServicesHandler) GetAllServicesByStatus(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	status := vars["status"]
	service, err := s.ServicesUseCase.GetAllServicesByStatus(status)
	if err != nil {
		w.Write([]byte("Data Not Found"))
	}
	var resp = Res{Msg: "getAllServices", Data: service}
	byteOfService, err := json.Marshal(resp)
	if err != nil {
		w.Write([]byte("Something when Wrong"))
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(byteOfService)
}

func (s ServicesHandler) AddNewServices(w http.ResponseWriter, r *http.Request) {
	dt := time.Now()
	day := dt.Format("2006.01.02 15:04:05")
	var resp = Res{}
	service := &models.ServicesModels{}
	w.Header().Set("Content-Type", "application/json")
	service.PriceDate = day
	err := json.NewDecoder(r.Body).Decode(&service)
	if err != nil {
		resp = Res{"Decode Failed", nil}
	}

	_, err = s.ServicesUseCase.AddNewServices(service)
	if err != nil {
		resp = Res{err.Error(), nil}
	} else {
		resp = Res{"Service Successfully Added", service}
	}
	byte, _ := json.Marshal(resp)
	w.Write(byte)
}

func (s ServicesHandler) UpdateServices(w http.ResponseWriter, r *http.Request) {
	services := &models.ServicesModels{}

	getJsonDataCheck := json.NewDecoder(r.Body).Decode(&services)
	utils.ErrorCheck(getJsonDataCheck, "Fatal")
	_, err := s.ServicesUseCase.UpdateServices(services)
	utils.ErrorCheck(err, "Fatal")
	w.Write([]byte("Service Updated"))

}

func (s ServicesHandler) UpdateServicesPrice(w http.ResponseWriter, r *http.Request) {
	dt := time.Now()
	day := dt.Format("2006.01.02 15:04:05")
	services := &models.ServicesModels{}
	getJsonDataCheck := json.NewDecoder(r.Body).Decode(&services)
	utils.ErrorCheck(getJsonDataCheck, "Fatal")
	_, err := s.ServicesUseCase.UpdateServicesPrice(day, services)
	utils.ErrorCheck(err, "Fatal")
	w.Write([]byte("Service Updated"))

}

func (s ServicesHandler) DeleteServices(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	_, err := s.ServicesUseCase.DeleteServices(vars["services_id"])
	if err != nil {
		w.Write([]byte("Data Not Found"))
	}
	w.Write([]byte("Data has been Deleted"))
}
