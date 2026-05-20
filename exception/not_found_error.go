package exception

// membuat struct untuk menangani not found error
type NotFoundError struct {
	// karena error merupakan kontrak yang bersifat interface,
	// maka kita buat atribute error dengan tipe data apapun (string)
	Message string
}

// membuat function untuk membuat pesan not found error
func NewNotFoundError(message string) NotFoundError {
	// mengembalikan error yang diterima oleh server
	return NotFoundError{
		Message: message,
	}
}

func (e NotFoundError) Error() string {
	return e.Message
}
