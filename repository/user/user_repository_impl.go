package user

// import (
// 	"context"
// 	"pinjam-buku/model/domain"

// 	"gorm.io/gorm"
// )

// // membuat implementasi dari repository (kontrak sebelumnya) untuk data user

// // membuat struct user implementation
// type UserRepositoryImpl struct {
// 	//
// }

// // mendefinisikan constructor
// func NewUserRepository() UserRepository {
// 	return &UserRepositoryImpl{}
// }

// // membuat method milik struct UserRepositoryImpl yang mana menerapkan kontrak sebelumnya
// func (repository *UserRepositoryImpl) SignUp(ctx context.Context, tx *gorm.DB, user domain.User) (domain.User, error) {

// 	// insert data user
// 	err := tx.WithContext(ctx).Create(&user).Error

// 	// mengecek error setelah melakukan insert
// 	if err != nil {
// 		return user, err
// 	}

// 	return user, nil
// }

// func (repository *UserRepositoryImpl) SignIn(ctx context.Context, tx *gorm.DB, email string) (domain.User, error) {
// 	// menyiapkan data user
// 	var user domain.User

// 	result := tx.WithContext(ctx).
// 		Where("email = ?", email).
// 		First(&user)

// 	if result.Error != nil {
// 		return user, result.Error
// 	}

// 	return user, nil
// }
