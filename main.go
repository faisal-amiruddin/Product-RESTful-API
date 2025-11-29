package main

import (
    "api-product/database"
    "api-product/models"
    "api-product/routes"
    "log"
    "os"

    "github.com/gofiber/fiber/v2"
    "github.com/joho/godotenv"
)

func main() {
    if err := godotenv.Load(); err != nil {
        log.Fatal("Error loading .env file")
    }

    database.Connect()

    database.DB.AutoMigrate(&models.Product{})

    app := fiber.New()

    routes.Setup(app)

    port := os.Getenv("PORT")
    app.Listen(":" + port)
}