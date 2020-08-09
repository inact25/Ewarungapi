package repositories

import "github.com/inact25/E-WarungApi/masters/apis/models"

type UsersRepo interface {
	OAuth(*models.UserModels) ([]*models.OAuth, error)
	GetUser(*models.UserModels) ([]*models.UserModels, error)
	GetAllUsers() ([]*models.UserModels, error)
	UpdateUser(*models.UserModels) (string, error)
}
