package main

import (
	"net/http"
	"pinjam-buku/app"
	"pinjam-buku/controller"
	"pinjam-buku/database"
	"pinjam-buku/helper"
	"pinjam-buku/middleware"
	repository "pinjam-buku/repository/category"
	"pinjam-buku/service"

	"github.com/go-playground/validator/v10"
)

func main() {
	// koneksi database
	db := app.NewDB()

	// hapus semua tabel (jika diperlukan)
	database.DropAllTables(db)

	// running auto migration
	database.RunMigration(db)

	// membuat validator
	validate := validator.New()

	// membuat repository
	categoryRepository := repository.NewCategoryRepository()

	// membuat service
	categoryService := service.NewCategoryService(categoryRepository, db, validate)

	// membuat category controller
	categoryController := controller.NewCategoryController(categoryService)
	
	// implementasi router
	router := app.NewRouter(categoryController)

	// membuat server
	server := http.Server{
		Addr: "localhost:3000",
		Handler: middleware.NewAuthMiddleware(router),

		// ketika auth middleware sudah ditambahkan, maka handler yang digunakan adalah-
		// router yang dibungkus dengan auth middleware
	}

	// menjalankan server
	err := server.ListenAndServe()

	// mengecek error
	helper.PanicIfError(err)
}