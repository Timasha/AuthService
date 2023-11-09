package errs

import "auth/internal/utils/errsutil"

type ErrInvalidOtp struct{}

func (e ErrInvalidOtp) Error() string {
	return "invalid otp code"
}

func (e ErrInvalidOtp) ErrCode() errsutil.AuthErrCode {
	return errsutil.ErrInvalidOtpCode
}
