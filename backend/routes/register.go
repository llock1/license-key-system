package routes

import (
	"fmt"
	"license/database"
	"license/models"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
)

func RegisterUser(c fiber.Ctx) error {

	payload := struct {
		Username string `json:"username" validate:"required,min=3,max=24,alphanum"`
		Email    string `json:"email" validate:"required,email,max=320"`
		Password string `json:"password" validate:"required,min=8,max=64"`
	}{}

	if err := c.Bind().JSON(&payload); err != nil {
		return err
	}

	validate := validator.New()

	if err := validate.Struct(payload); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": fmt.Sprintf("Validation error: %s", validationErrors),
			"user":    nil,
		})
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

	user := models.User{
		Username: payload.Username,
		Email:    payload.Email,

		IsSuperAdmin: false,
		IsAdmin:      false,
		IsModerator:  false,
		IsSupport:    false,
		IsStaff:      false,
		IsBanned:     false,
	}
	user.SetPassword(payload.Password)

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
