package repositories

import (
	"database/sql"
	"github.com/inact25/E-WarungApi/masters/apis/models"
	"log"
)

type TransactionRepoImpll struct {
	db *sql.DB
}

func (s TransactionRepoImpll) GetAllTransactions() ([]*models.TransactionModels, error) {
	var dataTransactions []*models.TransactionModels
	query := GetAllTransactionsQuery
	data, err := s.db.Query(query)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	for data.Next() {
		transactions := models.TransactionModels{}
		err := data.Scan(&transactions.TransactionID, &transactions.TransactionDate, &transactions.MenuDesc,
			&transactions.MenuPrice, &transactions.CategoryDesc, &transactions.CategoryPrice, &transactions.Qty,
			&transactions.ServicesDesc, &transactions.ServicePrice, &transactions.SubTotal)
		if err != nil {
			log.Fatal(err)
			return nil, err

		}
		dataTransactions = append(dataTransactions, &transactions)
	}
	return dataTransactions, nil
}

func (s TransactionRepoImpll) AddNewTransactions(transactions *models.TransactionModels) (string, error) {
	menuPrice := GetLatestMenuPriceByIDQuery
	categoryPrice := GetLatestCategoryPriceByIDQuery
	servicesPrice := GetLatestServicePriceByIDQuery

	errMenu := s.db.QueryRow(menuPrice, transactions.MenuDesc).Scan(&transactions.MenuPrice)
	if errMenu != nil {
		log.Fatal(errMenu)
		return "", errMenu
	}
	errCategory := s.db.QueryRow(categoryPrice, transactions.CategoryDesc).Scan(&transactions.CategoryPrice)
	if errCategory != nil {
		log.Fatal(errCategory)
		return "", errCategory
	}
	errServices := s.db.QueryRow(servicesPrice, transactions.ServicesDesc).Scan(&transactions.ServicePrice)
	if errServices != nil {
		log.Fatal(errServices)
		return "", errServices
	}

	tx, err := s.db.Begin()
	if err != nil {
		log.Fatal(err)
		return "", err
	}
	addTransactions, err := s.db.Prepare(AddNewTransactions)
	if err != nil {
		tx.Rollback()
		return "", err
	}
	defer addTransactions.Close()
	if _, err := addTransactions.Exec(transactions.TransactionDate, transactions.ServicesDesc, transactions.ServicePrice, transactions.MenuDesc,
		transactions.MenuPrice, transactions.CategoryDesc, transactions.CategoryPrice, transactions.Qty); err != nil {
		tx.Rollback()
		return "", err
	}
	return "", tx.Commit()
}

func (s TransactionRepoImpll) UpdateTransactions(transactions *models.TransactionModels) (string, error) {

	menuPrice := GetLatestMenuPriceByIDQuery
	categoryPrice := GetLatestCategoryPriceByIDQuery
	servicesPrice := GetLatestServicePriceByIDQuery

	errMenu := s.db.QueryRow(menuPrice, transactions.MenuDesc).Scan(&transactions.MenuPrice)
	if errMenu != nil {
		log.Fatal("1st", errMenu)
		return "", errMenu
	}
	errCategory := s.db.QueryRow(categoryPrice, transactions.CategoryDesc).Scan(&transactions.CategoryPrice)
	if errCategory != nil {
		log.Fatal(errCategory)
		return "", errCategory
	}
	errServices := s.db.QueryRow(servicesPrice, transactions.ServicesDesc).Scan(&transactions.ServicePrice)
	if errServices != nil {
		log.Fatal(errServices)
		return "", errServices
	}

	tx, err := s.db.Begin()
	if err != nil {
		log.Fatal(err)
		return "", err
	}
	updateTransactions, err := s.db.Prepare(UpdateTransactions)
	if err != nil {
		tx.Rollback()
		return "", err
	}
	defer updateTransactions.Close()
	if _, err := updateTransactions.Exec(transactions.ServicesDesc, transactions.ServicePrice, transactions.MenuDesc,
		transactions.MenuPrice, transactions.CategoryDesc, transactions.CategoryPrice, transactions.Qty, transactions.TransactionID); err != nil {
		tx.Rollback()
		return "", err
	}

	return "", tx.Commit()
}

func InitTransactionRepoImpl(db *sql.DB) TransactionRepositories {
	return &TransactionRepoImpll{db}

}
