package usecases

import "github.com/inact25/E-WarungApi/masters/apis/models"

type CategoryUseCases interface {
	GetAllCategories() ([]*models.CategoryModels, error)
	GetAllCategoriesPrice() ([]*models.CategoryPriceModels, error)
	AddNewCategory(category *models.CategoryModels) (string, error)
	AddNewCategoryPrice(day string, category *models.CategoryPriceModels) (string, error)
}
