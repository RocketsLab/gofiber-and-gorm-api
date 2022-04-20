package application

import (
	"github.com/RocketsLab/gofiber-and-gorm-api/http/middleware"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

var (
	App *fiber.App
)

func Start(config ...fiber.Config) {

	App = fiber.New(config...)

	App.Use(middleware.AuthMiddleware(), cors.New(cors.Config{
		AllowCredentials: true,
		AllowOrigins:     "*",
		AllowHeaders:     "Access-Control-Allow-Origin, Content-Type, Origin, Accept",
	}))
}
