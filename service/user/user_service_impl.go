package user

// import (
// 	"context"
// 	"pinjam-buku/helper"
// 	userHelper "pinjam-buku/helper/user"
// 	"pinjam-buku/model/domain"
// 	"pinjam-buku/model/web/user"
// 	repository "pinjam-buku/repository/user"

// 	"github.com/go-playground/validator/v10"
// 	"gorm.io/gorm"
// )

// // membuat struct untuk user service impl
// type UserServiceImpl struct {
// 	// di user service implementation membutuhkan repository, maka perlu di tambahkan di atribut
// 	UserRepository repository.UserRepository

// 	// butuh koneksi database juga, maka tambahkan attribute sql
// 	DB *gorm.DB

// 	// menambahkan attribute validate
// 	Validate *validator.Validate
// }

// // membuat untuk router dengan return UserService (interface)
// func NewUserService(userRepository repository.UserRepository, DB *gorm.DB, validate *validator.Validate) UserService {
// 	return &UserServiceImpl{
// 		UserRepository: userRepository,
// 		DB:             DB,
// 		Validate:       validate,
// 	}
// }

// // mengimplementasi user service (membuat method yang ada di user  agar dimiliki oleh UserServiceImpl)
// func (service UserServiceImpl) SignUp(ctx context.Context, request user.UserCreateRequest) (user.UserResponse, error) {
// 	// mengimplementasi validation sebelum melakukan transaction
// 	err := service.Validate.Struct(request)

// 	// mengecek error
// 	if err != nil {
// 		return user.UserResponse{}, err
// 	}

// 	// memulai koneksi database transactional
// 	tx := service.DB.Begin()

// 	// cek transaction error
// 	if tx.Error != nil {
// 		return user.UserResponse{}, tx.Error
// 	}

// 	// mencegah ketika terjadi error setelah proses selesai semuanya dengan defer (pencegahan terakhir jika terjadi error)
// 	defer func() {
// 		helper.CommitOrRollback(tx, err)
// 	}()

// 	// user entity
// 	userDomain := domain.User{
// 		Username: request.Username,
// 		Email:    request.Email,
// 		Password: request.Password,
// 	}

// 	// save user
// 	userDomain, err = service.UserRepository.SignUp(
// 		ctx,
// 		tx,
// 		userDomain,
// 	)

// 	// check error save
// 	if err != nil {
// 		return user.UserResponse{}, err
// 	}

// 	// return response
// 	return userHelper.ToUserResponse(userDomain), nil
// }
