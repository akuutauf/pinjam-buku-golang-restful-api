package web

// membuat standar web response
type WebResponse struct {
	Code   int    `json:"code"`
	Status string `json:"status"`

	// gunakan interface{}, agar datanya bisa di isi tipe data apapun
	Data interface{} `json:"data"`
}
