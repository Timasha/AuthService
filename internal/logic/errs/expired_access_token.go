package errs

import "auth/internal/utils/errsutil"

type ErrExpiredAccessToken struct {
}

func (e ErrExpiredAccessToken) Error() string {
	return "access token is expired"
}

func (e ErrExpiredAccessToken) ErrCode()int{
	return errsutil.ErrExpiredAccessTokenCode
}