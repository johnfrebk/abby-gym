package models

import (
	"time"
)

type ActivityLog struct {
	ID        uint      `gorm:"primaryKey"`
	Entity    string    `gorm:"not null;index:idx_activity_entity"`
	EntityID  uint      `gorm:"not null;index:idx_activity_entity_id"`
	Action    string    `gorm:"not null"`
	Summary   string    `gorm:"size:255"`
	CreatedAt time.Time `gorm:"index:idx_activity_created_at"`
}