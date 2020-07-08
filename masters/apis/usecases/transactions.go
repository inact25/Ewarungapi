package usecases

import (
	"github.com/inact25/E-WarungApi/masters/apis/models"
	"github.com/inact25/E-WarungApi/masters/apis/repositories"
	"log"
)

type TransactionUseCaseImpl struct {
	transactionRepo repositories.TransactionRepositories
}

func (s TransactionUseCaseImpl) GetAllTransactions() ([]*models.TransactionModels, error) {
	transaction, err := s.transactionRepo.GetAllTransactions()
	if err != nil {
		return nil, err
	}
	return transaction, nil
}

func (s TransactionUseCaseImpl) AddNewTransactions(day string, transactions *models.TransactionModels) (string, error) {
	log.Println("U :", transactions)
	transaction, err := s.transactionRepo.AddNewTransactions(day, transactions)
	if err != nil {
		return "", err
	}
	return transaction, nil
}

func (s TransactionUseCaseImpl) UpdateTransactions(transactions *models.TransactionModels) (string, error) {
	log.Println("U :", transactions)
	transaction, err := s.transactionRepo.UpdateTransactions(transactions)
	if err != nil {
		return "", err
	}
	return transaction, nil
}

func InitTransactionUseCase(transactionRepo repositories.TransactionRepositories) TransactionUseCases {
	return &TransactionUseCaseImpl{transactionRepo}
}
