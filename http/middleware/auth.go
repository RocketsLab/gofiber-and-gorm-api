package middleware

import (
	"github.com/RocketsLab/gofiber-and-gorm-api/services"
	"github.com/gofiber/fiber/v2"
)

func AuthMiddleware() fiber.Handler {
	return sessionAuthMiddleware
}

func sessionAuthMiddleware(ctx *fiber.Ctx) error {

	if ctx.Path() == "/auth/login" {
		return ctx.Next()
	}
	if ctx.Path() == "/auth/register" {
		return ctx.Next()
	}

	session, err := services.SessionStore.Get(ctx)
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "not authenticated",
		})
	}

	if session.Get(services.AuthKey) == nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "not authenticated",
		})
	}

	return ctx.Next()
}
