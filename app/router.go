package app

import (
	"pinjam-buku/controller"
	"pinjam-buku/exception"

	"github.com/gofiber/fiber/v2"
)

// membuat function untuk membuat route baru
func NewRouter(categoryController controller.CategoryController) *fiber.App {
	// membuat fiber app
	router := fiber.New(fiber.Config{
		// prefork cocok digunakan ketika aplikasi sudah di hosting di server
		Prefork: false,

		// global error handler
		ErrorHandler: exception.ErrorHandler,
	})

	// endpoint category
	router.Get("/api/categories", categoryController.FindAll)
	router.Get("/api/categories/:categoryId", categoryController.FindById)
	router.Post("/api/categories", categoryController.Create)
	router.Put("/api/categories/:categoryId", categoryController.Update)
	router.Delete("/api/categories/:categoryId", categoryController.Delete)

	return router
}
