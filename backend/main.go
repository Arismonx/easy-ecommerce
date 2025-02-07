package main

import (
	"github.com/Arismonx/easy-ecommerce/config"
	"github.com/Arismonx/easy-ecommerce/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	config.LoadENV()
	config.ConnectDB()
	// Create a new Fiber instance
	app := fiber.New()

	routes.Setup(app)

	app.Listen(":3000")
}
