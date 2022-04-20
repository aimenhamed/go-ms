package controllers

import (
	"github.com/aimenhamed/go-ms/services"
	"github.com/gin-gonic/gin"
)

type AccountController struct {
	accountService *services.AccountService
}

func NewAccountController(service *services.AccountService) AccountController {
	return AccountController{service}
}

func (c AccountController) RegisterRoutes(r *gin.RouterGroup) {
	r.GET("/accounts", c.accountService.GetAccount)
}
