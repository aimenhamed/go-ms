package controllers

import (
	"github.com/aimenhamed/go-ms/pkg/services"
	"github.com/gin-gonic/gin"
)

type NamesController struct {
	NameService *services.NamesService
}

func (c NamesController) RegisterRoutes(r *gin.RouterGroup) {
	r.GET("/names", c.NameService.GetName)
}
