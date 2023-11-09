package errs

import "auth/internal/utils/errsutil"

type ErrInvalidIntermediateToken struct {
}

func (e ErrInvalidIntermediateToken) Error() string {
	return "invalid intermediate token"
}

func (e ErrInvalidIntermediateToken) ErrCode() errsutil.AuthErrCode {
	return errsutil.ErrInvalidIntermediateTokenCode
}
