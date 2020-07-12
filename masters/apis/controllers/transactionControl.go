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

type TransactionHandler struct {
	TransactionUseCases usecases.TransactionUseCases
}

func TransactionControll(r *mux.Router, service usecases.TransactionUseCases) {
	TransactionHandler := TransactionHandler{service}
	//r.Use(middlewares.ActivityLogMiddleware)
	r.HandleFunc("/transactions", TransactionHandler.GetAllTransaction).Methods(http.MethodGet)
	r.HandleFunc("/transactions/", TransactionHandler.GetAllTransaction).Methods(http.MethodGet)

	r.HandleFunc("/transactions", TransactionHandler.AddNewTransaction).Methods(http.MethodPost)
	r.HandleFunc("/transactions/", TransactionHandler.AddNewTransaction).Methods(http.MethodPost)

	r.HandleFunc("/transactions", TransactionHandler.UpdateTransaction).Methods(http.MethodPut)
	r.HandleFunc("/transactions/", TransactionHandler.UpdateTransaction).Methods(http.MethodPut)
}

func (s TransactionHandler) GetAllTransaction(w http.ResponseWriter, r *http.Request) {

	transaction, err := s.TransactionUseCases.GetAllTransactions()
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

func (s TransactionHandler) AddNewTransaction(w http.ResponseWriter, r *http.Request) {
	dt := time.Now()
	day := dt.Format("2006.01.02 15:04:05")
	var resp = Res{}
	transactions := &models.TransactionModels{}
	w.Header().Set("Content-Type", "application/json")
	transactions.TransactionDate = day
	err := json.NewDecoder(r.Body).Decode(&transactions)
	if err != nil {
		resp = Res{"Decode Failed", nil}
	}
	_, err = s.TransactionUseCases.AddNewTransactions(transactions)
	if err != nil {
		resp = Res{err.Error(), nil}
	} else {
		resp = Res{"Category Successfully Added", nil}
	}
	byte, _ := json.Marshal(resp)
	w.Write(byte)
}

func (s TransactionHandler) UpdateTransaction(w http.ResponseWriter, r *http.Request) {
	transactions := &models.TransactionModels{}
	getJsonDataCheck := json.NewDecoder(r.Body).Decode(&transactions)
	log.Println("C :", transactions)
	utils.ErrorCheck(getJsonDataCheck, "Fatal")
	_, err := s.TransactionUseCases.UpdateTransactions(transactions)
	if err != nil {
		log.Println("err : ", err)
		w.Write([]byte("Data Not Found"))
	}
}
