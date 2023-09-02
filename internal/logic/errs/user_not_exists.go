package errs

import "auth/internal/utils/errsutil"

type ErrUserNotExists struct{}

func (e ErrUserNotExists) Error() string {
	return "user not exists"
}

func (e ErrUserNotExists) ErrCode() int {
	return errsutil.ErrUserNotExistsCode
}
