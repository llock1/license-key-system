package routes

import (
	"github.com/gofiber/fiber/v3"
)

func LogoutUser(c fiber.Ctx) error {

	type Response struct {
		Success bool   `json:"success"`
		Message string `json:"message"`
	}

	return c.JSON(Response{
		Success: false,
		Message: "Not implemented",
	})
}
