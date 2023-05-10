package services

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

type AccountService struct {
}

func NewAccountService() *AccountService {
	return &AccountService{}
}

func (s AccountService) GetAccount(c *fiber.Ctx) error {
	log.Printf("AccountService called.")
	return c.JSON(fiber.Map{
		"message": "list Account here",
	})
}
