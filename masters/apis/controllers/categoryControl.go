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

type CategoriesHandler struct {
	CategoriesUseCase usecases.CategoryUseCases
}

func CategoriesControll(r *mux.Router, service usecases.CategoryUseCases) {
	CategoriesHandler := CategoriesHandler{service}

	r.HandleFunc("/categories", CategoriesHandler.GetAllCategories).Methods(http.MethodGet)
	r.HandleFunc("/categories/", CategoriesHandler.GetAllCategories).Methods(http.MethodGet)

	r.HandleFunc("/categories/{status}", CategoriesHandler.GetAllCategoriesByStatus).Methods(http.MethodGet)
	r.HandleFunc("/categories/{status}/", CategoriesHandler.GetAllCategoriesByStatus).Methods(http.MethodGet)

	r.HandleFunc("/categories", CategoriesHandler.AddNewCategories).Methods(http.MethodPost)
	r.HandleFunc("/categories/", CategoriesHandler.AddNewCategories).Methods(http.MethodPost)

	r.HandleFunc("/categories", CategoriesHandler.UpdateCategories).Methods(http.MethodPut)
	r.HandleFunc("/categories/", CategoriesHandler.UpdateCategories).Methods(http.MethodPut)

	r.HandleFunc("/categories/prices", CategoriesHandler.UpdateCategoriesPrice).Methods(http.MethodPut)
	r.HandleFunc("/categories/prices/", CategoriesHandler.UpdateCategoriesPrice).Methods(http.MethodPut)

	r.HandleFunc("/categories/{categories_id}", CategoriesHandler.DeleteCategories).Methods(http.MethodDelete)
	r.HandleFunc("/categories/{categories_id}/", CategoriesHandler.DeleteCategories).Methods(http.MethodDelete)

}

func (s CategoriesHandler) GetAllCategories(w http.ResponseWriter, r *http.Request) {

	menu, err := s.CategoriesUseCase.GetAllCategories()
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

func (s CategoriesHandler) GetAllCategoriesByStatus(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	status := vars["status"]
	log.Println("C : ", status)
	menu, err := s.CategoriesUseCase.GetAllCategoriesByStatus(status)
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

func (s CategoriesHandler) AddNewCategories(w http.ResponseWriter, r *http.Request) {
	dt := time.Now()
	day := dt.Format("2006.01.02 15:04:05")
	categories := &models.CategoriesModels{}
	getJsonDataCheck := json.NewDecoder(r.Body).Decode(&categories)
	log.Println(getJsonDataCheck)
	log.Println(categories)
	utils.ErrorCheck(getJsonDataCheck, "Fatal")
	_, err := s.CategoriesUseCase.AddNewCategories(day, categories)
	utils.ErrorCheck(err, "Fatal")
	w.Write([]byte("Category Succesfully Added"))
}

func (s CategoriesHandler) UpdateCategories(w http.ResponseWriter, r *http.Request) {
	categories := &models.CategoriesModels{}

	getJsonDataCheck := json.NewDecoder(r.Body).Decode(&categories)
	log.Println("C :", categories)
	utils.ErrorCheck(getJsonDataCheck, "Fatal")
	_, err := s.CategoriesUseCase.UpdateCategories(categories)
	utils.ErrorCheck(err, "Fatal")
	w.Write([]byte("Menu Updated"))

}

func (s CategoriesHandler) UpdateCategoriesPrice(w http.ResponseWriter, r *http.Request) {
	dt := time.Now()
	day := dt.Format("2006.01.02 15:04:05")
	categories := &models.CategoriesModels{}
	getJsonDataCheck := json.NewDecoder(r.Body).Decode(&categories)
	log.Println("C :", categories)
	utils.ErrorCheck(getJsonDataCheck, "Fatal")
	_, err := s.CategoriesUseCase.UpdateCategoriesPrice(day, categories)
	utils.ErrorCheck(err, "Fatal")
	w.Write([]byte("Menu Updated"))

}

func (s CategoriesHandler) DeleteCategories(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	_, err := s.CategoriesUseCase.DeleteCategories(vars["categories_id"])
	if err != nil {
		w.Write([]byte("Data Not Found"))
	}
	w.Write([]byte("Data has been Deleted"))
}
