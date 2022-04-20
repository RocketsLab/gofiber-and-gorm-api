package service

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/storage/redis"
	"time"
)

var (
	SessionStore *session.Store
	AuthKey      string = "authenticated"
	UserId       string = "user_id"
)

func StartSession() {

	storage := redis.New()

	SessionStore = session.New(session.Config{
		CookieHTTPOnly: true,
		Expiration:     time.Hour * 1,
		Storage:        storage,
	})
}

func GetSessionStore(ctx *fiber.Ctx) (*session.Session, error) {
	sess, err := SessionStore.Get(ctx)
	return sess, err
}

func SessionError(ctx *fiber.Ctx, err error) error {
	return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
		"message": "something wrong: " + err.Error(),
	})
}

func SessionNotFound(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "session not found",
	})
}
