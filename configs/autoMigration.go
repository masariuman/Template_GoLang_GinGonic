package configs

import (
	"api_webservice_siasn/models/migrations"

	"gorm.io/gorm"
)

func autoMigration(connect *gorm.DB) {
	connect.AutoMigrate(migrations.Test{})
}
