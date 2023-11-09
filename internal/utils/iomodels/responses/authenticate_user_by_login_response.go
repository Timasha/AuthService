package responses

import "auth/internal/utils/errsutil"

type AuthenticateUserByLoginResponse struct {
	Err     string               `json:"error"`
	ErrCode errsutil.AuthErrCode `json:"errorCode"`

	OtpEnabled bool `json:"otpEnabled"`

	IntermediateToken string `json:"IntermediateToken"`

	AuthInfo struct {
		AccessToken  string `json:"accessToken"`
		RefreshToken string `json:"refreshtoken"`
	} `json:"authInfo"`
}
