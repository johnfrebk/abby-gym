package models

import (
	"time"
)

type Client struct {
	ID               uint           `gorm:"primaryKey"`
	FirstName        string         `gorm:"not null"`
	LastName         string         `gorm:"not null"`
	Email            string         `gorm:"index:idx_client_email"`
	Phone            string
	DNI              string         `gorm:"column:dni;index:idx_client_dni"`
	RegistrationDate string         `gorm:"not null"`
	CreatedAt        time.Time      `gorm:"index:idx_client_created_at"`
	UpdatedAt        time.Time

	Subscriptions []Subscription `gorm:"foreignKey:ClientID"`
	Sales         []Sale         `gorm:"foreignKey:ClientID"`
}
