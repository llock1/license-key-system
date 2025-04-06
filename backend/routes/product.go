package routes

import (
	"license/database"
	"license/dto"
	"license/helpers"
	"license/models"

	"github.com/gofiber/fiber/v3"
)

func GetProducts(c fiber.Ctx) error {

	var products []models.Product
	if err := database.Client.Find(&products).Error; err != nil {
		return err
	}

	productDTOs := []dto.ProductDTO{}
	for _, product := range products {
		productDTOs = append(productDTOs, dto.ProductDTO{
			ID:        product.ID,
			CreatedAt: product.CreatedAt,
			UpdatedAt: product.UpdatedAt,

			Name:        product.Name,
			OwnerID:     product.OwnerID,
			Title:       product.Title,
			Description: product.Description,
			Status:      product.Status,
			Version:     product.Version,
			IsListed:    product.IsListed,
		})
	}

	return c.JSON(fiber.Map{
		"success":  true,
		"products": productDTOs,
	})
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
