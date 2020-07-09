package usecases

import (
	"errors"
	"github.com/inact25/E-WarungApi/masters/apis/models"
	"github.com/inact25/E-WarungApi/masters/apis/repositories"
	"github.com/inact25/E-WarungApi/utils"
	"log"
)

type CategoryUseCaseImpl struct {
	categoryRepo repositories.CategoriesRepositories
}

func (s CategoryUseCaseImpl) GetAllCategoriesByStatus(status string) ([]*models.CategoriesModels, error) {
	if utils.IsStatusValid(status) != true {
		return nil, errors.New("ERROR")
	}
	menu, err := s.categoryRepo.GetAllCategoriesByStatus(status)
	log.Println("U : ", status)
	if err != nil {
		return nil, err
	}
	return menu, nil
}

func (s CategoryUseCaseImpl) GetAllCategories() ([]*models.CategoriesModels, error) {
	menu, err := s.categoryRepo.GetAllCategories()
	log.Println("U : ", menu)
	if err != nil {
		return nil, err
	}
	return menu, nil
}

func (s CategoryUseCaseImpl) AddNewCategories(day string, categories *models.CategoriesModels) (string, error) {
	category, err := s.categoryRepo.AddNewCategories(day, categories)
	if err != nil {
		return "", err
	}
	return category, nil
}

func (s CategoryUseCaseImpl) UpdateCategories(categories *models.CategoriesModels) (string, error) {
	log.Println("U :", categories)
	service, err := s.categoryRepo.UpdateCategories(categories)
	if err != nil {
		return "", err
	}
	return service, nil
}

func (s CategoryUseCaseImpl) UpdateCategoriesPrice(day string, categories *models.CategoriesModels) (string, error) {
	log.Println("U :", categories)
	product, err := s.categoryRepo.UpdateCategoriesPrice(day, categories)
	if err != nil {
		return "", err
	}
	return product, nil
}

func (s CategoryUseCaseImpl) DeleteCategories(categoriesID string) (string, error) {
	product, err := s.categoryRepo.DeleteCategories(categoriesID)
	if err != nil {
		return "", err
	}
	return product, nil
}

func InitCategoryUseCase(categories repositories.CategoriesRepositories) CategoryUseCases {
	return &CategoryUseCaseImpl{categories}
}
