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

func Login(c fiber.Ctx) error {
	payload := struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}{}

	if err := c.Bind().JSON(&payload); err != nil {
		return err
	}

	fmt.Println(payload.Username)
	fmt.Println(payload.Password)

	var user models.User
	var count int64

	if err := database.Client.Model(&models.User{}).Where("username = ?", payload.Username).Count(&count).Error; err != nil {
		return err
	}

	if count <= 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"message": "username or password is wrong",
			"user":    nil,
		})
	}

	if err := database.Client.Where("username = ?", payload.Username).First(&user).Error; err != nil {
		return err
	}

	if !core.CheckPasswordHash(payload.Password, user.Password) {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"success": false,
			"message": "username or password is wrong",
			"user":    nil,
		})
	}

	claims := jwt.MapClaims{
		"name": payload.Username,
		"exp":  time.Now().Add(time.Hour * 72).Unix(),
	}

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	jwtSecret := []byte(config.Vars.JWTSecret)
	t, err := token.SignedString(jwtSecret)
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "Successfully Authenticated",
		"token":   t,
		"user": UserDTO{
			Username: payload.Username,
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
