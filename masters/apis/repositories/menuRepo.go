package repositories

import "github.com/inact25/E-WarungApi/masters/apis/models"

type MenuRepositories interface {
	GetAllMenu() ([]*models.MenuModels, error)
	GetAllMenuByStatus(status string) ([]*models.MenuModels, error)
	GetAllMenuPrices() ([]*models.MenuPriceModels, error)
	AddNewMenu(day string, Menus *models.MenuModels) (string, error)
	UpdateMenu(products *models.MenuModels) (string, error)
	UpdateMenuPrice(day string, menus *models.MenuPriceModels) (string, error)
	DeleteMenu(MenuID string) (string, error)
}
