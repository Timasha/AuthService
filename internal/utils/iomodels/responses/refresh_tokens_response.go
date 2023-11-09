package responses

import "auth/internal/utils/errsutil"

type RefreshTokensResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`

	Err     string               `json:"error"`
	ErrCode errsutil.AuthErrCode `json:"errorCode"`
}
