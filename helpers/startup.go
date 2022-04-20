package helpers

import (
	"github.com/aimenhamed/go-ms/interfaces"
	"github.com/gin-gonic/gin"
)

func initialiseController(c interfaces.Controller, r *gin.RouterGroup) {
	c.RegisterRoutes(r)
}

func InitialiseControllers(c []interfaces.Controller, r *gin.RouterGroup) {
	for _, controller := range c {
		initialiseController(controller, r)
	}
}

func HealthCheck(r *gin.Engine) {
	r.GET("/health", func(c *gin.Context) {
		c.Status(200)
	})
}
