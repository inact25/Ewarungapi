package repositories

import "github.com/inact25/E-WarungApi/masters/apis/models"

type MenuRepositories interface {
	GetAllMenu() ([]*models.MenuModels, error)
	UpdateMenu(products *models.MenuModels) (string, error)
	DeleteMenu(MenuID string) (string, error)
	AddNewMenu(Menus *models.MenuModels) (string, error)
}
