package logic

import (
	"AuthService/internal/utils/errsutil"
	"context"
)

type AuthenticateUserByLoginArgs struct {
	Ctx      context.Context
	Login    string
	Password string
}

type AuthenticateUserByLoginReturned struct {
	OtpEnabled bool

	IntermediateToken string

	AuthInfo struct {
		AccessToken  string
		RefreshToken string
	}
	Err error
}

func (l *LogicProvider) AuthenticateUserByLogin(args AuthenticateUserByLoginArgs) (returned AuthenticateUserByLoginReturned) {
	user, getErr := l.userStorage.GetUserByLogin(args.Ctx, args.Login)

	if getErr != nil {
		returned.Err = getErr
		return
	}

	if !l.passwordHasher.Compare(args.Password, user.Password) {
		returned.Err = ErrInvalidPassword
		return
	}

	if user.OtpEnabled {
		returned.OtpEnabled = true

		returned.IntermediateToken, returned.Err = l.tokensProvider.CreateIntermediateToken(args.Login)

		return
	}

	returned.AuthInfo.AccessToken, returned.AuthInfo.RefreshToken, returned.Err = l.tokensProvider.CreateTokens(args.Login)

	return
}

//errors

var ErrInvalidPassword errsutil.AuthErr = errsutil.AuthErr{
	Msg:     "invalid password",
	ErrCode: errsutil.ErrInvalidPasswordCode,
}
