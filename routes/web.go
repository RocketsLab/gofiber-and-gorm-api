package routes

import (
	"github.com/RocketsLab/gofiber-and-gorm-api/application"
	"github.com/RocketsLab/gofiber-and-gorm-api/http/controllers"
)

func RegisterWebRoutes() {

	router := application.App

	//USERS
	userController := controllers.UserController{}
	router.Get("/users", userController.Index)
	router.Post("/users", userController.Store)
	router.Get("/users/:user", userController.Show)

	//AUTH
	authController := controllers.AuthController{}
	router.Post("/auth/login", authController.Login).Name("auth.login")
	router.Post("/auth/logout", authController.Logout).Name("auth.logout")
	router.Post("/auth/register", authController.Register).Name("auth.register")
}
