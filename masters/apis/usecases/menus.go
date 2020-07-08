package usecases

import (
	"github.com/inact25/E-WarungApi/masters/apis/models"
	"github.com/inact25/E-WarungApi/masters/apis/repositories"
	"log"
)

type MenuUseCaseImpl struct {
	menuRepo repositories.MenuRepositories
}

func (s MenuUseCaseImpl) GetAllMenu() ([]*models.MenuModels, error) {
	menu, err := s.menuRepo.GetAllMenu()
	log.Println("U : ", menu)
	if err != nil {
		return nil, err
	}
	return menu, nil
}

func (s MenuUseCaseImpl) GetAllMenuByStatus(status string) ([]*models.MenuModels, error) {
	menu, err := s.menuRepo.GetAllMenuByStatus(status)
	log.Println("U : ", status)
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

func (s MenuUseCaseImpl) AddNewMenu(day string, products *models.MenuModels) (string, error) {
	category, err := s.menuRepo.AddNewMenu(day, products)
	if err != nil {
		return "", err
	}
	return category, nil
}

func (s MenuUseCaseImpl) UpdateMenu(menu *models.MenuModels) (string, error) {
	log.Println("U :", menu)
	product, err := s.menuRepo.UpdateMenu(menu)
	if err != nil {
		return "", err
	}
	return product, nil
}

func (s MenuUseCaseImpl) UpdateMenuPrice(day string, menu *models.MenuPriceModels) (string, error) {
	log.Println("U :", menu)
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
