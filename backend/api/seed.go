package api

import (
	"log"

	"POS/backend/database/models"
	"POS/backend/database/sqlite"

	"golang.org/x/crypto/bcrypt"
)

func SeedDefaultAdmin() {
	var count int64
	sqlite.DB.Model(&models.User{}).Count(&count)
	if count > 0 {
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte("admin123"), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("error al crear admin por defecto: %v", err)
		return
	}

	admin := models.User{
		Email:    "admin@abbygym.com",
		Password: string(hash),
		Name:     "Administrador",
	}

	if err := sqlite.DB.Create(&admin).Error; err != nil {
		log.Printf("error al crear admin por defecto: %v", err)
		return
	}

	log.Println("Usuario admin creado: admin@abbygym.com / admin123")
}
