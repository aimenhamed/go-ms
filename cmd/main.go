package main

import (
	"github.com/aimenhamed/go-ms/pkg/controllers"
	"github.com/aimenhamed/go-ms/pkg/helpers"
	"github.com/aimenhamed/go-ms/pkg/interfaces"
	"github.com/aimenhamed/go-ms/pkg/services"
	"github.com/gin-gonic/gin"
)

var (
	nameService       = services.NamesService{}
  nameController    = controllers.NamesController{NameService: &nameService}
	accountService    = services.AccountService{}
  accountController = controllers.AccountController{AccountService: &accountService}
)

func main() {
	r := gin.Default()
	helpers.HealthCheck(r)

	v1 := r.Group("/api/v1")
	c := []interfaces.Controller{&nameController, &accountController}
	helpers.InitialiseControllers(c, v1)

	r.Run()
}
