package services

import (
	"github.com/gin-gonic/gin"
	"log"
)

type NamesService struct {
}

func NewNamesService() *NamesService {
	return &NamesService{}
}

func (s NamesService) GetName(c *gin.Context) {
	log.Printf("NamesService called.")
	c.JSON(200,
		gin.H{
			"message": "list Names here",
		})
}
