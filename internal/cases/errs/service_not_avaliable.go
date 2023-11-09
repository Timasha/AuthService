package errs

import "auth/internal/utils/errsutil"

type ErrServiceNotAvaliable struct{}

func (e ErrServiceNotAvaliable) Error() string {
	return "service not avaliable"
}

func (e ErrServiceNotAvaliable) ErrCode() errsutil.AuthErrCode {
	return errsutil.ErrServiceNotAvaliableCode
}
