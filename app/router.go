package app

import (
	category "pinjam-buku/controller/category"
	// user "pinjam-buku/controller/user"
	"pinjam-buku/exception"
	"pinjam-buku/middleware"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func NewRouter(
	categoryController category.CategoryController,
	// userController user.UserController,
) *fiber.App {

	// create fiber app
	router := fiber.New(fiber.Config{
		Prefork: false,

		// global error handler
		ErrorHandler: exception.ErrorHandler,
	})

	// global middleware
	middleware.LoggerMiddleware(router)

	// recover panic middleware
	router.Use(recover.New())

	// register routes
	RegisterCategoryRoute(router, categoryController)
	// RegisterUserRoute(router, userController)

	return router
}
