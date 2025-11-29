package routes

import (
    "api-product/controllers"
    "github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
    api := app.Group("/api")
	v1 := api.Group("/v1")

    v1.Get("/products", controllers.GetProducts)
    v1.Post("/products", controllers.CreateProduct)
}