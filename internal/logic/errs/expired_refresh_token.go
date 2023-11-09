package errs

import "auth/internal/utils/errsutil"

type ErrExpiredRefreshToken struct{}

func (e ErrExpiredRefreshToken) Error() string {
	return "refresh token is expired"
}

func (e ErrExpiredRefreshToken) ErrCode() errsutil.AuthErrCode {
	return errsutil.ErrExpiredRefreshTokenCode
}
