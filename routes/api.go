package routes

import (
    "api-product/controllers"
    "github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
    api := app.Group("/api")
	v1 := api.Group("/v1")

	product := v1.Group("/products")

    product.Get("/", controllers.GetProducts)
    product.Post("/", controllers.CreateProduct)
	product.Get("/:id", controllers.GetSingleProduct)
    product.Put("/:id", controllers.UpdateProduct)
    product.Delete("/:id", controllers.DeleteProduct)
}