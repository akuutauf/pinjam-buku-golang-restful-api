package domain

type ProfileUser struct {
	BaseModel // template id yang menggunakan uuid

	Name           string `gorm:"column:name;size:100;not null"`
	ProfilePicture string `gorm:"column:profile_picture;size:255"`
	Phone          string `gorm:"column:phone;size:15"`
	Gender         string `gorm:"column:gender;type:enum('L', 'P')"`
	Bio            string `gorm:"column:bio;type:text"`

	// foreign key ke model lain
	IdUser string `gorm:"column:id_user;size:36;not null"`

	BaseTime // template created_at, dan updated_at diletakkan di akhir kolom

	// relasi ke tabel parent
	User User `gorm:"foreignKey:id_user;references:id"`
}

// membuat method baru untuk mengganti nama tabel (alias)
func (u *ProfileUser) TableName() string {
	return "profile_users"
}
