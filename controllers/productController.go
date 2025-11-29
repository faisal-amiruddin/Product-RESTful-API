package controllers

import (
    "api-product/database"
    "api-product/models"
    "github.com/gofiber/fiber/v2"
)

func GetProducts(c *fiber.Ctx) error {
    var products []models.Product
    
    database.DB.Find(&products)

    return c.JSON(fiber.Map{
        "status": "success",
        "data": products,
    })
}

func CreateProduct(c *fiber.Ctx) error {
    product := new(models.Product)

    if err := c.BodyParser(product); err != nil {
        return c.Status(503).JSON(fiber.Map{
            "status":  "error",
            "message": "Input tidak valid",
        })
    }

    database.DB.Create(&product)

    return c.Status(201).JSON(fiber.Map{
        "status": "success",
        "data": product,
    })
}