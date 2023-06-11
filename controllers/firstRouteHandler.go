package controllers

import (
	"api_webservice_siasn/services"
	"api_webservice_siasn/universalfunctions"
	"net/http"

	"github.com/gin-gonic/gin"
)

type firstRouteHandler struct {
	firstRouteService services.FirstRouteService
}

func NewFristRouteHandler(firstRouteService services.FirstRouteService) *firstRouteHandler {
	return &firstRouteHandler{firstRouteService}
}

func (h *firstRouteHandler) RootHandler(c *gin.Context) {
	firstRoute, err := h.firstRouteService.FindAll()

	universalfunctions.PanicErr(err)

	c.JSON(http.StatusOK, firstRoute)
}
