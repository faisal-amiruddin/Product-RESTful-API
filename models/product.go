package models

import (
    "time"
    "gorm.io/gorm"
)

type CreateProductInput struct {
    Title       string  `json:"title" example:"Laptop Gaming"`
    Description string  `json:"description" example:"Laptop spek dewa"`
    Price       float64 `json:"price" example:"15000000"`
    Stock       int16   `json:"stock" example:"10"`
}

type Product struct {
    ID        uint           `gorm:"primaryKey" json:"id"`
    Title       string  `json:"title"`
    Description string  `json:"description"`
    Price       float64 `json:"price"`
    Stock       int16   `json:"stock"`
    CreatedAt time.Time      `json:"created_at"`
    UpdatedAt time.Time      `json:"updated_at"`
    DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at" swaggertype:"string"`
}