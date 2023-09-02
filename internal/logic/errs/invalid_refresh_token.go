package errs

import "auth/internal/utils/errsutil"

type ErrInvalidRefreshToken struct{}

func (e ErrInvalidRefreshToken) Error() string {
	return "invalid refresh token"
}

func (e ErrInvalidRefreshToken) ErrCode() int {
	return errsutil.ErrInvalidRefreshTokenCode
}
