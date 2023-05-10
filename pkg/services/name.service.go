package services

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

type NamesService struct {
}

func NewNamesService() *NamesService {
	return &NamesService{}
}

func (s NamesService) GetName(c *fiber.Ctx) error {
	log.Printf("NamesService called.")
	return c.JSON(
		fiber.Map{
			"message": "list Names here",
		})
}
