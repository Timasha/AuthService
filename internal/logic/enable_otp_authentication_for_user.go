package logic

import (
	"AuthService/internal/utils/errsutil"
	"context"
)

type EnableOtpAuthenticationForUserArgs struct{
	Ctx context.Context
	UserId string
}

type EnableOtpAuthenticationForUserReturned struct{
	OtpUrl string
	OtpKey string
	Err error
}

func (l *LogicProvider) EnableOtpAuthenticationForUser(args EnableOtpAuthenticationForUserArgs) (returned EnableOtpAuthenticationForUserReturned){
	gettedUser, getUserErr := l.userStorage.GetUserByUserId(args.Ctx,args.UserId)
	if (getUserErr != nil){
		returned.Err = getUserErr
		return
	}

	if (gettedUser.OtpEnabled){
		returned.Err = ErrOtpAlreadyEnabled
		return
	}

	otpKey, otpUrl, err := l.otpGenerator.GenerateKeys(gettedUser.Login)

	if err != nil{
		returned.Err = err
		return
	}

	gettedUser.OtpEnabled = true
	gettedUser.OtpKey = otpKey
	
	updateErr := l.userStorage.UpdateUserByLogin(args.Ctx,gettedUser.Login,gettedUser)
	if updateErr != nil{
		returned.Err = updateErr
		return
	}

	returned.OtpKey = otpKey
	returned.OtpUrl = otpUrl

	return
}

var ErrOtpAlreadyEnabled errsutil.AuthErr = errsutil.AuthErr{
	Msg:     "otp is already enabled",
	ErrCode: errsutil.ErrOtpAlreadyEnabledCode,
}
