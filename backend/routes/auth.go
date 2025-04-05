package routes

import (
	"license/auth"

	"github.com/gofiber/fiber/v3"
)

func CheckTokenHandler(c fiber.Ctx) error {

	payload := struct {
		Token string `json:"token"`
	}{}

	if err := c.Bind().JSON(&payload); err != nil {
		return err
	}

	err := auth.VerifyJWTToken(payload.Token)

	if err != nil {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	return c.SendStatus(fiber.StatusOK)
}
