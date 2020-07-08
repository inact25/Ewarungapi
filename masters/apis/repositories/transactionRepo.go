package repositories

import "github.com/inact25/E-WarungApi/masters/apis/models"

type TransactionRepositories interface {
	GetAllTransactions() ([]*models.TransactionModels, error)
	AddNewTransactions(day string, transactionModels *models.TransactionModels) (string, error)
	UpdateTransactions(transactions *models.TransactionModels) (string, error)
}
