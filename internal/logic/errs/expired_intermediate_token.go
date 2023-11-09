package errs

import "auth/internal/utils/errsutil"

type ErrExpiredIntermediateToken struct {
}

func (e ErrExpiredIntermediateToken) Error() string {
	return "intermediate token is expired"
}

func (e ErrExpiredIntermediateToken) ErrCode() errsutil.AuthErrCode {
	return errsutil.ErrExpiredIntermediateTokenCode
}
