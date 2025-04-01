package routes

import (
	"github.com/gofiber/fiber/v2"
	"license/database"
	"license/models"
)

func AllKeys(c *fiber.Ctx) error {
	var keys []models.LicenseKey

	err := database.Client.Find(&keys).Error

	if err != nil {
		return c.SendString("No Keys Found")
	}

	return c.JSON(keys)
}
