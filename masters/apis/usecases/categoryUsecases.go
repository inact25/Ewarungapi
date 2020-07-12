package usecases

import "github.com/inact25/E-WarungApi/masters/apis/models"

type CategoryUseCases interface {
	GetAllCategories() ([]*models.CategoriesModels, error)
	GetAllCategoriesByStatus(status string) ([]*models.CategoriesModels, error)
	AddNewCategories(categories *models.CategoriesModels) (string, error)
	UpdateCategories(categories *models.CategoriesModels) (string, error)
	UpdateCategoriesPrice(day string, categories *models.CategoriesModels) (string, error)
	DeleteCategories(categoriesID string) (string, error)
}
