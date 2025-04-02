package routes

import (
	"fmt"
	"github.com/gofiber/fiber/v3"
	"github.com/golang-jwt/jwt/v5"
	"license/config"
	"license/core"
	"license/database"
	"license/models"
	"time"
)

type UserDTO struct {
	Username string  `json:"username"`
	Email    *string `json:"email"`
}

func AuthUser(c fiber.Ctx) error {

	var userRequest models.User
	var user models.User

	//payload := struct {
	//	User     string `json:"user"`
	//	Password string `json:"password"`
	//}{}

	if err := c.Bind().JSON(&userRequest); err != nil {
		return err
	}

	// Throws Unauthorized error
	if err := database.Client.First(&user, "username = ?", "admin").Error; err != nil {
		return c.SendStatus(fiber.StatusNotFound)
	}

	if user.Username == "" {
		return c.SendStatus(fiber.StatusNotFound)
	}

	if !core.CheckPasswordHash(userRequest.Password, user.Password) {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	// Create the Claims
	claims := jwt.MapClaims{
		"name": userRequest.Username,
		"exp":  time.Now().Add(time.Hour * 72).Unix(),
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

func RestrictedExample(c fiber.Ctx) error {
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

func CheckTokenHandler(c fiber.Ctx) error {

	payload := struct {
		Token string `json:"token"`
	}{}

	if err := c.Bind().JSON(&payload); err != nil {
		return err
	}

	err := VerifyJWTToken(payload.Token)

	if err != nil {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	return c.SendStatus(fiber.StatusOK)
}
