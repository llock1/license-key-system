package middleware

import (
	"license/config"

	"github.com/gofiber/fiber/v3"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware() fiber.Handler {
	return func(c fiber.Ctx) error {
		tokenString := c.Get("Authorization")
		if tokenString == "" {
			return c.Status(fiber.StatusUnauthorized).
				SendString("no authorization token")
		}
		tokenString = tokenString[len("Bearer "):]
		jwtSecret := []byte(config.Vars.JWTSecret)
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return jwtSecret, nil
		})

		if err != nil {
			return c.Status(fiber.StatusInternalServerError).
				SendString("failed to parse token")
		}

		if !token.Valid {
			return c.Status(fiber.StatusUnauthorized).
				SendString("invalid token")
		}

		return c.Next()
	}
}
