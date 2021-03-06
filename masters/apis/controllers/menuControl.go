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

type MenusHandler struct {
	MenuUseCase usecases.MenuUseCases
}

func MenuControll(r *mux.Router, service usecases.MenuUseCases) {
	MenuHandler := MenusHandler{service}
	r.HandleFunc("/menus", MenuHandler.GetAllMenu).Methods(http.MethodGet)
	r.HandleFunc("/menus/", MenuHandler.GetAllMenu).Methods(http.MethodGet)

	r.HandleFunc("/menus/prices", MenuHandler.GetAllMenuPrices).Methods(http.MethodGet)
	r.HandleFunc("/menus/prices/", MenuHandler.GetAllMenuPrices).Methods(http.MethodGet)

	r.HandleFunc("/menus/{status}", MenuHandler.GetAllMenuByStatus).Methods(http.MethodGet)
	r.HandleFunc("/menus/{status}/", MenuHandler.GetAllMenuByStatus).Methods(http.MethodGet)

	r.HandleFunc("/menus", MenuHandler.AddNewMenu).Methods(http.MethodPost)
	r.HandleFunc("/menus/", MenuHandler.AddNewMenu).Methods(http.MethodPost)

	r.HandleFunc("/menus", MenuHandler.UpdateMenu).Methods(http.MethodPut)
	r.HandleFunc("/menus/", MenuHandler.UpdateMenu).Methods(http.MethodPut)

	r.HandleFunc("/menus/prices", MenuHandler.UpdateMenuPrice).Methods(http.MethodPut)
	r.HandleFunc("/menus/prices/", MenuHandler.UpdateMenuPrice).Methods(http.MethodPut)

	r.HandleFunc("/menus/{menu_id}", MenuHandler.DeleteMenu).Methods(http.MethodDelete)
	r.HandleFunc("/menus/{menu_id}/", MenuHandler.DeleteMenu).Methods(http.MethodDelete)

}

func (s MenusHandler) GetAllMenu(w http.ResponseWriter, r *http.Request) {

	menu, err := s.MenuUseCase.GetAllMenu()
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

func (s MenusHandler) GetAllMenuByStatus(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	status := vars["status"]
	log.Println("C : ", status)
	menu, err := s.MenuUseCase.GetAllMenuByStatus(status)
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

func (s MenusHandler) GetAllMenuPrices(w http.ResponseWriter, r *http.Request) {

	menu, err := s.MenuUseCase.GetAllMenuPrices()
	log.Println("c : ", menu)
	if err != nil {
		w.Write([]byte("Data Not Found"))
	}
	var resp = Res{Msg: "getAllMenuPrices", Data: menu}
	byteOfMenu, err := json.Marshal(resp)
	if err != nil {
		w.Write([]byte("Something when Wrong"))
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(byteOfMenu)
}

func (s MenusHandler) AddNewMenu(w http.ResponseWriter, r *http.Request) {
	dt := time.Now()
	day := dt.Format("2006.01.02 15:04:05")
	menus := &models.MenuModels{}
	getJsonDataCheck := json.NewDecoder(r.Body).Decode(&menus)
	log.Println(getJsonDataCheck)
	log.Println(menus)
	utils.ErrorCheck(getJsonDataCheck, "Fatal")
	_, err := s.MenuUseCase.AddNewMenu(day, menus)
	utils.ErrorCheck(err, "Fatal")
	w.Write([]byte("Category Succesfully Added"))

}

func (s MenusHandler) UpdateMenu(w http.ResponseWriter, r *http.Request) {
	menus := &models.MenuModels{}

	getJsonDataCheck := json.NewDecoder(r.Body).Decode(&menus)
	log.Println("C :", menus)
	utils.ErrorCheck(getJsonDataCheck, "Fatal")
	_, err := s.MenuUseCase.UpdateMenu(menus)
	utils.ErrorCheck(err, "Fatal")
	w.Write([]byte("Menu Updated"))

}

func (s MenusHandler) UpdateMenuPrice(w http.ResponseWriter, r *http.Request) {
	dt := time.Now()
	day := dt.Format("2006.01.02 15:04:05")
	menus := &models.MenuPriceModels{}
	getJsonDataCheck := json.NewDecoder(r.Body).Decode(&menus)
	log.Println("C :", menus)
	utils.ErrorCheck(getJsonDataCheck, "Fatal")
	_, err := s.MenuUseCase.UpdateMenuPrice(day, menus)
	utils.ErrorCheck(err, "Fatal")
	w.Write([]byte("Menu Updated"))

}

func (s MenusHandler) DeleteMenu(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	_, err := s.MenuUseCase.DeleteMenu(vars["menu_id"])
	if err != nil {
		w.Write([]byte("Data Not Found"))
	}
	w.Write([]byte("Data has been Deleted"))
}
