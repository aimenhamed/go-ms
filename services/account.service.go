package services

import (
	"github.com/gin-gonic/gin"
	"log"
)

type AccountService struct {
}

func NewAccountService() *AccountService {
	return &AccountService{}
}

func (s AccountService) GetAccount(c *gin.Context) {
	log.Printf("AccountService called.")
	c.JSON(200,
		gin.H{
			"message": "list Account here",
		})
}
