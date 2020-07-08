package usecases

import "github.com/inact25/E-WarungApi/masters/apis/models"

type TransactionUseCases interface {
	GetAllTransactions() ([]*models.TransactionModels, error)
	AddNewTransactions(day string, transactions *models.TransactionModels) (string, error)
	UpdateTransactions(transactions *models.TransactionModels) (string, error)
}
