package seeder

import (
	"log"
	"pinjam-buku/model/domain"

	"gorm.io/gorm"
)

func SeedCategory(db *gorm.DB) {
	// data category bawaan
	categories := []domain.Category{
		{Name: "Novel"},
		{Name: "Komik"},
		{Name: "Cerpen"},
		{Name: "Cergam"},
		{Name: "Dongeng"},
		{Name: "Referensi"},
		{Name: "Biografi"},
		{Name: "Autobiografi"},
		{Name: "Akademik"},
		{Name: "Sains"},
		{Name: "Matematika"},
		{Name: "Olahraga"},
		{Name: "Seni"},
		{Name: "Sastra"},
		{Name: "Geografi"},
		{Name: "Sejarah"},
		{Name: "Komputer"},
		{Name: "Filsafat"},
		{Name: "Psikologi"},
		{Name: "Pengembangan Diri"},
		{Name: "Karya Umum"},
		{Name: "Teknologi"},
	}

	for _, category := range categories {

		var existingCategory domain.Category

		err := db.
			Where("name = ?", category.Name).
			First(&existingCategory).
			Error

		if err == gorm.ErrRecordNotFound {

			createErr := db.Create(&category).Error

			if createErr != nil {
				log.Println("Seeder category failed:", createErr)
			}
		}
	}

	log.Println("Successful category seeder")
}
