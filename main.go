package main

import (
	"github.com/aimenhamed/go-ms/controllers"
	"github.com/aimenhamed/go-ms/interfaces"
	"github.com/aimenhamed/go-ms/services"
	"github.com/gin-gonic/gin"
)

var (
	nameService       = services.NewNamesService()
	nameController    = controllers.NewNameController(nameService)
	accountService    = services.NewAccountService()
	accountController = controllers.NewAccountController(accountService)
)

func initialiseController(c interfaces.Controller, r *gin.RouterGroup) {
	c.RegisterRoutes(r)
}

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	v1 := r.Group("/api/v1")
	initialiseController(nameController, v1)
	initialiseController(accountController, v1)

	r.Run()
}
