package requests

type AuthorizeUserRequest struct {
	AccessToken string `json:"accessToken"`
	Login       string `json:"login"`
}
