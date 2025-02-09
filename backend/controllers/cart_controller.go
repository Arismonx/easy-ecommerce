package controllers

import (
	"github.com/Arismonx/easy-ecommerce/config"
	"github.com/Arismonx/easy-ecommerce/models"
	"github.com/gofiber/fiber/v2"
)

func CreateCart(c *fiber.Ctx) error {
	newCart := new(models.Cart)

	if err := c.BodyParser(newCart); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}

	if newCart.ProductID == 0 || newCart.UserID == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "ProductID and UserID are Required",
		})
	}
	// Check ProductID delete?
	var product models.Products
	if err := config.DB.Unscoped().Where("id = ? AND deleted_at IS NOT NULL", newCart.ProductID).First(&product).Error; err == nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Product has been deleted",
		})
	}
	// Check ProductID delete?
	var user models.Users
	if err := config.DB.Unscoped().Where("id = ? AND deleted_at IS NOT NULL", newCart.ProductID).First(&user).Error; err == nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "User has been deleted",
		})
	}

	if err := config.DB.Create(newCart).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Cannot create cart",
		})
	}

	// Preload Product and User data
	if err := config.DB.Preload("Product").Preload("User").First(&newCart, newCart.ID).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Cannot load related data",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(newCart)
}

func GetCarts(c *fiber.Ctx) error {
	var carts []models.Cart
	if err := config.DB.Preload("Product").Preload("User").Find(&carts).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Cannot get Carts",
		})
	}
	return c.JSON(carts)
}

func GetCartByID(c *fiber.Ctx) error {
	cart := new(models.Cart)
	id := c.Params("id")

	if err := config.DB.Preload("Product").Preload("User").First(&cart, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Cart not Found",
		})
	}
	return c.JSON(cart)
}
