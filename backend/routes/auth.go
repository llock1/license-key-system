package routes

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"license/config"
	"time"
)

type UserDTO struct {
	Username string  `json:"username"`
	Email    *string `json:"email"`
}

func AuthUser(c *fiber.Ctx) error {
	// https://docs.gofiber.io/contrib/jwt/
	type User struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	var u User

	err := c.BodyParser(&u)

	// Throws Unauthorized error
	if u.Username != "john" || u.Password != "doe" {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	// Create the Claims
	claims := jwt.MapClaims{
		"name":  "John Doe",
		"admin": true,
		"exp":   time.Now().Add(time.Hour * 72).Unix(),
	}

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	jwtSecret := []byte(config.Vars.JWTSecret)
	t, err := token.SignedString(jwtSecret)
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.JSON(fiber.Map{"token": t})
}

func RestrictedExample(c *fiber.Ctx) error {
	return c.SendString("ok")
}

func VerifyJWTToken(tokenString string) error {
	jwtSecret := []byte(config.Vars.JWTSecret)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if err != nil {
		return err
	}

	if !token.Valid {
		return fmt.Errorf("invalid token")
	}

	return nil
}

func CheckTokenHandler(c *fiber.Ctx) error {

	type Token struct {
		Token string `json:"token"`
	}

	var t Token

	c.BodyParser(&t)

	err := VerifyJWTToken(t.Token)

	if err != nil {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	return c.SendStatus(fiber.StatusOK)
}
