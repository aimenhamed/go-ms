package controllers

import (
	"github.com/aimenhamed/go-ms/pkg/services"
	"github.com/gofiber/fiber/v2"
)

type AccountController struct {
	AccountService *services.AccountService
}

func (c AccountController) RegisterRoutes(r fiber.Router) {
	r.Get("/accounts", c.AccountService.GetAccount)
}
