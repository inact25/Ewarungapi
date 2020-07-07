package repositories

import (
	"database/sql"
	"github.com/inact25/E-WarungApi/masters/apis/models"

	"log"
)

type ServicesRepoImpl struct {
	db *sql.DB
}

func (c ServicesRepoImpl) GetAllServices() ([]*models.ServicesModels, error) {
	dataServices := []*models.ServicesModels{}
	query := "select s.servicesID, s.servicesDesc, max(sp.price) as price from services s inner join servicesprice sp on s.servicesID = sp.priceID group by s.servicesID"
	data, err := c.db.Query(query)
	if err != nil {
		return nil, err
	}
	for data.Next() {
		services := models.ServicesModels{}
		err := data.Scan(&services.ServicesID, &services.ServicesDate, &services.ServicesPrice)
		if err != nil {
			return nil, err
		}
		dataServices = append(dataServices, &services)
	}
	return dataServices, nil
}

func (c ServicesRepoImpl) AddNewServices(services *models.ServicesModels) (string, error) {
	tx, err := c.db.Begin()
	if err != nil {
		return "", err
	}
	stmt, err := c.db.Prepare("insert into services values (?,?)")
	if err != nil {
		tx.Rollback()
		return "", err
	}
	defer stmt.Close()

	if _, err := stmt.Exec(services.ServicesID, services.ServicesDesc); err != nil {
		tx.Rollback()
		log.Fatalf("%v", err)
		return "", err
	}
	return "", tx.Commit()
}

func (c ServicesRepoImpl) AddNewServicePrice(services *models.ServicesModels) (string, error) {
	tx, err := c.db.Begin()
	if err != nil {
		return "", err
	}
	stmt, err := c.db.Prepare("insert into servicesprice values (?,?,?)")
	if err != nil {
		tx.Rollback()
		return "", err
	}
	defer stmt.Close()

	if _, err := stmt.Exec(services.ServicesID, services.ServicesDate, services.ServicesPrice); err != nil {
		tx.Rollback()
		log.Fatalf("%v", err)
		return "", err
	}
	return "", tx.Commit()
}
func InitServiceRepoImpl(db *sql.DB) CategoryRepositories {
	return &CategoriesRepoImpl{db}

}
