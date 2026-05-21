package user

import (
	"context"
	"pinjam-buku/model/domain"

	"gorm.io/gorm"
)

// membuat repository khusus untuk user

// membuat kontrak dalam bentuk interface (best practice)
type UserRepository interface {
	// membuat semua repository function untuk CRUD
	SignUp(ctx context.Context, tx *gorm.DB, user domain.User) (domain.User, error)
	SignIn(ctx context.Context, tx *gorm.DB, user domain.User) (domain.User, error)
	UpdateUser(ctx context.Context, tx *gorm.DB, user domain.User) (domain.User, error)
	SignOut(ctx context.Context, tx *gorm.DB, user domain.User) error
	GetUserCurrent(ctx context.Context, tx *gorm.DB, userId string) (domain.User, error)
}
