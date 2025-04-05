package routes

import (
	"license/database"
	"license/dto"
	"license/helpers"
	"license/models"

	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
)

func GetLicenses(c fiber.Ctx) error {
	var keys []models.License

	err := database.Client.Find(&keys).Error
	if err != nil {
		return err
	}

	return c.JSON(keys)
}

func DeleteLicense(c fiber.Ctx) error {
	id := c.Params("id")
	var key models.License

	err := helpers.DeleteModel(&key, id)
	if err != nil {
		return err
	}
	return c.SendStatus(fiber.StatusOK)
}

func CreateLicense(c fiber.Ctx) error {
	key := models.License{
		CreatorID: 1,
		ProductID: 1,
		Key:       uuid.New().String(),
	}
	if err := database.Client.Create(&key).Error; err != nil {
		return err
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "Key created",
		"license": dto.LicenseDTO{
			ID:   key.ID,
			Key:  key.Key,
			Hwid: key.HWID,
		},
	})
}
