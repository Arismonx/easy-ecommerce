package routes

import (
	middleware "github.com/Arismonx/easy-ecommerce/Middleware"
	"github.com/Arismonx/easy-ecommerce/controllers"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Get("/", controllers.Hello)
	//API user
	app.Post("/register", controllers.Register)
	app.Post("/login", controllers.LoginUser)

	//API Product
	app.Get("/api/product", middleware.AuthRequired, controllers.GetProducts)
	app.Get("/api/product/:id", middleware.AuthRequired, controllers.GetProductByID)
	app.Post("/api/product", middleware.AuthRequired, controllers.CreateProduct)
	app.Put("/api/product/:id", middleware.AuthRequired, controllers.UpdateProductByID)
	app.Delete("/api/product/:id", middleware.AuthRequired, controllers.DeleteProductByID)
}
