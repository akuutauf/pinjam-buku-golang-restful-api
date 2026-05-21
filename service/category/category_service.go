package category

import (
	"context"
	"pinjam-buku/model/web/category"
)

// service disebut juga sebagai bussiness logic
// di dalam service harus membuat kontrak
// membuat interface (kontrak)
type CategoryService interface {
	// biasanya kontrak ini berisi api sesuai dengan jumlah function
	// contoh: sebuah api hanya bisa memanggil 1 function di sebuah service

	// isi parameter berisi context, request (optional), dan response (optional)
	// untuk request bisa kita custom seperti category_create_request
	// dan untuk responsenya bisa kita custom seperti category_response

	// kalau domain itu ranahnya ke repository tidak ke service. maka dari itu nanti request akan-
	// representasikan ke dalam bentuk yang lain
	// juga untuk response nya direkomendasikan jangan menggunakan domain, karena biasanya kalau-
	// ditampilkan dalam bentuk table semua informasi tabel akan ditampilkan

	Create(ctx context.Context, request category.CategoryCreateRequest) (category.CategoryResponse, error)
	Update(ctx context.Context, requst category.CategoryUpdateRequest) (category.CategoryResponse, error)
	Delete(ctx context.Context, categoryId string) error
	FindById(ctx context.Context, categoryId string) (category.CategoryResponse, error)
	FindAll(ctx context.Context) ([]category.CategoryResponse, error)
}
