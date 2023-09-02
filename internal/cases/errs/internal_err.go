package errs

import "auth/internal/utils/errsutil"

type ErrServiceInternal struct{}

func (e ErrServiceInternal) Error() string {
	return "internal service error"
}

func (e ErrServiceInternal) ErrCode() int {
	return errsutil.ErrServiceInternalCode
}
