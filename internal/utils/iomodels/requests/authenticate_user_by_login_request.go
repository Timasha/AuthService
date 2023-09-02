package requests

type AuthenticateUserByLoginRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}
