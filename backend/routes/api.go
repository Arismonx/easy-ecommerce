package routes

import (
	"github.com/Arismonx/easy-ecommerce/controllers"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Get("/", controllers.Hello)

	//API user
	app.Post("/register", controllers.Register)
	app.Post("/login", controllers.LoginUser)

	//API Product
	app.Get("/api/product", controllers.GetProducts)
	app.Get("/api/product/:id", controllers.GetProductByID)
	app.Post("/api/product", controllers.CreateProduct)
	app.Put("/api/product/:id", controllers.UpdateProductByID)
	app.Delete("/api/product/:id", controllers.DeleteProductByID)
}
