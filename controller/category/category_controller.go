package category

import (
	"github.com/gofiber/fiber/v2"
)

// membuat interface untuk controller nya
type CategoryController interface {
	// menambahkan kontrak untuk category controller
	Create(ctx *fiber.Ctx) error
	Update(ctx *fiber.Ctx) error
	Delete(ctx *fiber.Ctx) error
	FindById(ctx *fiber.Ctx) error
	FindAll(ctx *fiber.Ctx) error
}
