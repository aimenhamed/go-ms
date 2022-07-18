package controllers

import (
	"github.com/aimenhamed/go-ms/pkg/services"
	"github.com/gin-gonic/gin"
)

type AccountController struct {
	AccountService *services.AccountService
}

func (c AccountController) RegisterRoutes(r *gin.RouterGroup) {
	r.GET("/accounts", c.AccountService.GetAccount)
}
