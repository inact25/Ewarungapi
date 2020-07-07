package usecases

import (
	"github.com/inact25/E-WarungApi/masters/apis/models"
	"github.com/inact25/E-WarungApi/masters/apis/repositories"
)

type CategoryUseCaseImpl struct {
	categoryRepo repositories.CategoryRepositories
}

func (s CategoryUseCaseImpl) GetAllCategories() ([]*models.CategoryModels, error) {
	category, err := s.categoryRepo.GetAllCategories()
	if err != nil {
		return nil, err
	}
	return category, nil
}

func (s CategoryUseCaseImpl) GetAllCategoriesPrice() ([]*models.CategoryPriceModels, error) {
	category, err := s.categoryRepo.GetAllCategoriesPrice()
	if err != nil {
		return nil, err
	}
	return category, nil
}

func (s CategoryUseCaseImpl) AddNewCategory(categories *models.CategoryModels) (string, error) {
	category, err := s.categoryRepo.AddNewCategory(categories)
	if err != nil {
		return "", err
	}
	return category, nil
}

func (s CategoryUseCaseImpl) AddNewCategoryPrice(day string, categories *models.CategoryPriceModels) (string, error) {
	category, err := s.categoryRepo.AddNewCategoryPrice(day, categories)
	if err != nil {
		return "", err
	}
	return category, nil

}

func InitCategoryUseCase(categoryRepo repositories.CategoryRepositories) CategoryUseCases {
	return &CategoryUseCaseImpl{categoryRepo}
}
