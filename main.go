package main

import (
	"github.com/RocketsLab/gofiber-and-gorm-api/application"
	"github.com/RocketsLab/gofiber-and-gorm-api/http/service"
	"github.com/RocketsLab/gofiber-and-gorm-api/routes"
	"log"
)

func main() {

	service.StartSession()

	service.InitDatabase()

	application.Start()

	routes.RegisterWebRoutes()

	log.Fatal(application.App.Listen(":3333"))

}
