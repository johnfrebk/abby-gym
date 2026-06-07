package models

import (
	"time"
)

type Sale struct {
	ID        uint           `gorm:"primaryKey"`
	ClientID  uint           `gorm:"not null;index:idx_sale_client_id"`
	Total     float64        `gorm:"not null"`

	CreatedAt time.Time      `gorm:"index:idx_sale_created_at"`
	UpdatedAt time.Time

	Client      Client         `gorm:"foreignKey:ClientID"`
	SaleDetails []SalesDetail  `gorm:"foreignKey:SaleID;constraint:OnDelete:CASCADE"`
}
