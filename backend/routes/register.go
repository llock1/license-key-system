package routes

import (
	"license/database"
	"license/models"
	"strings"

	"github.com/gofiber/fiber/v3"
	"golang.org/x/crypto/bcrypt"
)

func RegisterUser(c fiber.Ctx) error {

	payload := struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}{}

	if err := c.Bind().JSON(&payload); err != nil {
		return err
	}

	tx := database.Client.Begin()

	var count int64
	if err := tx.Model(&models.User{}).Where("email = ? OR username = ?", strings.ToLower(payload.Email), payload.Username).
		Count(&count).Error; err != nil {
		return err
	}

	if count > 0 {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{
			"success": false,
			"message": "user already exists",
			"user":    nil,
		})
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(payload.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user := models.User{
		Username: payload.Username,
		Email:    payload.Email,
		Password: string(hashedPassword),

		IsSuperAdmin: false,
		IsAdmin:      false,
		IsModerator:  false,
		IsSupport:    false,
		IsStaff:      false,
		IsBanned:     false,
	}

	if err := tx.Create(&user).Error; err != nil {
		return err
	}

	if err := tx.Commit().Error; err != nil {
		return err
	}

	return c.JSON(fiber.Map{
		"success": true,
		"message": nil,
		"user": UserDTO{
			Email:    &user.Email,
			Username: user.Username,

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
