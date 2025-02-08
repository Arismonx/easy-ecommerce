package controllers

import (
	"github.com/Arismonx/easy-ecommerce/config"
	"github.com/Arismonx/easy-ecommerce/models"
	"github.com/gofiber/fiber/v2"
)

func CreateProduct(c *fiber.Ctx) error {
	newProduct := new(models.Products)

	if err := c.BodyParser(newProduct); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}
	config.DB.Create(newProduct)
	return c.JSON(newProduct)
}

func GetProducts(c *fiber.Ctx) error {
	var product []models.Products
	config.DB.Find(&product)
	return c.JSON(product)
}
