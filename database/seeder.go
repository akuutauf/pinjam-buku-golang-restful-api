package database

import (
	"pinjam-buku/database/seeder"

	"gorm.io/gorm"
)

// membuat fungsi untuk menjalankan semua seeder
func RunSeeder(db *gorm.DB) {

	// jalankan semua seeder
	seeder.SeedCategory(db)
	// seeder.SampleSeed(db) .. menjalankan seeder yang lain bila perlu
}
