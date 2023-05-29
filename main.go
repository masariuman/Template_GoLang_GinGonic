package main

import (
	"api_webservice_siasn/configs"
	"api_webservice_siasn/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	configs.ConnectDatabase()

	router := gin.Default()
	v1 := router.Group("/v1")
	routes.AddRoutes(v1)

	router.Run(":8877")
}
