package main

import (
	"github.com/RocketsLab/gofiber-and-gorm-api/application"
	"github.com/RocketsLab/gofiber-and-gorm-api/routes"
	"github.com/RocketsLab/gofiber-and-gorm-api/services"
	"log"
)

func main() {

	services.StartSession()

	services.InitDatabase()

	application.Start()

	routes.RegisterWebRoutes()

	log.Fatal(application.App.Listen(":3333"))

}
