package user

// representasi request create, yang mana hanya mengirimkan atribut yang diperlukan saja
// juga untuk id akan otomatis di generate karena menggunakan UUID dadi BaseModel

type UserAuthRequest struct {
	Username string `validate:"required,min=1,max=20" json:"username"`
	Password string `validate:"required,min=1,max=20" json:"password"`
}
