package user

import "github.com/gofiber/fiber/v2"

// membuat interface untuk controller nya
type UserController interface {
	// menambahkan kontrak untuk user controller
	SignUp(ctx *fiber.Ctx) error
	SignIn(ctx *fiber.Ctx) error
	UpdateUser(ctx *fiber.Ctx) error
	SignOut(ctx *fiber.Ctx) error
	GetUserCurrent(ctx *fiber.Ctx) error
}
