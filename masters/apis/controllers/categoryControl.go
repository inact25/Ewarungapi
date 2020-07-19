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

	category, err := s.CategoriesUseCase.GetAllCategories()
	if err != nil {
		w.Write([]byte("Data Not Found"))
	}
	var resp = Res{Msg: "getAllCategory", Data: category}
	byteOfCategory, err := json.Marshal(resp)
	if err != nil {
		w.Write([]byte("Something when Wrong"))
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(byteOfCategory)
}

func (s CategoriesHandler) GetAllCategoriesByStatus(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	status := vars["status"]
	category, err := s.CategoriesUseCase.GetAllCategoriesByStatus(status)
	if err != nil {
		w.Write([]byte("Data Not Found"))
	}
	var resp = Res{Msg: "getAllCategory", Data: category}
	byteOfCategory, err := json.Marshal(resp)
	if err != nil {
		w.Write([]byte("Something when Wrong"))
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(byteOfCategory)
}

func (s CategoriesHandler) AddNewCategories(w http.ResponseWriter, r *http.Request) {
	dt := time.Now()
	day := dt.Format("2006.01.02 15:04:05")
	var resp = Res{}
	category := &models.CategoriesModels{}
	w.Header().Set("Content-Type", "application/json")
	category.PriceDate = day
	err := json.NewDecoder(r.Body).Decode(&category)
	if err != nil {
		resp = Res{"Decode Failed", nil}
	}
	_, err = s.CategoriesUseCase.AddNewCategories(category)
	if err != nil {
		resp = Res{err.Error(), nil}
	} else {
		resp = Res{"Category Successfully Added", category}
	}
	byte, _ := json.Marshal(resp)
	w.Write(byte)
}

func (s CategoriesHandler) UpdateCategories(w http.ResponseWriter, r *http.Request) {
	categories := &models.CategoriesModels{}

	getJsonDataCheck := json.NewDecoder(r.Body).Decode(&categories)
	utils.ErrorCheck(getJsonDataCheck, "Fatal")
	_, err := s.CategoriesUseCase.UpdateCategories(categories)
	utils.ErrorCheck(err, "Fatal")
	w.Write([]byte("Category Updated"))

}

func (s CategoriesHandler) UpdateCategoriesPrice(w http.ResponseWriter, r *http.Request) {
	dt := time.Now()
	day := dt.Format("2006.01.02 15:04:05")
	categories := &models.CategoriesModels{}
	getJsonDataCheck := json.NewDecoder(r.Body).Decode(&categories)
	utils.ErrorCheck(getJsonDataCheck, "Fatal")
	_, err := s.CategoriesUseCase.UpdateCategoriesPrice(day, categories)
	utils.ErrorCheck(err, "Fatal")
	w.Write([]byte("Category Updated"))

}

func (s CategoriesHandler) DeleteCategories(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	_, err := s.CategoriesUseCase.DeleteCategories(vars["categories_id"])
	if err != nil {
		w.Write([]byte("Data Not Found"))
	}
	w.Write([]byte("Data has been Deleted"))
}
