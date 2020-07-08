package usecases

import (
	"github.com/inact25/E-WarungApi/masters/apis/models"
	"github.com/inact25/E-WarungApi/masters/apis/repositories"
	"log"
)

type MenuUseCaseImpl struct {
	menuRepo repositories.MenuRepositories
}

func (s MenuUseCaseImpl) UpdateMenu(menu *models.MenuModels) (string, error) {
	product, err := s.menuRepo.UpdateMenu(menu)
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

func (s MenuUseCaseImpl) AddNewMenu(products *models.MenuModels) (string, error) {
	category, err := s.menuRepo.AddNewMenu(products)
	if err != nil {
		return "", err
	}
	return category, nil
}

func (s MenuUseCaseImpl) GetAllMenu() ([]*models.MenuModels, error) {
	menu, err := s.menuRepo.GetAllMenu()
	log.Println("U : ", menu)
	if err != nil {
		return nil, err
	}
	return menu, nil
}

func (s MenuUseCaseImpl) GetAllMenuPrices() ([]*models.MenuPriceModels, error) {
	menu, err := s.menuRepo.GetAllMenuPrices()
	log.Println("U : ", menu)
	if err != nil {
		return nil, err
	}
	return menu, nil
}

func (s MenuUseCaseImpl) AddNewMenuPrices(day string, products *models.MenuPriceModels) (string, error) {
	category, err := s.menuRepo.AddNewMenuPrices(day, products)
	if err != nil {
		return "", err
	}
	return category, nil
}

func InitMenuUseCase(menuRepo repositories.MenuRepositories) MenuUseCases {
	return &MenuUseCaseImpl{menuRepo}
}
