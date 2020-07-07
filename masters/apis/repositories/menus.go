package repositories

import (
	"database/sql"
	"github.com/inact25/E-WarungApi/masters/apis/models"
	"log"
)

type MenuRepoImpl struct {
	db *sql.DB
}

func (s MenuRepoImpl) UpdateMenu(menu *models.MenuModels) (string, error) {
	tx, err := s.db.Begin()
	if err != nil {
		log.Fatalf("%v", err)
		return "", err
	}
	stmt, err := s.db.Prepare("update menu set menuDesc = ?, menuStock = ? where menutID = ?")
	if err != nil {
		tx.Rollback()
		log.Fatalf("%v", err)
		return "", err
	}
	defer stmt.Close()

	if _, err := stmt.Exec(menu.MenuDesc, menu.MenuStock); err != nil {
		tx.Rollback()
		log.Fatalf("%v", err)
		return "", err
	}
	return "", tx.Commit()
}

func (s MenuRepoImpl) DeleteMenu(menuID string) (string, error) {
	tx, err := s.db.Begin()
	if err != nil {
		return "", err
	}
	stmt, err := s.db.Prepare("delete from menu where menuID = ?")
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

func (s MenuRepoImpl) AddNewMenu(menus *models.MenuModels) (string, error) {
	tx, err := s.db.Begin()
	if err != nil {
		return "", err
	}
	stmt, err := s.db.Prepare("insert into menu values (?,?,?)")
	if err != nil {
		tx.Rollback()
		return "", err
	}
	defer stmt.Close()
	if _, err := stmt.Exec(menus.MenuID, menus.MenuDesc, menus.MenuStock); err != nil {
		tx.Rollback()
		return "", err
	}
	return "", tx.Commit()
}

func (s MenuRepoImpl) GetAllMenu() ([]*models.MenuModels, error) {
	dataMenus := []*models.MenuModels{}
	query := "select m.menuID, m.menuDesc, max(p.price) as price, m.menuStock from menu m left join price p on m.menuID = p.priceID group by menuID;"
	data, err := s.db.Query(query)
	log.Println("R : ", data)
	if err != nil {
		return nil, err
	}
	for data.Next() {
		menu := models.MenuModels{}
		err := data.Scan(&menu.MenuID, &menu.MenuDesc, &menu.MenuDesc, &menu.MenuPrice, &menu.MenuStock)
		if err != nil {
			return nil, err

		}
		dataMenus = append(dataMenus, &menu)
		log.Println("R : ", dataMenus)
	}
	return dataMenus, nil
}

func InitMenuRepoImpl(db *sql.DB) MenuRepositories {
	return &MenuRepoImpl{db}

}
