package errs

import "auth/internal/utils/errsutil"

type ErrInvalidPassword struct{}

func (e ErrInvalidPassword) Error() string {
	return "invalid password"
}

func (e ErrInvalidPassword) ErrCode() int {
	return errsutil.ErrInvalidPasswordCode
}
