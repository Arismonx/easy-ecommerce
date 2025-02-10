package controllers

import (
	"time"

	"github.com/Arismonx/easy-ecommerce/config"
	"github.com/Arismonx/easy-ecommerce/models"
	"github.com/gofiber/fiber/v2"
)

func CreateOrder(c *fiber.Ctx) error {
	newOrder := new(models.Orders)

	if err := c.BodyParser(newOrder); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}

	// Check UserID delete?
	var user models.Users
	if err := config.DB.Unscoped().Where("id = ? AND deleted_at IS NOT NULL", newOrder.UserID).First(&user).Error; err == nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "User has been deleted",
		})
	}
	newOrder.OrderAddress = user.Address
	newOrder.OrderDate = time.Now()
	if newOrder.OrderStatus == "" {
		newOrder.OrderStatus = "Draft"
	}

	if newOrder.UserID == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "UserID are Required",
		})
	}

	if err := config.DB.Create(newOrder).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Cannot Create Order",
		})
	}

	if err := config.DB.Preload("User").First(&newOrder, newOrder.ID).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Cannot load related data",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(newOrder)
}

func GetOrders(c *fiber.Ctx) error {
	var order []models.Orders

	if err := config.DB.Preload("User").Find(&order).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Cannot get orders",
		})
	}

	return c.JSON(order)
}

func GetOrderByID(c *fiber.Ctx) error {
	order := new(models.Orders)
	id := c.Params("id")

	if err := config.DB.Preload("User").First(&order, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Cart not Found",
		})
	}
	return c.JSON(order)
}
