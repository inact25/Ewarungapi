package usecases

import "github.com/inact25/E-WarungApi/masters/apis/models"

type MenuUseCases interface {
	GetAllMenu() ([]*models.MenuModels, error)
	UpdateMenu(menus *models.MenuModels) (string, error)
	DeleteMenu(menusID string) (string, error)
	AddNewMenu(day string, Menus *models.MenuModels) (string, error)
	GetAllMenuPrices() ([]*models.MenuPriceModels, error)
}
