package controllers

import (
	"github.com/Arismonx/easy-ecommerce/config"
	"github.com/Arismonx/easy-ecommerce/models"
	"github.com/gofiber/fiber/v2"
)

func CreateOrderline(c *fiber.Ctx) error {
	newOrderline := new(models.Orderlines)

	if err := c.BodyParser(newOrderline); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}

	var product models.Products
	if err := config.DB.First(&product, newOrderline.ProductID).Error; err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Product not found",
		})
	}

	var order models.Orders
	if err := config.DB.First(&order, newOrderline.OrderID).Error; err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Order not found",
		})
	}

	// Set OrderLineDescription, OrderLineQuantity, and OrderLineUnitPrice from Product
	newOrderline.OrderLineDescription = product.ProductDescription
	newOrderline.OrderLineQuantity = product.ProductQuantity
	newOrderline.OrderLineUnitPrice = product.ProductPrice

	// Calculate and set OrderLinePriceSubtotal
	newOrderline.OrderLinePriceSubtotal = newOrderline.OrderLineQuantity * newOrderline.OrderLineUnitPrice

	if err := config.DB.Create(newOrderline).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Cannot create Orderline",
		})
	}

	if err := config.DB.Preload("Product").Preload("Order").First(&newOrderline, newOrderline.ID).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Cannot load related data",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(newOrderline)
}

func GetOrderlines(c *fiber.Ctx) error {
	var orderline []models.Orderlines

	if err := config.DB.Preload("Product").Preload("Order").Find(&orderline).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Cannot get orderlines",
		})
	}

	return c.JSON(orderline)
}

func GetOrderlineByID(c *fiber.Ctx) error {
	orderlien := new(models.Orderlines)
	id := c.Params("id")

	if err := config.DB.First(orderlien, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Orderlien not found",
		})
	}

	return c.JSON(orderlien)
}

func DeleteOrderlineByID(c *fiber.Ctx) error {
	orderline := new(models.Orderlines)
	id := c.Params("id")

	if err := config.DB.First(&orderline, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Orderline not found",
		})
	}

	if err := config.DB.Delete(&orderline, id).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Cannot delete orderline",
		})
	}

	return c.Status(fiber.StatusNoContent).JSON(fiber.Map{
		"message": "Delete Successful",
	})
}
