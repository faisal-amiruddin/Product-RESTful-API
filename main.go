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

// @title           Simple Product REST API
// @version         1.0
// @description     REST API menggunakan Fiber dan GORM.

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:3000
// @BasePath  /api/v1
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