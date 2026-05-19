package domain

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	BaseModel // gunakan id, createdAt, dan updatedAt dari base model

	Username  string    `gorm:"column:username;size:20;not null;unique"`
	Email     string    `gorm:"column:email;size:60;not null;unique"`
	Password  string    `gorm:"column:password;size:255;not null"`

	// implementasi one to one (has one)
	ProfileUser *ProfileUser `gorm:"foreignKey:id_user;references:ID"`

	// relasi ke tabel child (has many)
	Book []Book `gorm:"foreignKey:id_user;references:id"`
	Loan []Loan `gorm:"foreignKey:id_user;references:id"`
}

// membuat method baru untuk mengganti nama tabel (alias)
func (u *User) TableName() string {
	return "users"
}

// menambahkan hook sebelum menambahkan data user, maka generate uuid untuk id user dilakukan disini
func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.NewString()
	return
}