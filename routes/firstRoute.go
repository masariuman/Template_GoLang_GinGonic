package routes

import (
	"api_webservice_siasn/controllers"

	"github.com/gin-gonic/gin"
)

func firstRoute(superRoute *gin.RouterGroup) {
	firstRoute := superRoute.Group("/firstRoute")
	{
		firstRoute.GET("", controllers.RootHandler)
	}
}
