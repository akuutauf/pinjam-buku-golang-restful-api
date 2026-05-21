package user

import (
	"context"
	"pinjam-buku/model/web/user"
)

// service disebut juga sebagai bussiness logic
// di dalam service harus membuat kontrak
// membuat interface (kontrak)
type UserService interface {
	// biasanya kontrak ini berisi api sesuai dengan jumlah function
	// contoh: sebuah api hanya bisa memanggil 1 function di sebuah service

	// isi parameter berisi context, request (optional), dan response (optional)
	// untuk request bisa kita custom seperti category_create_request
	// dan untuk responsenya bisa kita custom seperti category_response

	SignUp(ctx context.Context, request user.UserCreateRequest) (user.UserResponse, error)
	SignIn(ctx context.Context, request user.UserAuthRequest) (user.UserAuthResponse, error)
	UpdateUser(ctx context.Context, request user.UserUpdateRequest) (user.UserResponse, error)
	SignOut(ctx context.Context, userId string) error
	GetUserCurrent(ctx context.Context, userId string) (user.UserResponse, error)
}
