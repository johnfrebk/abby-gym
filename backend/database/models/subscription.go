package models

import (
	"time"
)

type Subscription struct {
	ID           uint      `gorm:"primaryKey"`
	ClientID     uint      `gorm:"not null;index:idx_subscription_client_id"`
	MembershipID uint      `gorm:"not null;index:idx_subscription_membership_id"`
	Price        float64   `gorm:"not null"`
	StartDate    string    `gorm:"not null;index:idx_subscription_start_date"`
	EndDate      string    `gorm:"not null;index:idx_subscription_end_date"`
	CreatedAt    time.Time `gorm:"index:idx_subscription_created_at"`
	UpdatedAt    time.Time

	Client     Client     `gorm:"foreignKey:ClientID"`
	Membership Membership `gorm:"foreignKey:MembershipID"`
}
