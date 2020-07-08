package usecases

import (
	"github.com/inact25/E-WarungApi/masters/apis/models"
	"github.com/inact25/E-WarungApi/masters/apis/repositories"
	"log"
)

type ServiceUseCaseImpl struct {
	serviceRepo repositories.ServiceRepositories
}

func (s ServiceUseCaseImpl) GetAllServicesByStatus(status string) ([]*models.ServicesModels, error) {
	menu, err := s.serviceRepo.GetAllServicesByStatus(status)
	log.Println("U : ", status)
	if err != nil {
		return nil, err
	}
	return menu, nil
}

func (s ServiceUseCaseImpl) GetAllServices() ([]*models.ServicesModels, error) {
	menu, err := s.serviceRepo.GetAllServices()
	log.Println("U : ", menu)
	if err != nil {
		return nil, err
	}
	return menu, nil
}

func (s ServiceUseCaseImpl) AddNewServices(day string, services *models.ServicesModels) (string, error) {
	category, err := s.serviceRepo.AddNewServices(day, services)
	if err != nil {
		return "", err
	}
	return category, nil
}

func InitServiceUseCase(serviceRepo repositories.ServiceRepositories) ServiceUseCases {
	return &ServiceUseCaseImpl{serviceRepo}
}
