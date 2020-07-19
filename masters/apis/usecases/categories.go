package usecases

import (
	"errors"
	"github.com/inact25/E-WarungApi/masters/apis/models"
	"github.com/inact25/E-WarungApi/masters/apis/repositories"
	"github.com/inact25/E-WarungApi/utils/validation"
)

type CategoryUseCaseImpl struct {
	categoryRepo repositories.CategoriesRepositories
}

func (s CategoryUseCaseImpl) GetAllCategoriesByStatus(status string) ([]*models.CategoriesModels, error) {
	if validation.IsStatusValid(status) != true {
		return nil, errors.New("ERROR")
	}
	menu, err := s.categoryRepo.GetAllCategoriesByStatus(status)
	if err != nil {
		return nil, err
	}
	return menu, nil
}

func (s CategoryUseCaseImpl) GetAllCategories() ([]*models.CategoriesModels, error) {
	menu, err := s.categoryRepo.GetAllCategories()
	if err != nil {
		return nil, err
	}
	return menu, nil
}

func (s CategoryUseCaseImpl) AddNewCategories(categories *models.CategoriesModels) (string, error) {
	err := validation.CheckEmpty(categories.CategoriesID, categories.CategoriesDesc, categories.CategoriesStatus, categories.CategoriesPrice)
	if err != nil {
		return "", err
	}
	err = validation.CheckInt(categories.CategoriesPrice)
	if err != nil {
		return "", err
	}
	if validation.IsStatusValid(categories.CategoriesStatus) != true {
		return "", errors.New("Status not valid")
	}
	category, err := s.categoryRepo.AddNewCategories(categories)
	if err != nil {
		return "", err
	}
	return category, nil
}

func (s CategoryUseCaseImpl) UpdateCategories(categories *models.CategoriesModels) (string, error) {
	err := validation.CheckEmpty(categories.CategoriesID, categories.CategoriesDesc, categories.CategoriesStatus)
	if err != nil {
		return "", err
	}
	if validation.IsStatusValid(categories.CategoriesStatus) != true {
		return "", errors.New("Status not valid")
	}
	service, err := s.categoryRepo.UpdateCategories(categories)
	if err != nil {
		return "", err
	}
	return service, nil
}

func (s CategoryUseCaseImpl) UpdateCategoriesPrice(day string, categories *models.CategoriesModels) (string, error) {
	err := validation.CheckEmpty(categories.CategoriesID, categories.CategoriesPrice)
	if err != nil {
		return "", err
	}
	err = validation.CheckInt(categories.CategoriesPrice)
	if err != nil {
		return "", err
	}
	product, err := s.categoryRepo.UpdateCategoriesPrice(day, categories)
	if err != nil {
		return "", err
	}
	return product, nil
}

func (s CategoryUseCaseImpl) DeleteCategories(categoriesID string) (string, error) {
	err := validation.CheckEmpty(categoriesID)
	if err != nil {
		return "", err
	}
	product, err := s.categoryRepo.DeleteCategories(categoriesID)
	if err != nil {
		return "", err
	}
	return product, nil
}

func InitCategoryUseCase(categories repositories.CategoriesRepositories) CategoryUseCases {
	return &CategoryUseCaseImpl{categories}
}
