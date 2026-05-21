package user

import (
	"pinjam-buku/model/domain"
	"pinjam-buku/model/web/user"
)

// membuat function untuk mengkonversi data menjadi tipe User Response
func ToUserResponse(data domain.User) user.UserResponse {
	return user.UserResponse{
		ID:       data.ID,
		Username: data.Username,
		Email:    data.Email,
	}
}
