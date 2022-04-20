package controllers

import (
	"fmt"
	"github.com/RocketsLab/gofiber-and-gorm-api/http/repositories"
	"github.com/RocketsLab/gofiber-and-gorm-api/http/requests"
	"github.com/RocketsLab/gofiber-and-gorm-api/models"
	"github.com/gofiber/fiber/v2"
)

type UserController struct{}

func (c *UserController) Store(ctx *fiber.Ctx) error {

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

func (c *UserController) Index(ctx *fiber.Ctx) error {

	users, _ := repositories.UserAll()

	return ctx.JSON(fiber.Map{
		"users":   users,
		"count":   len(users),
		"message": "OK",
	})
}

func (c *UserController) Show(ctx *fiber.Ctx) error {
	user, err := repositories.UserFindByID(ctx.Params("user"))
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "user not found",
		})
	}
	fmt.Println(user, err)
	return ctx.JSON(fiber.Map{
		"user":    user,
		"message": "OK",
	})
}
