package repositories

import "github.com/inact25/E-WarungApi/masters/apis/models"

type MenuRepositories interface {
	GetAllMenu() ([]*models.MenuModels, error)
	GetAllMenuPrices() ([]*models.MenuPriceModels, error)
	UpdateMenu(products *models.MenuModels) (string, error)
	DeleteMenu(MenuID string) (string, error)
	AddNewMenu(day string, Menus *models.MenuModels) (string, error)
}
