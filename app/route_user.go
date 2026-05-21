package app

import (
	"pinjam-buku/controller/user"

	"github.com/gofiber/fiber/v2"
)

// mendaftarkan route untuk endpoint user
func RegisterUserRoute(
	router *fiber.App,
	userController user.UserController,
) {
	// kumpulan route dibagi per group
	user := router.Group("/api/users")

	// route untuk masing masing endpoint
	user.Post("/register", userController.SignUp)
	user.Post("/login", userController.SignIn)
	user.Get("/current", userController.GetUserCurrent)
	user.Put("/:userId", userController.UpdateUser)
	user.Delete("/userId", userController.SignOut)
}
