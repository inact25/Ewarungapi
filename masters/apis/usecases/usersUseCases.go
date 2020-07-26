package usecases

import "github.com/inact25/E-WarungApi/masters/apis/models"

type UsersUseCases interface {
	OAuth(*models.SelfUserModels) ([]*models.OAuth, error)
}
