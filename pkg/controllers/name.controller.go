package controllers

import (
	"github.com/aimenhamed/go-ms/pkg/services"
	"github.com/gin-gonic/gin"
)

type NamesController struct {
	nameService *services.NamesService
}

func NewNameController(service *services.NamesService) NamesController {
	return NamesController{service}
}

func (c NamesController) RegisterRoutes(r *gin.RouterGroup) {
	r.GET("/names", c.nameService.GetName)
}
