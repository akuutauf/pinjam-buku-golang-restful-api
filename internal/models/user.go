package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID        string    `gorm:"primary_key;column:id;size:36;not null;<-:create"`
	Username  string    `gorm:"column:username;size:20;not null;unique"`
	Email     string    `gorm:"column:email;size:60;not null;unique"`
	Password  string    `gorm:"column:password;size:255;not null"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime;<-:create"` 
	UpdatedAt time.Time `gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
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