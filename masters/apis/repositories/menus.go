package repositories

import (
	"database/sql"
	"github.com/inact25/E-WarungApi/masters/apis/models"
	"log"
)

type MenuRepoImpl struct {
	db *sql.DB
}

func (s MenuRepoImpl) GetAllMenu() ([]*models.MenuModels, error) {
	dataMenus := []*models.MenuModels{}
	query := GetAllMenusQuery
	data, err := s.db.Query(query)
	log.Println("R : ", data)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	for data.Next() {
		menu := models.MenuModels{}
		err := data.Scan(&menu.MenuID, &menu.MenuDesc, &menu.MenuPrice, &menu.MenuStock, &menu.MenuStatus, &menu.PriceDate)
		if err != nil {
			log.Fatal(err)
			return nil, err

		}
		dataMenus = append(dataMenus, &menu)
		log.Println("R : ", dataMenus)
	}
	return dataMenus, nil
}

func (s MenuRepoImpl) GetAllMenuByStatus(status string) ([]*models.MenuModels, error) {
	log.Println("R : ", status)
	dataMenus := []*models.MenuModels{}
	query := GetAllMenusByStatusQuery
	data, err := s.db.Query(query, status)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	for data.Next() {
		menu := models.MenuModels{}
		err := data.Scan(&menu.MenuID, &menu.MenuDesc, &menu.MenuPrice, &menu.MenuStock, &menu.MenuStatus, &menu.PriceDate)
		if err != nil {
			log.Fatal(err)
			return nil, err

		}
		dataMenus = append(dataMenus, &menu)
		log.Println("R : ", dataMenus)
	}
	return dataMenus, nil
}

func (s MenuRepoImpl) GetAllMenuPrices() ([]*models.MenuPriceModels, error) {
	dataMenus := []*models.MenuPriceModels{}
	query := GetAllMenusPrice
	data, err := s.db.Query(query)
	log.Println("R : ", data)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	for data.Next() {
		menu := models.MenuPriceModels{}
		err := data.Scan(&menu.MenuID, &menu.MenuDesc, &menu.MenuPrice, &menu.PriceDate)
		if err != nil {
			log.Fatal(err)
			return nil, err

		}
		dataMenus = append(dataMenus, &menu)
		log.Println("R : ", dataMenus)
	}
	return dataMenus, nil
}

func (s MenuRepoImpl) AddNewMenu(day string, menus *models.MenuModels) (string, error) {
	tx, err := s.db.Begin()
	if err != nil {
		log.Fatal(err)
		return "", err
	}
	addMenu, err := s.db.Prepare(AddNewMenuQuery)
	if err != nil {
		tx.Rollback()
		return "", err
	}
	defer addMenu.Close()
	if _, err := addMenu.Exec(menus.MenuID, menus.MenuDesc, menus.MenuStock, &menus.MenuStatus); err != nil {
		tx.Rollback()
		return "", err
	}
	addPrice, err := s.db.Prepare(AddNewMenuPricesQuery)
	if err != nil {
		tx.Rollback()
		return "", err
	}
	defer addPrice.Close()
	if _, err := addPrice.Exec(menus.MenuID, day, menus.MenuPrice); err != nil {
		tx.Rollback()
		return "", err
	}
	return "", tx.Commit()
}

func (s MenuRepoImpl) UpdateMenu(menu *models.MenuModels) (string, error) {

	log.Println("R :", menu)
	tx, err := s.db.Begin()
	if err != nil {
		log.Fatal(err)
		return "", err
	}
	putMenu, err := s.db.Prepare(UpdateMenuQuery)
	if err != nil {
		tx.Rollback()
		return "", err
	}
	defer putMenu.Close()
	if _, err := putMenu.Exec(menu.MenuDesc, menu.MenuStock, menu.MenuStatus, menu.MenuID); err != nil {
		tx.Rollback()
		return "", err
	}
	return "", tx.Commit()
}

func (s MenuRepoImpl) UpdateMenuPrice(day string, menu *models.MenuPriceModels) (string, error) {

	log.Println("R :", menu)
	tx, err := s.db.Begin()
	if err != nil {
		log.Fatal(err)
		return "", err
	}
	putMenu, err := s.db.Prepare(UpdateMenuPriceQuery)
	if err != nil {
		tx.Rollback()
		return "", err
	}
	defer putMenu.Close()
	if _, err := putMenu.Exec(menu.MenuID, day, menu.MenuPrice); err != nil {
		tx.Rollback()
		return "", err
	}
	return "", tx.Commit()
}

func (s MenuRepoImpl) DeleteMenu(menuID string) (string, error) {
	tx, err := s.db.Begin()
	if err != nil {
		return "", err
	}
	stmt, err := s.db.Prepare(DeleteMenuQuery)
	if err != nil {
		tx.Rollback()
		return "", err
	}
	defer stmt.Close()
	if _, err := stmt.Exec(menuID); err != nil {
		tx.Rollback()
		return "", err
	}
	return "", tx.Commit()
}

func InitMenuRepoImpl(db *sql.DB) MenuRepositories {
	return &MenuRepoImpl{db}

}
