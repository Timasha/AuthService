package responses

import "auth/internal/utils/errsutil"

type AuthorizeUserResponses struct {
	Uuid string `json:"uuid"`

	Err     string               `json:"error"`
	ErrCode errsutil.AuthErrCode `json:"errorCode"`
}
