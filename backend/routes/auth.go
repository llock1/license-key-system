package routes

import (
	"github.com/gofiber/fiber/v3"
)

type UserDTO struct {
	Username string  `json:"username"`
	Email    *string `json:"email"`
}

func AuthUser(c fiber.Ctx) error {

	type Response struct {
		Success bool     `json:"success"`
		Message string   `json:"message"`
		User    *UserDTO `json:"user"`
	}

	return c.JSON(Response{
		Success: false,
		Message: "Not implemented",
		User:    nil,
	})
}
