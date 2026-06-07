package models

type User struct {
	ID        uint   `gorm:"primaryKey" json:"id"`
	Email     string `gorm:"uniqueIndex;not null;size:255" json:"email"`
	Password  string `gorm:"not null;size:255" json:"-"`
	Name      string `gorm:"not null;size:255" json:"name"`
}
