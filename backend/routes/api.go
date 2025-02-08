package routes

import (
	"github.com/Arismonx/easy-ecommerce/controllers"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Get("/", controllers.Hello)

	//API user
	// app.Post("/api/user", controllers.CreateUser)

	//API Product
	app.Get("/api/product", controllers.GetProducts)
	app.Post("/api/product", controllers.CreateProduct)
}
