package errs

import "auth/internal/utils/errsutil"

type ErrInvalidLoginOrPassword struct{}

func (e ErrInvalidLoginOrPassword) Error() string {
	return "invalid login or password"
}

func (e ErrInvalidLoginOrPassword) ErrCode() errsutil.AuthErrCode {
	return errsutil.ErrInvalidLoginOrPasswordCode
}
