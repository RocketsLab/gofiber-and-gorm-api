package routes

import (
	"github.com/RocketsLab/gofiber-and-gorm-api/application"
	"github.com/RocketsLab/gofiber-and-gorm-api/http/controllers"
)

func RegisterWebRoutes() {
	//USERS
	userController := controllers.UserController{}
	application.App.Get("/users", userController.Index)
	application.App.Post("/users", userController.Store)
	application.App.Get("/users/:user", userController.Show)

	//AUTH
	authController := controllers.AuthController{}
	application.App.Post("/auth/login", authController.Login).Name("auth.login")
	application.App.Post("/auth/logout", authController.Logout).Name("auth.logout")
	application.App.Post("/auth/register", authController.Register).Name("auth.register")
}
