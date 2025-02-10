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

	newOrder.OrderDate = time.Now()
	newOrder.OrderAddress = user.Address
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
