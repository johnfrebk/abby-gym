package models

import (
	"time"
)

type Product struct {
	ID        uint           `gorm:"primaryKey"`
	Name      string         `gorm:"not null;index:idx_product_name"`
	Price     float64        `gorm:"not null"`
	Stock     int            `gorm:"not null"`
	CreatedAt time.Time      `gorm:"index:idx_product_created_at"`
	UpdatedAt time.Time

	SaleDetails []SalesDetail `gorm:"foreignKey:ProductID"`
}