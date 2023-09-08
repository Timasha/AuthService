package responses

type RefreshTokensResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`

	Err     string `json:"error"`
	ErrCode int    `json:"errorCode"`
}
