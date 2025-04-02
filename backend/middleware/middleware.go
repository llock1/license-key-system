package middleware

import (
	"github.com/gofiber/fiber/v3"
	"github.com/golang-jwt/jwt/v5"
	"license/config"
)

func AuthMiddleware() fiber.Handler {
	return func(c fiber.Ctx) error {
		tokenString := c.Get("Authorization")
		if tokenString == "" {
			return c.SendStatus(fiber.StatusUnauthorized)
		}
		tokenString = tokenString[len("Bearer "):]
		jwtSecret := []byte(config.Vars.JWTSecret)
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return jwtSecret, nil
		})

		if err != nil {
			return c.SendStatus(fiber.StatusUnauthorized)
		}

		if !token.Valid {
			return c.SendStatus(fiber.StatusUnauthorized)
		}

		return c.Next()
	}
}
