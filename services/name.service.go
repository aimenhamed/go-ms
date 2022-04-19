package services

import (
	"github.com/gin-gonic/gin"
	"log"
)

type NamesService struct {
}

func (s NamesService) GetName(c *gin.Context) {
	log.Printf("NamesService called.")
	c.String(200, "list here")
}
