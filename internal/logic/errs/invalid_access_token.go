package errs

import "auth/internal/utils/errsutil"

type ErrInvalidAccessToken struct{}

func (e ErrInvalidAccessToken) Error() string {
	return "invalid access token"
}

func (e ErrInvalidAccessToken) ErrCode() errsutil.AuthErrCode {
	return errsutil.ErrInvalidAccessTokenCode
}
