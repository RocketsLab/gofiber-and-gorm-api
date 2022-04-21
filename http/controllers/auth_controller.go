package controllers

import (
	"github.com/RocketsLab/gofiber-and-gorm-api/http/repositories"
	"github.com/RocketsLab/gofiber-and-gorm-api/http/requests"
	"github.com/RocketsLab/gofiber-and-gorm-api/models"
	"github.com/RocketsLab/gofiber-and-gorm-api/services"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

type AuthController struct{}

func (c *AuthController) Login(ctx *fiber.Ctx) error {
	var data requests.LoginRequest
	var user models.User

	_ = ctx.BodyParser(&data)
	result := services.DbConnection.Where("email = ?", data.Email).First(&user)
	if err := result.Error; err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "password or user not match",
			"error":   err.Error(),
		})
	}
	// PROCEED TO LOGIN
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(data.Password)); err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "password or user not match",
			"error":   err.Error(),
		})
	}

	session, err := services.GetSessionStore(ctx)
	if err != nil {
		return services.SessionError(ctx, err)
	}

	session.Set(services.AuthKey, true)
	session.Set(services.UserId, user.ID)

	err = session.Save()
	if err != nil {
		return services.SessionError(ctx, err)
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"user":    &user,
		"message": "logged in",
	})
}

func (c *AuthController) Logout(ctx *fiber.Ctx) error {

	session, err := services.GetSessionStore(ctx)
	if err != nil {
		return services.SessionNotFound(ctx)
	}

	err = session.Destroy()
	if err != nil {
		return services.SessionError(ctx, err)
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "logged out",
	})
}

func (c *AuthController) Register(ctx *fiber.Ctx) error {

	var data requests.UserRequest
	_ = ctx.BodyParser(&data)

	user := models.User{
		Name:     data.Name,
		Email:    data.Email,
		Password: data.Password,
	}

	err := repositories.UserSave(&user)
	if err != nil {
		return ctx.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"message": "failed to create user",
			"error":   err.Error(),
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"user":    &user,
		"message": "OK",
	})
}
