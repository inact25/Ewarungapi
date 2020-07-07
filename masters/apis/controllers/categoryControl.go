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

type CategoryHandler struct {
	CategoryUseCases usecases.CategoryUseCases
}

func CategoriesControll(r *mux.Router, service usecases.CategoryUseCases) {
	CategoryHandler := CategoryHandler{service}
	//r.Use(middlewares.ActivityLogMiddleware)
	r.HandleFunc("/categories", CategoryHandler.GetAllCategories).Methods(http.MethodGet)
	r.HandleFunc("/categories/", CategoryHandler.GetAllCategories).Methods(http.MethodGet)

	r.HandleFunc("/categories/prices", CategoryHandler.GetAllCategoriesPrice).Methods(http.MethodGet)
	r.HandleFunc("/categories/prices/", CategoryHandler.GetAllCategoriesPrice).Methods(http.MethodGet)

	r.HandleFunc("/categories", CategoryHandler.AddNewCategory).Methods(http.MethodPost)
	r.HandleFunc("/categories/", CategoryHandler.AddNewCategory).Methods(http.MethodPost)

	r.HandleFunc("/categories/price", CategoryHandler.AddNewCategoryPrice).Methods(http.MethodPost)
	r.HandleFunc("/categories/price/", CategoryHandler.AddNewCategoryPrice).Methods(http.MethodPost)
}

func (s CategoryHandler) GetAllCategories(w http.ResponseWriter, r *http.Request) {

	categories, err := s.CategoryUseCases.GetAllCategories()
	if err != nil {
		w.Write([]byte("Data Not Found"))
	}
	var resp = Res{Msg: "getAllCategories", Data: categories}
	byteOfCategories, err := json.Marshal(resp)
	if err != nil {
		w.Write([]byte("Something when Wrong"))
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(byteOfCategories)
}

func (s CategoryHandler) GetAllCategoriesPrice(w http.ResponseWriter, r *http.Request) {

	categories, err := s.CategoryUseCases.GetAllCategoriesPrice()
	if err != nil {
		w.Write([]byte("Data Not Found"))
	}
	var resp = Res{Msg: "getAllCategoriesPRice", Data: categories}
	byteOfCategories, err := json.Marshal(resp)
	if err != nil {
		w.Write([]byte("Something when Wrong"))
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(byteOfCategories)
}

func (s CategoryHandler) AddNewCategory(w http.ResponseWriter, r *http.Request) {
	categories := &models.CategoryModels{}
	getJsonDataCheck := json.NewDecoder(r.Body).Decode(&categories)
	utils.ErrorCheck(getJsonDataCheck, "Print")
	_, err := s.CategoryUseCases.AddNewCategory(categories)
	utils.ErrorCheck(err, "Print")
	w.Write([]byte("Category Succesfully Added"))

}

func (s CategoryHandler) AddNewCategoryPrice(w http.ResponseWriter, r *http.Request) {
	dt := time.Now()
	day := dt.Format("2006.01.02 15:04:05")
	categories := &models.CategoryPriceModels{}
	var getJsonDataCheck = json.NewDecoder(r.Body).Decode(&categories)
	utils.ErrorCheck(getJsonDataCheck, "Print")
	_, err := s.CategoryUseCases.AddNewCategoryPrice(day, categories)
	utils.ErrorCheck(err, "Print")
	w.Write([]byte("Category Succesfully Added"))

}
