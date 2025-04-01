package middleware

import (
	"github.com/gofiber/fiber/v2"
	"license/core"
)

func AuthMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		err := core.IsUserAuthed(c)
		if err != nil {
			return err
		}
		return c.Next()
	}
}
