package errs

import "auth/internal/utils/errsutil"

type ErrUserAlreadyExists struct{}

func (e ErrUserAlreadyExists) Error() string {
	return "user already exists"
}

func (e ErrUserAlreadyExists) ErrCode() int {
	return errsutil.ErrUserAlreadyExistsCode
}
