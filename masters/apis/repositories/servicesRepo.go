package repositories

import "github.com/inact25/E-WarungApi/masters/apis/models"

type ServiceRepositories interface {
	GetAllCategories() ([]*models.CategoryModels, error)
	AddNewCategory(category *models.CategoryModels) (string, error)
	AddNewCategoryPrice(categories *models.CategoryModels) (string, error)
}
