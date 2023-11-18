package logic

import (
	"AuthService/internal/utils/errsutil"
	"context"
)

type ContinueAuthenticateOtpUserByLoginArgs struct {
	Ctx context.Context

	IntermediateToken string
	OtpCode           string
}

type ContinueAuthenticateOtpUserByLoginReturned struct {
	AuthInfo struct {
		AccessToken  string
		RefreshToken string
	}
	Err error
}

func (l *LogicProvider) ContinueAuthenticateOtpUserByLogin(args ContinueAuthenticateOtpUserByLoginArgs) (returned ContinueAuthenticateOtpUserByLoginReturned) {

	login, validErr := l.tokensProvider.ValidIntermediateToken(args.IntermediateToken)

	if validErr != nil {
		returned.Err = validErr
		return
	}

	user, getErr := l.userStorage.GetUserByLogin(args.Ctx, login)

	if getErr != nil {
		returned.Err = getErr
		return
	}

	if !l.otpGenerator.ValidOtp(args.OtpCode, user.OtpKey) {
		returned.Err = ErrInvalidOtp
		return
	}

	returned.AuthInfo.AccessToken, returned.AuthInfo.RefreshToken, returned.Err = l.tokensProvider.CreateTokens(login)

	return
}

var ErrInvalidOtp errsutil.AuthErr = errsutil.AuthErr{
	Msg:     "invalid otp code",
	ErrCode: errsutil.ErrInvalidOtpCode,
}
