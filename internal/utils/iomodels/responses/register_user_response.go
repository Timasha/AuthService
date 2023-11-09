package responses

import "auth/internal/utils/errsutil"

type RegisterUserResponses struct {
	Err     string               `json:"error"`
	ErrCode errsutil.AuthErrCode `json:"errorCode"`
}
