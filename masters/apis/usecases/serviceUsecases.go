package usecases

import (
	"github.com/inact25/E-WarungApi/masters/apis/models"
)

type ServiceUseCases interface {
	GetAllServices() ([]*models.ServicesModels, error)
	GetAllServicesByStatus(status string) ([]*models.ServicesModels, error)
	AddNewServices(day string, services *models.ServicesModels) (string, error)
}
