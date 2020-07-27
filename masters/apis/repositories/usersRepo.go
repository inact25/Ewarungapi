package repositories

import "github.com/inact25/E-WarungApi/masters/apis/models"

type UsersRepo interface {
	OAuth(*models.SelfUserModels) ([]*models.OAuth, error)
	GetSelfUser(*models.SelfUserModels) ([]*models.SelfUserModels, error)
	GetAllUsers() ([]*models.UserModels, error)
}
