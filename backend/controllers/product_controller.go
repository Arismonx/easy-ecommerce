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

	if err := config.DB.Create(newProduct).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Cannot Create Product",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(newProduct)
}

func GetProducts(c *fiber.Ctx) error {
	var product []models.Products
	if err := config.DB.Find(&product).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Cannot get Products",
		})
	}
	return c.JSON(product)
}

func GetProductByID(c *fiber.Ctx) error {
	product := new(models.Products)
	id := c.Params("id")
	if err := config.DB.First(&product, id).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Product not found",
		})
	}
	return c.JSON(product)
}

func UpdateProductByID(c *fiber.Ctx) error {
	product := new(models.Products)
	id := c.Params("id")
	if err := config.DB.First(&product, id).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Product not found",
		})
	}

	if err := c.BodyParser(product); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}
	if err := config.DB.Save(&product).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Cannot update product",
		})
	}

	return c.JSON(product)
}

func DeleteProductByID(c *fiber.Ctx) error {
	product := new(models.Products)
	id := c.Params("id")

	if err := config.DB.First(&product, id).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Product not found",
		})
	}

	if err := config.DB.Delete(&product, id).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Cannot delete product",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Delete Successful!",
	})
}
