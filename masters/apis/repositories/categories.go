package repositories

import (
	"database/sql"
	"github.com/inact25/E-WarungApi/masters/apis/models"

	"log"
)

type CategoriesRepoImpl struct {
	db *sql.DB
}

func (s CategoriesRepoImpl) GetAllCategories() ([]*models.CategoriesModels, error) {
	dataCategories := []*models.CategoriesModels{}
	query := GetAllCategories
	data, err := s.db.Query(query)
	log.Println("R : ", data)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	for data.Next() {
		categories := models.CategoriesModels{}
		err := data.Scan(&categories.CategoriesID, &categories.CategoriesDesc, &categories.CategoriesPrice, &categories.CategoriesStatus, &categories.PriceDate)
		if err != nil {
			log.Fatal(err)
			return nil, err

		}
		dataCategories = append(dataCategories, &categories)
		log.Println("R : ", dataCategories)
	}
	return dataCategories, nil
}

func (s CategoriesRepoImpl) GetAllCategoriesByStatus(status string) ([]*models.CategoriesModels, error) {
	log.Println("R : ", status)
	dataCategories := []*models.CategoriesModels{}
	query := GetAllCategoriesByStatus
	data, err := s.db.Query(query, status)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	for data.Next() {
		categories := models.CategoriesModels{}
		err := data.Scan(&categories.CategoriesID, &categories.CategoriesDesc, &categories.CategoriesPrice, &categories.CategoriesStatus, &categories.PriceDate)
		if err != nil {
			log.Fatal(err)
			return nil, err

		}
		dataCategories = append(dataCategories, &categories)
		log.Println("R : ", dataCategories)
	}
	return dataCategories, nil
}

func (s CategoriesRepoImpl) AddNewCategories(categories *models.CategoriesModels) (string, error) {
	tx, err := s.db.Begin()
	if err != nil {
		log.Fatal(err)
		return "", err
	}
	addCategories, err := s.db.Prepare(AddNewCategoriesQuery)
	if err != nil {
		tx.Rollback()
		return "", err
	}
	defer addCategories.Close()
	if _, err := addCategories.Exec(categories.CategoriesID, categories.CategoriesDesc, categories.CategoriesStatus); err != nil {
		tx.Rollback()
		return "", err
	}
	addPrice, err := s.db.Prepare(AddNewCategoriesPricesQuery)
	if err != nil {
		tx.Rollback()
		return "", err
	}
	defer addPrice.Close()
	if _, err := addPrice.Exec(categories.CategoriesID, categories.PriceDate, categories.CategoriesPrice); err != nil {
		tx.Rollback()
		return "", err
	}
	return "", tx.Commit()
}

func (s CategoriesRepoImpl) UpdateCategories(categories *models.CategoriesModels) (string, error) {
	log.Println("R :", categories)
	tx, err := s.db.Begin()
	if err != nil {
		log.Fatal(err)
		return "", err
	}
	putCategories, err := s.db.Prepare(UpdateCategoriesQuery)
	if err != nil {
		tx.Rollback()
		return "", err
	}
	defer putCategories.Close()
	if _, err := putCategories.Exec(categories.CategoriesDesc, categories.CategoriesStatus, categories.CategoriesID); err != nil {
		tx.Rollback()
		return "", err
	}
	return "", tx.Commit()
}

func (s CategoriesRepoImpl) UpdateCategoriesPrice(day string, categories *models.CategoriesModels) (string, error) {
	log.Println("price : ", categories.CategoriesPrice)
	log.Println("id :", categories.CategoriesID)
	log.Println("R :", day)
	tx, err := s.db.Begin()
	if err != nil {
		log.Fatal(err)
		return "", err
	}
	putCategories, err := s.db.Prepare(UpdateCategoriesPriceQuery)
	if err != nil {
		tx.Rollback()
		return "", err
	}
	defer putCategories.Close()
	if _, err := putCategories.Exec(categories.CategoriesID, day, categories.CategoriesPrice); err != nil {
		tx.Rollback()
		return "", err
	}
	return "", tx.Commit()

}

func (s CategoriesRepoImpl) DeleteCategories(categoriesID string) (string, error) {
	tx, err := s.db.Begin()
	if err != nil {
		log.Fatal(err)
		return "", err
	}
	stmt, err := s.db.Prepare(DeleteCategoriesQuery)
	if err != nil {
		tx.Rollback()
		return "", err
	}
	defer stmt.Close()
	if _, err := stmt.Exec(categoriesID); err != nil {
		tx.Rollback()
		return "", err
	}
	return "", tx.Commit()
}

func InitCategoriesRepoImpl(db *sql.DB) CategoriesRepositories {
	return &CategoriesRepoImpl{db}
}
