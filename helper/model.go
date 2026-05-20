package helper

import (
	"pinjam-buku/model/domain"
	"pinjam-buku/model/web/category"
)

// membuat function untuk mengkonversi data menjadi tipe Category Response
func ToCategoryResponse(data domain.Category) category.CategoryResponse {
	return category.CategoryResponse{
		ID:   data.ID,
		Name: data.Name,
	}
}

// membuat function untuk mengkonversi data slice menjadi slice Category Response
func ToCategoryResponses(categories []domain.Category) []category.CategoryResponse {
	// membuat variabel baru (berjenis slice category response)
	var categoryResponses []category.CategoryResponse

	// melakukan konversi data categories (slice), ke category responses (slice)
	for _, category := range categories {
		categoryResponses = append(categoryResponses, ToCategoryResponse(category))
	}

	// karena data category sudah dikonversi ke category responses (slice), maka langsung bisa di return
	return categoryResponses
}
