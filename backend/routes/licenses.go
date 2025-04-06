package routes

import (
	"license/database"
	"license/dto"
	"license/models"

	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
)

func GetLicenses(c fiber.Ctx) error {

	var licenses []models.License
	if err := database.Client.Find(&licenses).Error; err != nil {
		return err
	}

	licenseDTOs := []dto.LicenseDTO{}
	for _, license := range licenses {
		licenseDTOs = append(licenseDTOs, dto.LicenseDTO{
			ID:        license.ID,
			CreatedAt: license.CreatedAt,
			UpdatedAt: license.UpdatedAt,

			CreatorID: license.CreatorID,
			ProductID: license.ProductID,
			UserID:    license.UserID,

			Key:         license.Key,
			HWID:        license.HWID,
			HWIDResetAt: license.HWIDResetAt,
		})
	}

	return c.JSON(fiber.Map{
		"success":  true,
		"licenses": licenseDTOs,
	})
}

func DeleteLicense(c fiber.Ctx) error {
	id := c.Params("id")

	if err := database.Client.Delete(&models.License{}, id).Error; err != nil {
		return err
	}

	return c.SendStatus(fiber.StatusOK)
}

func CreateLicense(c fiber.Ctx) error {

	license := models.License{
		CreatorID: 1,
		ProductID: 1,
		Key:       uuid.New().String(),
	}

	if err := database.Client.Create(&license).Error; err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "Key created",
		"license": dto.LicenseDTO{
			ID:        license.ID,
			CreatedAt: license.CreatedAt,
			UpdatedAt: license.UpdatedAt,

			CreatorID: license.CreatorID,
			ProductID: license.ProductID,
			UserID:    license.UserID,

			Key:         license.Key,
			HWID:        license.HWID,
			HWIDResetAt: license.HWIDResetAt,
		},
	})
}
