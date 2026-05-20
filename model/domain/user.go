package domain

type User struct {
	BaseModel // template id yang menggunakan uuid

	Username string `gorm:"column:username;size:20;not null;unique"`
	Email    string `gorm:"column:email;size:60;not null;unique"`
	Password string `gorm:"column:password;size:255;not null"`

	BaseTime // template created_at, dan updated_at diletakkan di akhir kolom

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
