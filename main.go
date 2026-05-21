package main

import (
	"pinjam-buku/app"
	"pinjam-buku/controller"

	// "pinjam-buku/database"
	"pinjam-buku/helper"
	repository "pinjam-buku/repository/category"
	"pinjam-buku/service"

	"github.com/go-playground/validator/v10"
)

func main() {
	// target :
	// 1. otentikasi / auth (register/login)

	// koneksi database
	db := app.NewDB()

	// hapus semua tabel (jika diperlukan)
	// database.DropAllTables(db)

	// running auto migration
	// database.RunMigration(db)

	// seeder (jika diperlukan/biasanya cukup sekali)
	// database.RunSeeder(db)

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

	// menjalankan server
	err := router.Listen(":3000")

	// mengecek error
	helper.PanicIfError(err)
}
