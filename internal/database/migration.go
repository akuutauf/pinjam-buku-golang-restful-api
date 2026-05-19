package database

import (
	"log"
	"pinjam-buku/internal/models"

	"gorm.io/gorm"
)

// migrate table
func RunMigration(db *gorm.DB) {
	err := db.AutoMigrate(
		&models.User{},
		&models.Book{},
		&models.ProfileUser{},
		&models.Loan{},
	)

	// mengecek error migration
	if err != nil {
		log.Fatal("Database migration failed")
	}

	log.Println("Successful database migration")
}

// hapus semua tabel
func DropAllTables(db *gorm.DB) {
	err := db.Migrator().DropTable(
		// urutan tabel dihapus dari paling bawah, agar tidak terjadi erorr FK 
		&models.Loan{},
		&models.Book{},
		&models.ProfileUser{},
		&models.User{},
	)

	if err != nil {
		log.Fatal("Drop table failed")
	}

	log.Println("All tables dropped successfully")
}