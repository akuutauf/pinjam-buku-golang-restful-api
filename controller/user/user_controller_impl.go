package user

import (
	"pinjam-buku/model/web"
	webUser "pinjam-buku/model/web/user"
	"pinjam-buku/service/user"

	"github.com/gofiber/fiber/v2"
)

// membuat struct untuk user controller impl
type UserControllerImpl struct {
	// menambahkan user service
	UserService user.UserService
}

// membuat untuk router dengan return UserController (interface)
func NewUserController(userService user.UserService) UserController {
	// melakukan return dengan ponter
	return &UserControllerImpl{
		UserService: userService,
	}
}

// ini adalah kontrak untuk user controller
func (controller *UserControllerImpl) SignUp(ctx *fiber.Ctx) error {
	// membuat request body
	request := webUser.UserCreateRequest{}

	// parsing body request
	err := ctx.BodyParser(&request)

	// melakukan pengecekan jika error pada saat parsing
	if err != nil {
		return err
	}

	// memanggil service sign up
	response, err := controller.UserService.SignUp(
		ctx.Context(),
		request,
	)

	// melakukan pengecekan jika terdpat error saat running service
	if err != nil {
		return err
	}

	// return data hasil request
	return ctx.Status(fiber.StatusOK).JSON(
		web.WebResponse{
			Code:   200,
			Status: "OK",
			Data:   response,
		},
	)
}

func (controller *UserControllerImpl) SignIn(ctx *fiber.Ctx) error {
	// membuat request body
	request := webUser.UserAuthRequest{}

	// parsing body request
	err := ctx.BodyParser(&request)

	// melakukan pengecekan jika error pada saat parsing
	if err != nil {
		return err
	}

	// memanggil service sign in
	response, err := controller.UserService.SignIn(
		ctx.Context(),
		request,
	)

	// melakukan pengecekan jika terdpat error saat running service
	if err != nil {
		return err
	}

	// return data hasil request
	return ctx.Status(fiber.StatusOK).JSON(
		web.WebResponse{
			Code:   200,
			Status: "OK",
			Data:   response,
		},
	)
}

func (controller *UserControllerImpl) UpdateUser(ctx *fiber.Ctx) error {
	// request body
	request := webUser.UserUpdateRequest{}

	// parsing body
	err := ctx.BodyParser(&request)

	// melakukan pengecekan jika error pada saat parsing
	if err != nil {
		return err
	}

	// ambil parameter userId menggunakan user current
	userId := ctx.Locals("userId").(string)

	// set id ke request
	request.Id = userId

	// memanggil service update user
	response, err := controller.UserService.UpdateUser(
		ctx.Context(),
		request,
	)

	// melakukan pengecekan jika terdpat error saat running service
	if err != nil {
		return err
	}

	// return data hasil request
	return ctx.Status(fiber.StatusOK).JSON(
		web.WebResponse{
			Code:   200,
			Status: "OK",
			Data:   response,
		},
	)
}

func (controller *UserControllerImpl) SignOut(ctx *fiber.Ctx) error {
	// ambil parameter userId menggunakan user current
	userId := ctx.Locals("userId").(string)

	// panggil service sign out
	err := controller.UserService.SignOut(
		ctx.Context(),
		userId,
	)

	// melakukan pengecekan jika terdpat error saat running service
	if err != nil {
		return err
	}

	// return data hasil request
	return ctx.Status(fiber.StatusOK).JSON(
		web.WebResponse{
			Code:   200,
			Status: "OK",
			Data:   "OK",
		},
	)
}

func (controller *UserControllerImpl) GetUserCurrent(ctx *fiber.Ctx) error {
	// ambil userId dari middleware auth
	userId := ctx.Locals("userId").(string)

	// panggil service get user current
	userResponse, err := controller.UserService.GetUserCurrent(
		ctx.Context(),
		userId,
	)

	// melakukan pengecekan jika terdpat error saat running service
	if err != nil {
		return err
	}

	// return data hasil request
	return ctx.Status(fiber.StatusOK).JSON(
		web.WebResponse{
			Code:   200,
			Status: "OK",
			Data:   userResponse,
		},
	)
}
