package domain

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BaseModel struct {
	ID string `gorm:"primary_key;column:id;size:36;not null;<-:create"`
}

// menambahkan hook sebelum menambahkan data pada tabel yang mengadaptasi base model,-
// maka generate uuid untuk data di tabel tersebut akan otomatis dilakukan disini
func (b *BaseModel) BeforeCreate(tx *gorm.DB) error {
	b.ID = uuid.NewString()
	return nil
}
