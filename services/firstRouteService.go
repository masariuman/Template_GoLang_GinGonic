package services

import (
	"api_webservice_siasn/models/migrations"
	"api_webservice_siasn/repositories"
)

type FirstRouteService interface {
	FindAll() ([]migrations.Test, error)
}

type service struct {
	repository repositories.FirstRouteRepository
}

func NewFirstRouteService(repository repositories.FirstRouteRepository) *service {
	return &service{repository}
}

func (s *service) FindAll() ([]migrations.Test, error) {
	return s.repository.FindAll()
}
