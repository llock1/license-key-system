package routes

import (
	"errors"
	"license/config"
	"license/core"
	"license/database"
	"license/models"
	"strings"
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

func LoginUser(c fiber.Ctx) error {
	payload := struct {
		User     string `json:"user"`
		Password string `json:"password"`
	}{}

	if err := c.Bind().JSON(&payload); err != nil {
		return err
	}

	var user models.User
	if err := database.Client.Where("username = ? OR email = ?", payload.User, strings.ToLower(payload.User)).
		Find(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"success": false,
				"message": "user doesn't exist",
				"user":    nil,
			})
		}
		return err
	}

	if !core.CheckPasswordHash(payload.Password, user.Password) {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"success": false,
			"message": "invalid credentials",
			"user":    nil,
		})
	}

	claims := jwt.MapClaims{
		"name": user.Username,
		"exp":  time.Now().Add(time.Hour * 72).Unix(),
	}

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	jwtSecret := []byte(config.Vars.JWTSecret)
	tokenStr, err := token.SignedString(jwtSecret)
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": nil,
		"token":   tokenStr,
		"user": UserDTO{
			Username: user.Username,
			Email:    &user.Email,

			EmailVerifiedAt: user.EmailVerifiedAt,
			PasswordResetAt: user.PasswordResetAt,

			IsSuperAdmin: user.IsSuperAdmin,
			IsAdmin:      user.IsAdmin,
			IsModerator:  user.IsModerator,
			IsSupport:    user.IsSupport,
			IsStaff:      user.IsStaff,
			IsBanned:     user.IsBanned,
		},
	})
}
