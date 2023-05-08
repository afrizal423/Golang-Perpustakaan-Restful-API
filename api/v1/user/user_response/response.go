package user_response

type LoginResponse struct {
	Username string `json:"username"`
	Token    string `json:"token"`
}

func NewLoginResponse(username string, token string) *LoginResponse {
	var loginResponse LoginResponse
	loginResponse.Username = username
	loginResponse.Token = token
	return &loginResponse
}
