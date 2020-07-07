package usecases

import "github.com/inact25/E-WarungApi/masters/apis/models"

type ServiceUseCases interface {
	GetAllServices() ([]*models.ServicesModels, error)
	AddNewServices(services *models.ServicesModels) (string, error)
	AddNewServicesPrice(services *models.ServicesModels) (string, error)
}
