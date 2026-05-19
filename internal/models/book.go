package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Book struct {
	BaseModel // gunakan id, createdAt, dan updatedAt dari base model

	Title  string   `gorm:"column:title;size:200;not null"`
	Author string	`gorm:"column:author;size:50;not null"`
	Stock uint8	`gorm:"column:stock;not null"`
	IsAvailable bool	`gorm:"column:is_available;size:1;not null"`
	Cover string	`gorm:"column:cover;size:255"`
	QrContent string	`gorm:"column:qr_content;size:255"`
	Status string	`gorm:"column:status;type:enum('Normal', 'Hilang', 'Rusak');not null"`
	Publisher string	`gorm:"column:publisher;size:100;not null"`
	PublishYear int16	`gorm:"column:publish_year;type:year;not null"`
	Description string	`gorm:"column:description;type:text"`
	
	// foreign key ke model lain
	IdUser    string `gorm:"column:id_user;size:36;not null"`

	// relasi ke tabel parent
	User User `gorm:"foreignKey:id_user;references:id"`

	// relasi ke tabel child (has many)
	Loan []Loan `gorm:"foreignKey:id_book;references:id"`
}

// membuat method baru untuk mengganti nama tabel (alias)
func (u *Book) TableName() string {
	return "books"
}

// menambahkan hook sebelum menambahkan data book, maka generate uuid untuk id book dilakukan disini
func (u *Book) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.NewString()
	return
}