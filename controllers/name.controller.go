package controllers

import (
	"github.com/aimenhamed/go-ms/services"
	"github.com/gin-gonic/gin"
)

type NamesController struct {
	nameService services.NamesService
}

func (c NamesController) RegisterRoutes(r *gin.RouterGroup) {
	c.nameService = services.NamesService{}
	r.GET("/names", c.nameService.GetName)
}
