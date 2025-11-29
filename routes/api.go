package routes

import (
    "api-product/controllers"
    _ "api-product/docs" 
    "github.com/gofiber/fiber/v2"
    "github.com/gofiber/swagger" 
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

    app.Get("/swagger/*", swagger.New(swagger.Config{
        URL: "/swagger/doc.json",
    }))
    
}