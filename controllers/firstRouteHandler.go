package controllers

import (
	"api_webservice_siasn/configs"
	"api_webservice_siasn/repositories"
	"api_webservice_siasn/services"
	"api_webservice_siasn/universalfunctions"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RootHandler(c *gin.Context) {
	firstRouteService := services.NewFirstRouteService(repositories.NewFirstRouteRepository(configs.Connect))
	firstRoute, err := firstRouteService.FindAll()

	universalfunctions.PanicErr(err)

	c.JSON(http.StatusOK, firstRoute)
}
