package user

// representasi request update, yang diperlukan saja
// meskipun id akan otomatis di generate karena UUID dari BaseModel
// namun update membutuhkan id sebagai parameter
// proses validasi dibutuhkan pada saat update

type UserUpdateRequest struct {
	// menjabarkan aturan validasi untuk request body
	Id       string `validate:"required" json:"id"`
	Username string `validate:"required,min=1,max=20" json:"username"`
	Email    string `validate:"required,min=1,max=60" json:"email"`
	Password string `validate:"required,min=1,max=20" json:"password"`
}
