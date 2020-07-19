package usecases

import (
	"errors"
	"github.com/inact25/E-WarungApi/masters/apis/models"
	"github.com/inact25/E-WarungApi/masters/apis/repositories"
	"github.com/inact25/E-WarungApi/utils/validation"
)

type MenuUseCaseImpl struct {
	menuRepo repositories.MenuRepositories
}

func (s MenuUseCaseImpl) GetAllMenu() ([]*models.MenuModels, error) {
	menu, err := s.menuRepo.GetAllMenu()
	if err != nil {
		return nil, err
	}
	return menu, nil
}

func (s MenuUseCaseImpl) GetAllMenuByStatus(status string) ([]*models.MenuModels, error) {
	if validation.IsStatusValid(status) != true {
		return nil, errors.New("ERROR")
	}
	menu, err := s.menuRepo.GetAllMenuByStatus(status)
	if err != nil {
		return nil, err
	}

	return menu, nil
}

func (s MenuUseCaseImpl) GetAllMenuPrices() ([]*models.MenuPriceModels, error) {
	menu, err := s.menuRepo.GetAllMenuPrices()
	if err != nil {
		return nil, err
	}
	return menu, nil
}

func (s MenuUseCaseImpl) AddNewMenu(products *models.MenuModels) (string, error) {
	err := validation.CheckEmpty(products.MenuID, products.MenuDesc, products.MenuStock, products.MenuPrice)
	if err != nil {
		return "", err
	}
	err = validation.CheckInt(products.MenuStock, products.MenuPrice)
	if err != nil {
		return "", err
	}
	if validation.IsStatusValid(products.MenuStatus) != true {
		return "", errors.New("Status not valid")
	}
	category, err := s.menuRepo.AddNewMenu(products)
	if err != nil {
		return "", err
	}
	return category, nil
}

func (s MenuUseCaseImpl) UpdateMenu(menu *models.MenuModels) (string, error) {
	product, err := s.menuRepo.UpdateMenu(menu)
	if err != nil {
		return "", err
	}
	return product, nil
}

func (s MenuUseCaseImpl) UpdateMenuPrice(day string, menu *models.MenuPriceModels) (string, error) {
	product, err := s.menuRepo.UpdateMenuPrice(day, menu)
	if err != nil {
		return "", err
	}
	return product, nil
}

func (s MenuUseCaseImpl) DeleteMenu(menuID string) (string, error) {
	product, err := s.menuRepo.DeleteMenu(menuID)
	if err != nil {
		return "", err
	}
	return product, nil
}

func InitMenuUseCase(menuRepo repositories.MenuRepositories) MenuUseCases {
	return &MenuUseCaseImpl{menuRepo}
}
