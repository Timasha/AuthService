package errs

import "auth/internal/utils/errsutil"

type ErrTooShortLoginOrPassword struct{}

func (e ErrTooShortLoginOrPassword) Error() string {
	return "too short login or password"
}

func (e ErrTooShortLoginOrPassword) ErrCode() int {
	return errsutil.ErrTooShortLoginOrPasswordCode
}
