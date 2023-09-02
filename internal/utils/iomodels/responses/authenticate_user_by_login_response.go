package responses

type AuthenticateUserByLoginResponse struct {
	Err     string `json:"error"`
	ErrCode int    `json:"errorCode"`

	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}
