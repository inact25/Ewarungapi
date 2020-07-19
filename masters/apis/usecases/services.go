package usecases

import (
	"errors"
	"github.com/inact25/E-WarungApi/masters/apis/models"
	"github.com/inact25/E-WarungApi/masters/apis/repositories"
	"github.com/inact25/E-WarungApi/utils/validation"
)

type ServiceUseCaseImpl struct {
	serviceRepo repositories.ServiceRepositories
}

func (s ServiceUseCaseImpl) GetAllServicesByStatus(status string) ([]*models.ServicesModels, error) {
	if validation.IsStatusValid(status) != true {
		return nil, errors.New("ERROR")
	}
	menu, err := s.serviceRepo.GetAllServicesByStatus(status)
	if err != nil {
		return nil, err
	}
	return menu, nil
}

func (s ServiceUseCaseImpl) GetAllServices() ([]*models.ServicesModels, error) {
	menu, err := s.serviceRepo.GetAllServices()
	if err != nil {
		return nil, err
	}
	return menu, nil
}

func (s ServiceUseCaseImpl) AddNewServices(services *models.ServicesModels) (string, error) {
	err := validation.CheckEmpty(services.ServicesID, services.ServicesDesc, services.ServicesStatus, services.ServicePrice)
	if err != nil {
		return "", err
	}
	err = validation.CheckInt(services.ServicePrice)
	if err != nil {
		return "", err
	}
	if validation.IsStatusValid(services.ServicesStatus) != true {
		return "", errors.New("Status not valid")
	}
	service, err := s.serviceRepo.AddNewServices(services)
	if err != nil {
		return "", err
	}
	return service, nil
}

func (s ServiceUseCaseImpl) UpdateServices(services *models.ServicesModels) (string, error) {
	err := validation.CheckEmpty(services.ServicesID, services.ServicesDesc, services.ServicesStatus)
	if err != nil {
		return "", err
	}
	if validation.IsStatusValid(services.ServicesStatus) != true {
		return "", errors.New("Status not valid")
	}
	service, err := s.serviceRepo.UpdateServices(services)
	if err != nil {
		return "", err
	}
	return service, nil
}

func (s ServiceUseCaseImpl) UpdateServicesPrice(day string, services *models.ServicesModels) (string, error) {
	err := validation.CheckEmpty(services.ServicesID, services.ServicePrice)
	if err != nil {
		return "", err
	}
	err = validation.CheckInt(services.ServicePrice)
	if err != nil {
		return "", err
	}
	product, err := s.serviceRepo.UpdateServicesPrice(day, services)
	if err != nil {
		return "", err
	}
	return product, nil
}

func (s ServiceUseCaseImpl) DeleteServices(servicesID string) (string, error) {
	err := validation.CheckEmpty(servicesID)
	if err != nil {
		return "", err
	}
	product, err := s.serviceRepo.DeleteServices(servicesID)
	if err != nil {
		return "", err
	}
	return product, nil
}

func InitServiceUseCase(serviceRepo repositories.ServiceRepositories) ServiceUseCases {
	return &ServiceUseCaseImpl{serviceRepo}
}
