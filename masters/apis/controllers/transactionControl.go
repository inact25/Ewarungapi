package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/inact25/E-WarungApi/masters/apis/models"
	"github.com/inact25/E-WarungApi/masters/apis/usecases"
	"github.com/inact25/E-WarungApi/utils"
	"net/http"
	"time"
)

type TransactionHandler struct {
	TransactionUseCases usecases.TransactionUseCases
}

func TransactionControll(r *mux.Router, service usecases.TransactionUseCases) {
	TransactionHandler := TransactionHandler{service}
	//r.Use(middlewares.ActivityLogMiddleware)
	r.HandleFunc("/transactions", TransactionHandler.GetAllTransaction).Methods(http.MethodGet)
	r.HandleFunc("/transactions/", TransactionHandler.GetAllTransaction).Methods(http.MethodGet)

	r.HandleFunc("/transactions/daily", TransactionHandler.GetDailyTransaction).Methods(http.MethodGet)
	r.HandleFunc("/transactions/daily/", TransactionHandler.GetDailyTransaction).Methods(http.MethodGet)

	r.HandleFunc("/transactions", TransactionHandler.AddNewTransaction).Methods(http.MethodPost)
	r.HandleFunc("/transactions/", TransactionHandler.AddNewTransaction).Methods(http.MethodPost)

	r.HandleFunc("/transactions", TransactionHandler.UpdateTransaction).Methods(http.MethodPut)
	r.HandleFunc("/transactions/", TransactionHandler.UpdateTransaction).Methods(http.MethodPut)

	r.HandleFunc("/transactions/{transaction_id}", TransactionHandler.DeleteTransaction).Methods(http.MethodDelete)
	r.HandleFunc("/transactions/{transaction_id}/", TransactionHandler.DeleteTransaction).Methods(http.MethodDelete)

}

func (s TransactionHandler) GetAllTransaction(w http.ResponseWriter, r *http.Request) {

	transaction, err := s.TransactionUseCases.GetAllTransaction()
	if err != nil {
		w.Write([]byte("Data Not Found"))
	}
	var resp = Res{Msg: "getAllTransaction", Data: transaction}
	byteOfTransaction, err := json.Marshal(resp)
	if err != nil {
		w.Write([]byte("Something when Wrong"))
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(byteOfTransaction)
}

func (s TransactionHandler) GetDailyTransaction(w http.ResponseWriter, r *http.Request) {
	dt := time.Now()
	day := dt.Format("2006-01-02")
	currentDay := fmt.Sprintf("%%" + day + "%%")
	transaction, err := s.TransactionUseCases.GetDailyTransaction(currentDay)
	if err != nil {
		w.Write([]byte("Data Not Found"))
	}
	var resp = Res{Msg: "getDailyTransaction", Data: transaction}
	byteOfTransaction, err := json.Marshal(resp)
	if err != nil {
		w.Write([]byte("Something when Wrong"))
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(byteOfTransaction)
}

func (s TransactionHandler) AddNewTransaction(w http.ResponseWriter, r *http.Request) {
	dt := time.Now()
	day := dt.Format("2006.01.02 15:04:05")
	transactions := &models.TransactionModels{}
	getJsonDataCheck := json.NewDecoder(r.Body).Decode(&transactions)
	utils.ErrorCheck(getJsonDataCheck, "Fatal")
	_, err := s.TransactionUseCases.AddTransaction(day, transactions)
	if err != nil {
		w.Write([]byte("Data Not Found"))
	}
	w.Write([]byte("Data Successfully Added"))
	w.Header().Set("Content-Type", "application/json")
}

func (s TransactionHandler) DeleteTransaction(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	transactionID := vars["transaction_id"]
	_, err := s.TransactionUseCases.DeleteTransaction(transactionID)
	if err != nil {
		w.Write([]byte("Data Not Found"))
	}

	w.Write([]byte("Data Successfully Deleted"))
	w.Header().Set("Content-Type", "application/json")
}

func (s TransactionHandler) UpdateTransaction(w http.ResponseWriter, r *http.Request) {
	transactions := &models.TransactionModels{}
	getJsonDataCheck := json.NewDecoder(r.Body).Decode(&transactions)
	utils.ErrorCheck(getJsonDataCheck, "Fatal")
	_, err := s.TransactionUseCases.UpdateTransaction(transactions)
	if err != nil {
		w.Write([]byte("Data Not Found"))
	}
	w.Write([]byte("Data Successfully Updated"))
	w.Header().Set("Content-Type", "application/json")

}
