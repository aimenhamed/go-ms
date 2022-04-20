package main

import (
	"github.com/aimenhamed/go-ms/controllers"
	"github.com/aimenhamed/go-ms/helpers"
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

func main() {
	r := gin.Default()
	helpers.HealthCheck(r)

	v1 := r.Group("/api/v1")
	c := []interfaces.Controller{nameController, accountController}
	helpers.InitialiseControllers(c, v1)

	r.Run()
}
