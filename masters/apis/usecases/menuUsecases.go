package usecases

import "github.com/inact25/E-WarungApi/masters/apis/models"

type MenuUseCases interface {
	GetAllMenu() ([]*models.MenuModels, error)
	GetAllMenuByStatus(status string) ([]*models.MenuModels, error)
	GetAllMenuPrices() ([]*models.MenuPriceModels, error)
	AddNewMenu(day string, menus *models.MenuModels) (string, error)
	UpdateMenu(menus *models.MenuModels) (string, error)
	UpdateMenuPrice(day string, menus *models.MenuPriceModels) (string, error)
	DeleteMenu(menusID string) (string, error)
}
