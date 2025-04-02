package routes

import (
	"license/database"
	"license/models"

	"github.com/gofiber/fiber/v3"
)

func AllKeys(c fiber.Ctx) error {
	var keys []models.LicenseKey

	err := database.Client.Find(&keys).Error

	if err != nil {
		return c.SendString("No Keys Found")
	}

	return c.JSON(keys)
}

func DeleteKey(c fiber.Ctx) error {
	id := c.Params("id")
	var key models.LicenseKey
	database.Client.First(&key, id)
	database.Client.Delete(&key)
	return c.SendStatus(fiber.StatusOK)
}
