package repositories

import "github.com/inact25/E-WarungApi/masters/apis/models"

type CategoryRepositories interface {
	GetAllCategories() ([]*models.CategoryModels, error)
	GetAllCategoriesPrice() ([]*models.CategoryPriceModels, error)
	AddNewCategory(category *models.CategoryModels) (string, error)
	AddNewCategoryPrice(day string, categories *models.CategoryPriceModels) (string, error)
}
