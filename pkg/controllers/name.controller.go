package controllers

import (
	"github.com/aimenhamed/go-ms/pkg/services"
	"github.com/gofiber/fiber/v2"
)

type NamesController struct {
	NameService *services.NamesService
}

func (c NamesController) RegisterRoutes(r fiber.Router) {
	r.Get("/names", c.NameService.GetName)
}
