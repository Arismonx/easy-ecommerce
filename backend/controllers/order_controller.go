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
	orderID := c.Params("id")

	if err := config.DB.Preload("User").First(&order, orderID).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Order not found",
		})
	}
	return c.JSON(order)
}

func GetOrdersByUserID(c *fiber.Ctx) error {
	var orders []models.Orders
	userID := c.Query("userID")

	if userID != "" {
		if err := config.DB.Preload("User").Where("user_id = ?", userID).Find(&orders).Error; err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Orders not found",
			})
		}
	} else {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "UserID is required",
		})
	}

	return c.JSON(orders)
}

// func UpdateOrderByID(c *fiber.Ctx) error {
// 	updateCart := new(models.Cart)
// 	id := c.Params("id")

// 	// Load the existing cart
// 	if err := config.DB.First(&updateCart, id).Error; err != nil {
// 		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
// 			"error": "Cart not found",
// 		})
// 	}

// 	if err := c.BodyParser(updateCart); err != nil {
// 		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
// 			"error": "Cannot parse JSON",
// 		})
// 	}

// 	if updateCart.ProductID == 0 || updateCart.UserID == 0 {
// 		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
// 			"error": "ProductID and UserID are Required",
// 		})
// 	}

// 	// Check ProductID delete?
// 	var product models.Products
// 	if err := config.DB.Unscoped().Where("id = ? AND deleted_at IS NOT NULL", updateCart.ProductID).First(&product).Error; err == nil {
// 		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
// 			"error": "Product has been deleted",
// 		})
// 	}
// 	// Check ProductID delete?
// 	var user models.Users
// 	if err := config.DB.Unscoped().Where("id = ? AND deleted_at IS NOT NULL", updateCart.ProductID).First(&user).Error; err == nil {
// 		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
// 			"error": "User has been deleted",
// 		})
// 	}

// 	if err := config.DB.Save(&updateCart).Error; err != nil {
// 		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
// 			"error": "Cannot Update cart",
// 		})
// 	}

// 	// Preload Product and User data
// 	if err := config.DB.Preload("Product").Preload("User").First(&updateCart, updateCart.ID).Error; err != nil {
// 		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
// 			"error": "Cannot load related data",
// 		})
// 	}
// 	return c.JSON(updateCart)

// }
