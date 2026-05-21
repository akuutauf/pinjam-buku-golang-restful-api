package user

// membuat representasi response untuk endpoint user
type UserAuthResponse struct {
	Token string `json:"token"`
}
