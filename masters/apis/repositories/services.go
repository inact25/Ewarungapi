package repositories

import (
	"database/sql"
	"github.com/inact25/E-WarungApi/masters/apis/models"

	"log"
)

type ServicesRepoImpl struct {
	db *sql.DB
}

func (s ServicesRepoImpl) GetAllServices() ([]*models.ServicesModels, error) {
	dataServices := []*models.ServicesModels{}
	query := GetAllServices
	data, err := s.db.Query(query)
	log.Println("R : ", data)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	for data.Next() {
		services := models.ServicesModels{}
		err := data.Scan(&services.ServicesID, &services.ServicesDesc, &services.ServicePrice, &services.ServicesStatus, &services.PriceDate)
		if err != nil {
			log.Fatal(err)
			return nil, err

		}
		dataServices = append(dataServices, &services)
		log.Println("R : ", dataServices)
	}
	return dataServices, nil
}

func (s ServicesRepoImpl) GetAllServicesByStatus(status string) ([]*models.ServicesModels, error) {
	log.Println("R : ", status)
	dataServices := []*models.ServicesModels{}
	query := GetAllServicesByStatus
	data, err := s.db.Query(query, status)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	for data.Next() {
		services := models.ServicesModels{}
		err := data.Scan(&services.ServicesID, &services.ServicesDesc, &services.ServicePrice, &services.ServicesStatus, &services.PriceDate)
		if err != nil {
			log.Fatal(err)
			return nil, err

		}
		dataServices = append(dataServices, &services)
		log.Println("R : ", dataServices)
	}
	return dataServices, nil
}

func (s ServicesRepoImpl) AddNewServices(day string, services *models.ServicesModels) (string, error) {
	tx, err := s.db.Begin()
	if err != nil {
		log.Fatal(err)
		return "", err
	}
	addService, err := s.db.Prepare(AddNewServicesQuery)
	if err != nil {
		tx.Rollback()
		return "", err
	}
	defer addService.Close()
	if _, err := addService.Exec(services.ServicesID, services.ServicesDesc, services.ServicesStatus); err != nil {
		tx.Rollback()
		return "", err
	}
	addPrice, err := s.db.Prepare(AddNewServicesPricesQuery)
	if err != nil {
		tx.Rollback()
		return "", err
	}
	defer addPrice.Close()
	if _, err := addPrice.Exec(services.ServicesID, day, services.PriceDate); err != nil {
		tx.Rollback()
		return "", err
	}
	return "", tx.Commit()
}

func InitServiceRepoImpl(db *sql.DB) ServiceRepositories {
	return &ServicesRepoImpl{db}

}
