package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ProfileUser struct {
	BaseModel // menggunakan id, created_at, dan updated_at sesuai dengan base models

	Name           string `gorm:"column:name;size:100;not null"`
	ProfilePicture string `gorm:"column:profile_picture;size:255"`
	Phone          string `gorm:"column:phone;size:15"`
	Gender         string `gorm:"column:gender;type:enum('L', 'P')"`
	Bio            string `gorm:"column:bio;type:text"`

	// foreign key ke model lain
	IdUser string `gorm:"column:id_user;size:36;not null"`

	// relasi ke tabel parent
	User User `gorm:"foreignKey:id_user;references:id"`
}

// membuat method baru untuk mengganti nama tabel (alias)
func (u *ProfileUser) TableName() string {
	return "profile_users"
}

// menambahkan hook sebelum menambahkan data profile_user, maka generate uuid untuk id profile_user dilakukan disini
func (u *ProfileUser) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.NewString()
	return
}