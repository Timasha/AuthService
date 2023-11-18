package logic

import (
	"AuthService/internal/utils/errsutil"
	"context"
)

type DisableOtpAuthenticationForUserArgs struct{
	Ctx context.Context
	UserId string
}

type DisableOtpAuthenticationForUserReturned struct{
	Err error
}

func (l *LogicProvider) DisableOtpAuthenticationForUser(args DisableOtpAuthenticationForUserArgs)(returned DisableOtpAuthenticationForUserReturned){
	gettedUser, getUserErr := l.userStorage.GetUserByUserId(args.Ctx,args.UserId)
	if (getUserErr != nil){
		returned.Err = getUserErr
		return
	}

	if !gettedUser.OtpEnabled{
		returned.Err = ErrOtpAlreadyDisabled
		return
	}

	gettedUser.OtpEnabled = false
	gettedUser.OtpKey = ""
	updateErr := l.userStorage.UpdateUserByLogin(args.Ctx,gettedUser.Login,gettedUser)

	returned.Err = updateErr

	return
}

var ErrOtpAlreadyDisabled errsutil.AuthErr = errsutil.AuthErr{
	Msg:     "otp is already disabled",
	ErrCode: errsutil.ErrOtpAlreadyDisabledCode,
}
