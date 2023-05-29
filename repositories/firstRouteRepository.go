package repositories

import (
	"api_webservice_siasn/models/migrations"

	"gorm.io/gorm"
)

type FirstRouteRepository interface {
	FindAll() ([]migrations.Test, error)
}

type repository struct {
	database *gorm.DB
}

func NewFirstRouteRepository(database *gorm.DB) *repository {
	return &repository{database}
}

func (r *repository) FindAll() ([]migrations.Test, error) {
	var artikels []migrations.Test
	err := r.database.Debug().Find(&artikels).Error
	return artikels, err
}
