package app

import (
	"pinjam-buku/controller"

	"github.com/gofiber/fiber/v2"
)

func RegisterCategoryRoute(
	router *fiber.App,
	categoryController controller.CategoryController,
) {
	// kumpulan route dibagi per group
	category := router.Group("/api/categories")
	category.Get("/", categoryController.FindAll)
	category.Get("/:categoryId", categoryController.FindById)
	category.Post("/", categoryController.Create)
	category.Put("/:categoryId", categoryController.Update)
	category.Delete("/:categoryId", categoryController.Delete)
}
