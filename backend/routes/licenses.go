package routes

import (
	"license/database"
	"license/models"

	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
)

type LicenseDTO struct {
	ID   uint   `json:"id"`
	Key  string `json:"key"`
	Hwid string `json:"hwid"`
}

func GetLicenses(c fiber.Ctx) error {
	var keys []models.License

	err := database.Client.Find(&keys).Error

	if err != nil {
		return c.SendString("No licenses found")
	}

	return c.JSON(keys)
}

func DeleteLicense(c fiber.Ctx) error {
	id := c.Params("id")
	var key models.License
	database.Client.First(&key, id)
	database.Client.Delete(&key)
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
		"license": LicenseDTO{
			ID:   key.ID,
			Key:  key.Key,
			Hwid: key.HWID,
		},
	})
}
