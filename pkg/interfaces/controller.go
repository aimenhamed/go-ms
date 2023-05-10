package interfaces

import (
	"github.com/gofiber/fiber/v2"
)

type Controller interface {
	RegisterRoutes(r fiber.Router)
}
