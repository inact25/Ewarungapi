package repositories

import (
	"database/sql"
	"github.com/inact25/E-WarungApi/masters/apis/models"
	"log"
)

type TransactionRepoImpll struct {
	db *sql.DB
}

func (t TransactionRepoImpll) UpdateTransaction(transactions *models.TransactionModels) (string, error) {

	tx, err := t.db.Begin()
	if err != nil {
		log.Fatalf("%v", err)
		return "", err
	}
	stmt, err := t.db.Prepare("update transaction set serviceID = ?, menuID = ?, categoryID = ?, Qty = ? where transactionID = ?")
	if err != nil {
		tx.Rollback()
		log.Fatalf("%v", err)
		return "", err
	}
	defer stmt.Close()

	if _, err := stmt.Exec(transactions.ServiceDesc, transactions.MenuDesc, transactions.CategoryDesc, transactions.Quantity, transactions.TransactionID); err != nil {
		tx.Rollback()
		log.Fatalf("%v", err)
		return "", err
	}
	return "", tx.Commit()
}

func (t TransactionRepoImpll) AddTransaction(day string, transactions *models.TransactionModels) (string, error) {
	tx, err := t.db.Begin()
	if err != nil {
		return "", err
	}
	stmt, err := t.db.Prepare("insert into transaction values (uuid(),?,?,?,?,?) ")
	if err != nil {
		tx.Rollback()
		return "", err
	}
	defer stmt.Close()

	if _, err := stmt.Exec(day, transactions.ServiceDesc, transactions.MenuDesc, transactions.CategoryDesc, transactions.Quantity); err != nil {
		tx.Rollback()
		return "", err
	}
	return "", tx.Commit()
}

func (t TransactionRepoImpll) DeleteTransaction(transactionID string) (string, error) {
	log.Println("r : ", transactionID)
	tx, err := t.db.Begin()
	if err != nil {
		return "", err
	}
	stmt, err := t.db.Prepare("delete from transaction where transactionID = ?")
	if err != nil {
		tx.Rollback()
		return "", err
	}
	defer stmt.Close()
	if _, err := stmt.Exec(transactionID); err != nil {
		tx.Rollback()
		return "", err
	}
	return "", tx.Commit()
}

func (t TransactionRepoImpll) GetAllTransaction() ([]*models.TransactionModels, error) {
	dataTransaction := []*models.TransactionModels{}
	query := "select t.transactionID, t.transactionDate, m.menuDesc, p.price,t.Qty, c.categoryDesc, cp.price as 'Favor Price', s.servicesDesc, sp.Price as 'Services Price', (p.price*t.Qty)+cp.price+sp.Price as'Sub Total'  from transaction t  inner join menu m on t.menuID = m.menuID inner join price p on p.priceID = m.menuID inner join category c on c.categoryID = t.categoryID inner join categoriesprice cp on cp.priceID = c.categoryID inner join services s on s.servicesID = t.servicesID inner join servicesprice sp on sp.PriceID = s.servicesID group by t.transactionID;"
	data, err := t.db.Query(query)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	for data.Next() {
		transaction := models.TransactionModels{}
		err := data.Scan(&transaction.TransactionID, &transaction.TransactionDate,
			&transaction.MenuDesc, &transaction.MenuPrice, &transaction.Quantity,
			&transaction.CategoryDesc, &transaction.FavorPrice, &transaction.ServiceDesc, &transaction.ServicePrice, &transaction.SubTotal)
		if err != nil {
			log.Fatal(err)
			return nil, err
		}
		dataTransaction = append(dataTransaction, &transaction)
	}
	return dataTransaction, nil

}

func (t TransactionRepoImpll) GetDailyTransaction(date string) ([]*models.TransactionModels, error) {
	dataTransaction := []*models.TransactionModels{}
	query, err := t.db.Query(`select t.transactionID, t.transactionDate, m.menuDesc, p.price,t.Qty, c.categoryDesc, cp.price as 'Favor Price', s.servicesDesc, sp.Price as 'Services Price', (p.price*t.Qty)+cp.price+sp.Price as'Sub Total'  from transaction t  inner join menu m on t.menuID = m.menuID inner join price p on p.priceID = m.menuID inner join category c on c.categoryID = t.categoryID inner join categoriesprice cp on cp.priceID = c.categoryID inner join services s on s.servicesID = t.servicesID inner join servicesprice sp on sp.PriceID = s.servicesID where t.transactionDate like ? group by t.transactionID`, date)

	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	for query.Next() {
		transaction := models.TransactionModels{}
		err := query.Scan(&transaction.TransactionID, &transaction.TransactionDate,
			&transaction.MenuDesc, &transaction.MenuPrice, &transaction.Quantity,
			&transaction.CategoryDesc, &transaction.FavorPrice, &transaction.ServiceDesc, &transaction.ServicePrice, &transaction.SubTotal)
		if err != nil {
			log.Fatal(err)
			return nil, err
		}
		dataTransaction = append(dataTransaction, &transaction)
	}
	return dataTransaction, nil

}

func InitTransactionRepoImpl(db *sql.DB) TransactionRepositories {
	return &TransactionRepoImpll{db}

}
