package routes

import (
	"fmt"
	"license/config"
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/golang-jwt/jwt/v5"
)

type UserDTO struct {
	Username string  `json:"username"`
	Email    *string `json:"email"`

	EmailVerifiedAt *time.Time `json:"email_verified_at"`
	PasswordResetAt *time.Time `json:"password_reset_at"`

	IsSuperAdmin bool `json:"is_super_admin"`
	IsAdmin      bool `json:"is_admin"`
	IsModerator  bool `json:"is_moderator"`
	IsSupport    bool `json:"is_support"`
	IsStaff      bool `json:"is_staff"`
	IsBanned     bool `json:"is_banned"`
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
