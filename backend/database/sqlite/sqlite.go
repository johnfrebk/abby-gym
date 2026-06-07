package sqlite

import (
	"log"
	"os"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
	"POS/backend/database/models"
	"POS/backend/utils"
)

var DB *gorm.DB

func Init() {
	var err error
	var dbPath string

	envPath := os.Getenv("DB_PATH")
	if envPath != "" {
		dbPath = envPath
	} else {
		dbPath, err = utils.GetDatabasePath("AbbyGym", "GYM.db")
		if err != nil {
			log.Fatal("No se pudo montar la base de datos en el sistema operativo actual")
		}
	}

	DB, err = gorm.Open(sqlite.Open(dbPath+"?_pragma=journal_mode(WAL)&_pragma=busy_timeout(5000)"), &gorm.Config{})
	if err != nil {
		log.Fatal("Error al conectar con SQLite:", err)
	}

	log.Println("Base de datos SQLite conectada")

	err = autoMigrate()
	if err != nil {
		log.Fatal("Error en migraciones:", err)
	}

	log.Println("Migraciones aplicadas")
}

func autoMigrate() error {	
	return DB.AutoMigrate(
		&models.User{},
		&models.Client{},
		&models.Product{},
		&models.Sale{},
		&models.SalesDetail{},
		&models.Membership{},
		&models.Subscription{},
		&models.ActivityLog{},
	)
}
