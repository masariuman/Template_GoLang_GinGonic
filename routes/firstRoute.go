package routes

import (
	"api_webservice_siasn/configs"
	"api_webservice_siasn/controllers"
	"api_webservice_siasn/repositories"
	"api_webservice_siasn/services"

	"github.com/gin-gonic/gin"
)

func firstRoute(superRoute *gin.RouterGroup) {
	firsRouteRepository := repositories.NewFirstRouteRepository(configs.Connect)
	firstRouteService := services.NewFirstRouteService(firsRouteRepository)
	firstRouteHandler := controllers.NewFristRouteHandler(firstRouteService)
	firstRoute := superRoute.Group("/firstRoute")
	{
		firstRoute.GET("", firstRouteHandler.RootHandler)
	}
}
