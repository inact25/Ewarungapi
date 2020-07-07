package usecases

import (
	"github.com/inact25/E-WarungApi/masters/apis/models"
	"github.com/inact25/E-WarungApi/masters/apis/repositories"
)

type TransactionUseCaseImpl struct {
	transactionRepo repositories.TransactionRepositories
}

func (s TransactionUseCaseImpl) AddTransaction(day string, transactions *models.TransactionModels) (string, error) {
	transaction, err := s.transactionRepo.AddTransaction(day, transactions)
	if err != nil {
		return "", err
	}
	return transaction, nil
}

func (s TransactionUseCaseImpl) UpdateTransaction(transactions *models.TransactionModels) (string, error) {
	transaction, err := s.transactionRepo.UpdateTransaction(transactions)
	if err != nil {
		return "", err
	}
	return transaction, nil
}

func (s TransactionUseCaseImpl) DeleteTransaction(transactionID string) (string, error) {
	transaction, err := s.transactionRepo.DeleteTransaction(transactionID)
	if err != nil {
		return "", err
	}
	return transaction, nil
}

func (s TransactionUseCaseImpl) GetAllTransaction() ([]*models.TransactionModels, error) {
	transaction, err := s.transactionRepo.GetAllTransaction()
	if err != nil {
		return nil, err
	}
	return transaction, nil
}

func (s TransactionUseCaseImpl) GetDailyTransaction(date string) ([]*models.TransactionModels, error) {
	transaction, err := s.transactionRepo.GetDailyTransaction(date)
	if err != nil {
		return nil, err
	}
	return transaction, nil
}

func InitTransactionUseCase(transactionRepo repositories.TransactionRepositories) TransactionUseCases {
	return &TransactionUseCaseImpl{transactionRepo}
}
