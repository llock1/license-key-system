package routes

import (
	"license/database"
	"license/helpers"
	"license/models"

	"github.com/gofiber/fiber/v3"
)

func GetProducts(c fiber.Ctx) error {
	var products []models.Product

	err := database.Client.Find(&products).Error

	if err != nil {
		return err
	}

	return c.JSON(products)
}

func DeleteProduct(c fiber.Ctx) error {
	id := c.Params("id")
	product := models.Product{}

	err := helpers.DeleteModel(&product, id)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "Successfully deleted product",
	})

}
