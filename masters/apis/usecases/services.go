package usecases

import (
	"github.com/inact25/E-WarungApi/masters/apis/models"
	"github.com/inact25/E-WarungApi/masters/apis/repositories"
)

type ServiceUseCaseImpl struct {
	serviceRepo repositories.ServicesRepoImpl
}

func (s ServiceUseCaseImpl) GetAllServices() ([]*models.ServicesModels, error) {
	services, err := s.serviceRepo.GetAllServices()
	if err != nil {
		return nil, err
	}
	return services, nil
}

func (s ServiceUseCaseImpl) AddNewServices(services *models.ServicesModels) (string, error) {
	service, err := s.serviceRepo.AddNewServices(services)
	if err != nil {
		return "", err
	}
	return service, nil
}

func (s ServiceUseCaseImpl) AddNewServicesPrice(services *models.ServicesModels) (string, error) {
	service, err := s.serviceRepo.AddNewServicePrice(services)
	if err != nil {
		return "", err
	}
	return service, nil

}

func InitServiceUseCase(serviceRepo repositories.ServicesRepoImpl) ServiceUseCases {
	return &ServiceUseCaseImpl{serviceRepo}
}
