package database

import (
	"log"
	"pinjam-buku/internal/models"

	"gorm.io/gorm"
)

func RunMigration(db *gorm.DB) {
	err := db.AutoMigrate(
		&models.User{},
		&models.Book{},
	)

	// mengecek error migration
	if err != nil {
		log.Fatal("Database migration failed")
	}

	log.Println("Successful database migration")
}