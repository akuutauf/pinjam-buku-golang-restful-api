package helper

// membuat helper error agar tidak duplikat (agar reusable)
func PanicIfError(err error) {
	// mengecek error
	if err != nil {
		panic(err)
	}
}
