package domain

// membuat struct category untuk merepresentasikan sebuah table
type Category struct {
	BaseModel // template id yang menggunakan uuid

	Name string `gorm:"column:name;size:20;not null;unique"`

	BaseTime // template created_at, dan updated_at diletakkan di akhir kolom

	// relasi ke tabel child (has many)
	Book []Book `gorm:"foreignKey:id_category;references:id"`
}

// membuat method baru untuk mengganti nama tabel (alias)
func (u *Category) TableName() string {
	return "categories"
}
