package main

import (
	"pinjam-buku/app"
	"pinjam-buku/database"
	"pinjam-buku/helper"

	categoryController "pinjam-buku/controller/category"
	categoryRepository "pinjam-buku/repository/category"
	categoryService "pinjam-buku/service/category"

	"github.com/go-playground/validator/v10"
)

func main() {
	// target :
	// 1. otentikasi / auth (register/login)

	// koneksi database
	db := app.NewDB()

	// hapus semua tabel (jika diperlukan)
	database.DropAllTables(db)

	// running auto migration
	database.RunMigration(db)

	// seeder (jika diperlukan/biasanya cukup sekali)
	database.RunSeeder(db)

	// membuat validator
	validate := validator.New()

	// membuat repository
	categoryRepository := categoryRepository.NewCategoryRepository()

	// membuat service
	categoryService := categoryService.NewCategoryService(categoryRepository, db, validate)

	// membuat category controller
	categoryController := categoryController.NewCategoryController(categoryService)

	// implementasi router
	router := app.NewRouter(categoryController)

	// menjalankan server
	err := router.Listen(":3000")

	// mengecek error
	helper.PanicIfError(err)
}
