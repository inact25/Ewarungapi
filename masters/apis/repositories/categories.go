package repositories

import (
	"database/sql"
	"github.com/inact25/E-WarungApi/masters/apis/models"

	"log"
)

type CategoriesRepoImpl struct {
	db *sql.DB
}

func (c CategoriesRepoImpl) GetAllCategories() ([]*models.CategoryModels, error) {
	dataCategories := []*models.CategoryModels{}
	//query := "select * from category"
	query := `select c.categoryID, c.categoryDesc, ifnull(max(cp.price),"Empty") as price from category c left join categoriesprice cp on c.categoryID = cp.priceID group by c.categoryID;`
	data, err := c.db.Query(query)
	if err != nil {
		return nil, err
	}
	for data.Next() {
		categories := models.CategoryModels{}
		err := data.Scan(&categories.CategoryID, &categories.CategoryDesc, &categories.CategoryPrice)
		if err != nil {
			log.Fatal(err)
			return nil, err

		}
		dataCategories = append(dataCategories, &categories)
	}
	return dataCategories, nil
}

func (c CategoriesRepoImpl) AddNewCategory(categories *models.CategoryModels) (string, error) {
	tx, err := c.db.Begin()
	if err != nil {
		return "", err
	}
	stmt, err := c.db.Prepare("insert into category values (?,?)")
	if err != nil {
		tx.Rollback()
		return "", err
	}
	defer stmt.Close()

	if _, err := stmt.Exec(categories.CategoryID, categories.CategoryDesc); err != nil {
		tx.Rollback()
		log.Fatalf("%v", err)
		return "", err
	}
	return "", tx.Commit()
}

func (c CategoriesRepoImpl) GetAllCategoriesPrice() ([]*models.CategoryPriceModels, error) {
	dataCategories := []*models.CategoryPriceModels{}
	//query := "select * from category"
	query := `select * from categoriesprice`
	data, err := c.db.Query(query)
	if err != nil {
		return nil, err
	}
	for data.Next() {
		categories := models.CategoryPriceModels{}
		err := data.Scan(&categories.PriceID, &categories.PriceDate, &categories.Price)
		if err != nil {
			log.Fatal(err)
			return nil, err

		}
		dataCategories = append(dataCategories, &categories)
	}
	return dataCategories, nil
}

func (c CategoriesRepoImpl) AddNewCategoryPrice(day string, categories *models.CategoryPriceModels) (string, error) {
	tx, err := c.db.Begin()
	if err != nil {
		return "", err
	}
	stmt, err := c.db.Prepare("insert into categoriesprice values (?,?,?)")
	if err != nil {
		tx.Rollback()
		return "", err
	}
	defer stmt.Close()

	if _, err := stmt.Exec(categories.PriceID, day, categories.Price); err != nil {
		tx.Rollback()
		log.Fatalf("%v", err)
		return "", err
	}
	return "", tx.Commit()
}
func InitCategoryRepoImpl(db *sql.DB) CategoryRepositories {
	return &CategoriesRepoImpl{db}

}
