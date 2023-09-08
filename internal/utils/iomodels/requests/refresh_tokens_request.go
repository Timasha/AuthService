package requests

type RefreshTokenRequest struct {
	RefreshToken string `json:"refreshToken"`
	AccessToken  string `json:"accessToken"`
	Login        string `json:"login"`
}
