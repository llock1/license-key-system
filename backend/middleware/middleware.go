package middleware

import (
	"license/core"

	"github.com/gofiber/fiber/v3"
)

func AuthMiddleware() fiber.Handler {
	return func(c fiber.Ctx) error {
		err := core.IsUserAuthed(c)
		if err != nil {
			return err
		}
		return c.Next()
	}
}
