package domain

import (
	"time"
)

type BaseTime struct {
	// khusus untuk kolom created_at dan updated_at
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime;not null;<-:create"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoCreateTime;autoUpdateTime;not null"`
}
