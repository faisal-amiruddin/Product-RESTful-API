package controllers

import (
	"api-product/database"
	"api-product/models"
	"github.com/gofiber/fiber/v2"
)

// GetProducts godoc
// @Summary      Get all products
// @Description  Mengambil daftar semua produk dari database
// @Tags         Products
// @Accept       json
// @Produce      json
// @Success      200 {object} map[string][]models.Product "Returns list of products"
// @Router       /products [get]
func GetProducts(c *fiber.Ctx) error {
	var products []models.Product
	
	database.DB.Find(&products)

	return c.JSON(fiber.Map{
		"status": "success",
		"data":   products,
	})
}

// CreateProduct godoc
// @Summary      Create a new product
// @Description  Membuat entri produk baru di database
// @Tags         Products
// @Accept       json
// @Produce      json
// @Param        product body models.CreateProductInput true "Product data input"
// @Success      201 {object} map[string]models.Product "Returns created product data"
// @Failure      400 {object} map[string]string "Input tidak valid"
// @Router       /products [post]
func CreateProduct(c *fiber.Ctx) error {
	input := new(models.CreateProductInput)

	if err := c.BodyParser(input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Input tidak valid",
		})
	}

	product := models.Product{
		Title:       input.Title,
		Description: input.Description,
		Price:       input.Price,
		Stock:       input.Stock,
	}

	database.DB.Create(&product)

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status": "success",
		"data":   product,
	})
}

// GetSingleProduct godoc
// @Summary      Get a single product by ID
// @Description  Mengambil data produk berdasarkan ID
// @Tags         Products
// @Accept       json
// @Produce      json
// @Param        id path int true "Product ID"
// @Success      200 {object} map[string]models.Product "Returns product data"
// @Failure      404 {object} map[string]string "Product not found"
// @Router       /products/{id} [get]
func GetSingleProduct(c *fiber.Ctx) error {
	id := c.Params("id")
	var product models.Product

	if result := database.DB.First(&product, id); result.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  "error",
			"message": "Produk tidak ditemukan",
		})
	}

	return c.JSON(fiber.Map{
		"status": "success",
		"data":   product,
	})
}

// UpdateProduct godoc
// @Summary      Update an existing product
// @Description  Memperbarui data produk yang sudah ada
// @Tags         Products
// @Accept       json
// @Produce      json
// @Param        id path int true "Product ID"
// @Param        product body models.CreateProductInput true "Updated product data"
// @Success      200 {object} map[string]models.Product "Product berhasil diperbarui"
// @Failure      400 {object} map[string]string "Input tidak valid"
// @Failure      404 {object} map[string]string "Product not found"
// @Router       /products/{id} [put]
func UpdateProduct(c *fiber.Ctx) error {
	id := c.Params("id")
	var product models.Product
	if result := database.DB.First(&product, id); result.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  "error",
			"message": "Produk tidak ditemukan untuk diupdate",
		})
	}

	input := new(models.CreateProductInput)
	if err := c.BodyParser(input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "Input tidak valid",
		})
	}

	database.DB.Model(&product).Updates(models.Product{
		Title:       input.Title,
		Description: input.Description,
		Price:       input.Price,
		Stock:       input.Stock,
	})

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "Produk berhasil diperbarui",
		"data":    product,
	})
}

// DeleteProduct godoc
// @Summary      Delete a product
// @Description  Menghapus produk berdasarkan ID (Soft Delete GORM)
// @Tags         Products
// @Accept       json
// @Produce      json
// @Param        id path int true "Product ID"
// @Success      204 {string} string "No Content"
// @Failure      404 {object} map[string]string "Product not found"
// @Router       /products/{id} [delete]
func DeleteProduct(c *fiber.Ctx) error {
	id := c.Params("id")
	var product models.Product

	if result := database.DB.First(&product, id); result.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  "error",
			"message": "Produk tidak ditemukan untuk dihapus",
		})
	}

	database.DB.Delete(&product)

	return c.SendStatus(fiber.StatusNoContent)
}