package domain

import (
	"time"
)

type Loan struct {
	BaseModel // template id yang menggunakan uuid

	BorrowerName  string    `gorm:"column:borrow_name;size:100;not null"`
	BorrowerPhone string    `gorm:"column:borrow_phone;size:15;not null"`
	BorrowerDate  time.Time `gorm:"column:borrow_date;type:datetime;not null"`
	DueDate       time.Time `gorm:"column:due_date;type:date;not null"`
	ReturnedAt    time.Time `gorm:"column:returned_at;type:datetime"`
	Status        string    `gorm:"column:status;type:enum('Dipinjam', 'Terlambat', 'Selesai');not null"`
	Note          string    `gorm:"column:note;type:text"`

	// foreign key ke model lain
	IdUser string `gorm:"column:id_user;size:36;not null"`
	IdBook string `gorm:"column:id_book;size:36;not null"`

	BaseTime // template created_at, dan updated_at diletakkan di akhir kolom

	// relasi ke tabel parent
	User User `gorm:"foreignKey:id_user;references:id"`
	Book Book `gorm:"foreignKey:id_book;references:id"`
}

// membuat method baru untuk mengganti nama tabel (alias)
func (u *Loan) TableName() string {
	return "loans"
}
