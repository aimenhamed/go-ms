package main

import (
	"log"

	"github.com/aimenhamed/go-ms/pkg/controllers"
	"github.com/aimenhamed/go-ms/pkg/helpers"
	"github.com/aimenhamed/go-ms/pkg/interfaces"
	"github.com/aimenhamed/go-ms/pkg/services"
	"github.com/gofiber/fiber/v2"
)

var (
	nameService       = services.NamesService{}
	nameController    = controllers.NamesController{NameService: &nameService}
	accountService    = services.AccountService{}
	accountController = controllers.AccountController{AccountService: &accountService}
)

func main() {
	app := fiber.New()
	helpers.HealthCheck(app)

	v1 := app.Group("/api/v1")
	c := []interfaces.Controller{&nameController, &accountController}
	helpers.InitialiseControllers(c, v1)

	log.Fatal(app.Listen(":8080"))
}
