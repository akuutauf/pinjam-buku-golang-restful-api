package domain

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BaseModel struct {
	ID        string `gorm:"primary_key;column:id;size:36;not null;<-:create"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime;not null;<-:create"` 
	UpdatedAt time.Time `gorm:"column:updated_at;autoCreateTime;autoUpdateTime;not null"`
}

func (b *BaseModel) BeforeCreate(tx *gorm.DB) error {
	b.ID = uuid.NewString()
	return nil
}