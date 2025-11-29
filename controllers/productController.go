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

func GetSingleProduct(c *fiber.Ctx) error {
	id := c.Params("id")
	var product models.Product

	if result := database.DB.First(&product, id); result.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status": "error",
			"message": "Produk tidak ditemukan",
		})
	}

	return c.JSON(fiber.Map{
		"status": "success",
		"data": product,
	})
}

func UpdateProduct(c *fiber.Ctx) error {
	id := c.Params("id")
	product := new(models.Product)
	productUpdate := new(models.Product)

	if err := c.BodyParser(productUpdate); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": "error",
			"message": "Input tidak valid",
		})
	}

	if result := database.DB.First(&product, id); result.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status": "error",
			"message": "Produk tidak ditemukan untuk diupdate",
		})
	}

	database.DB.Model(&product).Updates(productUpdate)

	return c.JSON(fiber.Map{
		"status": "success",
		"message": "Produk berhasil diperbarui",
		"data": product,
	})
}

func DeleteProduct(c *fiber.Ctx) error {
	id := c.Params("id")
	var product models.Product

	if result := database.DB.First(&product, id); result.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status": "error",
			"message": "Produk tidak ditemukan untuk dihapus",
		})
	}

	database.DB.Delete(&product)

	return c.SendStatus(fiber.StatusNoContent)
}