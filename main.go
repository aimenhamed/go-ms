package main

import (
	"github.com/aimenhamed/go-ms/controllers"
	"github.com/aimenhamed/go-ms/interfaces"
	"github.com/gin-gonic/gin"
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
	initialiseController(controllers.NamesController{}, v1)

	r.Run()
}
