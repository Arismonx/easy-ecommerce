package main

import (
	"github.com/Arismonx/easy-ecommerce/config"
	"github.com/Arismonx/easy-ecommerce/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	config.LoadENV()
	config.ConnectDB()
	// Create a new Fiber instance
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:3000",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))
	routes.Setup(app)

	app.Listen(":8080")
}
