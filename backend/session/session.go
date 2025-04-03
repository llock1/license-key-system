package session

import (
	"errors"

	"github.com/gofiber/fiber/v3"
	"github.com/golang-jwt/jwt/v5"
)

func IsUserAuthed(c fiber.Ctx) error {
	user, ok := c.Locals("user").(*jwt.Token)
	if !ok {
		return errors.New("missing or invalid token")
	}

	claims, ok := user.Claims.(jwt.MapClaims)
	if !ok {
		return errors.New("invalid token claims")
	}

	name, ok := claims["name"].(string)
	if !ok || name == "" {
		return errors.New("missing or invalid username in token")
	}

	return nil
}
