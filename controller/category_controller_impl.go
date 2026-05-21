package controller

import (
	"pinjam-buku/model/web"
	webCategory "pinjam-buku/model/web/category"
	"pinjam-buku/service"

	"github.com/gofiber/fiber/v2"
)

// membuat struct
type CategoryControllerImpl struct {
	// menambahkan category service
	// CategoryService adalah sebuah interface, maka tidak perlu menggunakan pointer
	CategoryService service.CategoryService
}

// menambahkan function untuk kebutuhan endpoint (bukan method milik CategoryController)

// membuat untuk router dengan return CategoryController (interface)
func NewCategoryController(categoryService service.CategoryService) CategoryController {
	// melakukan return dengan ponter
	return &CategoryControllerImpl{
		CategoryService: categoryService,
	}
}

// ini adalah kontrak untuk category controller
func (controller *CategoryControllerImpl) Create(ctx *fiber.Ctx) error {
	// membuat request body
	categoryCreateRequest := webCategory.CategoryCreateRequest{}

	// parsing body request
	err := ctx.BodyParser(&categoryCreateRequest)

	// jika error parsing
	if err != nil {
		return err
	}

	// memanggil service
	categoryResponse, err := controller.CategoryService.Create(
		ctx.Context(),
		categoryCreateRequest,
	)

	// jika error service
	if err != nil {
		return err
	}

	// return data
	return ctx.Status(fiber.StatusOK).JSON(
		web.WebResponse{
			Code:   200,
			Status: "OK",
			Data:   categoryResponse,
		},
	)
}

func (controller *CategoryControllerImpl) Update(ctx *fiber.Ctx) error {
	// request body
	categoryUpdateRequest := webCategory.CategoryUpdateRequest{}

	// parsing body
	err := ctx.BodyParser(&categoryUpdateRequest)

	// jika error parsing
	if err != nil {
		return err
	}

	// ambil parameter categoryId
	categoryId := ctx.Params("categoryId")

	// set id ke request
	categoryUpdateRequest.Id = categoryId

	// panggil service
	categoryResponse, err := controller.CategoryService.Update(
		ctx.Context(),
		categoryUpdateRequest,
	)

	// jika error service
	if err != nil {
		return err
	}

	// response
	return ctx.Status(fiber.StatusOK).JSON(
		web.WebResponse{
			Code:   200,
			Status: "OK",
			Data:   categoryResponse,
		},
	)
}

func (controller *CategoryControllerImpl) Delete(ctx *fiber.Ctx) error {
	// ambil parameter categoryId
	categoryId := ctx.Params("categoryId")

	// panggil service delete
	err := controller.CategoryService.Delete(
		ctx.Context(),
		categoryId,
	)

	// melakukan cek error
	if err != nil {
		return err
	}

	// response
	return ctx.Status(fiber.StatusOK).JSON(
		web.WebResponse{
			Code:   200,
			Status: "OK",
			Data:   "OK",
		},
	)
}

func (controller *CategoryControllerImpl) FindById(ctx *fiber.Ctx) error {
	// ambil parameter categoryId
	categoryId := ctx.Params("categoryId")

	// panggil service
	categoryResponse, err := controller.CategoryService.FindById(
		ctx.Context(),
		categoryId,
	)

	// jika error service
	if err != nil {
		return err
	}

	// response
	return ctx.Status(fiber.StatusOK).JSON(
		web.WebResponse{
			Code:   200,
			Status: "OK",
			Data:   categoryResponse,
		},
	)
}

func (controller *CategoryControllerImpl) FindAll(ctx *fiber.Ctx) error {
	// memanggil service
	categoryResponses, err := controller.CategoryService.FindAll(
		ctx.Context(),
	)

	// jika error service
	if err != nil {
		return err
	}

	// response
	return ctx.Status(fiber.StatusOK).JSON(
		web.WebResponse{
			Code:   200,
			Status: "OK",
			Data:   categoryResponses,
		},
	)
}
