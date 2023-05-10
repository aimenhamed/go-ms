package helpers

import (
	"github.com/aimenhamed/go-ms/pkg/interfaces"
	"github.com/gofiber/fiber/v2"
)

func initialiseController(c interfaces.Controller, r fiber.Router) {
	c.RegisterRoutes(r)
}

func InitialiseControllers(c []interfaces.Controller, r fiber.Router) {
	for _, controller := range c {
		initialiseController(controller, r)
	}
}

func HealthCheck(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		c.Status(200)
		return c.JSON("healthy")
	})
	app.Get("/health", func(c *fiber.Ctx) error {
		c.Status(200)
		return c.JSON("healthy")
	})
}
