package configs

import (
	"api_webservice_siasn/universals"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Connect *gorm.DB

func ConnectDatabase() {
	USERNAME := "root"
	PASSWORD := ""
	HOST := "127.0.0.1"
	PORT := "3306"
	DATABASE := "testtest"

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", USERNAME, PASSWORD, HOST, PORT, DATABASE)
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	universals.PanicErr(err)

	autoMigration(database)
	Connect = database
}
