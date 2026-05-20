package exception

import (
	"errors"
	"pinjam-buku/model/web"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// membuat function untuk penangan error handler pada endpoint
// global error handler fiber
func ErrorHandler(ctx *fiber.Ctx, err error) error {

	// validation error
	var validationErrors validator.ValidationErrors

	if errors.As(err, &validationErrors) {
		return ctx.Status(fiber.StatusBadRequest).JSON(
			web.WebResponse{
				Code:   400,
				Status: "BAD REQUEST",
				Data:   validationErrors.Error(),
			},
		)
	}

	// gorm record not found
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return ctx.Status(fiber.StatusNotFound).JSON(
			web.WebResponse{
				Code:   404,
				Status: "NOT FOUND",
				Data:   "Data not found",
			},
		)
	}

	// custom not found error
	var notFoundError NotFoundError

	if errors.As(err, &notFoundError) {
		return ctx.Status(fiber.StatusNotFound).JSON(
			web.WebResponse{
				Code:   404,
				Status: "NOT FOUND",
				Data:   notFoundError.Message,
			},
		)
	}

	// default internal server error
	return ctx.Status(fiber.StatusInternalServerError).JSON(
		web.WebResponse{
			Code:   500,
			Status: "INTERNAL SERVER ERROR",
			Data:   err.Error(),
		},
	)
}
