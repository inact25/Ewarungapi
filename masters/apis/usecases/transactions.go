package usecases

import (
	"github.com/inact25/E-WarungApi/masters/apis/models"
	"github.com/inact25/E-WarungApi/masters/apis/repositories"
	"github.com/inact25/E-WarungApi/utils/validation"
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

func (s TransactionUseCaseImpl) AddNewTransactions(transactions *models.TransactionModels) (string, error) {
	log.Println("u:", transactions)
	err := validation.CheckEmpty(transactions.ServicesDesc, transactions.MenuDesc, transactions.CategoryDesc, transactions.Qty)
	if err != nil {
		return "", err
	}
	err = validation.CheckInt(transactions.Qty)
	if err != nil {
		return "", err
	}
	transaction, err := s.transactionRepo.AddNewTransactions(transactions)
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
