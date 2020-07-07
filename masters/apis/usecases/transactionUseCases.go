package usecases

import "github.com/inact25/E-WarungApi/masters/apis/models"

type TransactionUseCases interface {
	GetAllTransaction() ([]*models.TransactionModels, error)
	GetDailyTransaction(date string) ([]*models.TransactionModels, error)
	AddTransaction(day string, transactions *models.TransactionModels) (string, error)
	UpdateTransaction(transactions *models.TransactionModels) (string, error)
	DeleteTransaction(transactionID string) (string, error)
}
