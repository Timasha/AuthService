package cases

import (
	"AuthService/internal/logic"
	"AuthService/internal/utils/logger"
	"context"
	"time"
)

type EnableOtpAuthenticationForUserArgs struct{
	Ctx context.Context
	UserId string
}

type EnableOtpAuthenticationForUserReturned struct{
	OtpKey string
	OtpUrl string
	Err error
}

func (c *CasesProvider)EnableOtpAuthenticationForUser(args EnableOtpAuthenticationForUserArgs) (returned EnableOtpAuthenticationForUserReturned){

	logicArgs := logic.EnableOtpAuthenticationForUserArgs{
		Ctx: args.Ctx,
		UserId: args.UserId,
	}
	logicReturned := c.logic.EnableOtpAuthenticationForUser(logicArgs)
	if logicReturned.Err == logic.ErrUserNotExists || logicReturned.Err == logic.ErrOtpAlreadyEnabled{
		returned.Err = logicReturned.Err
		return
	}else if logicReturned.Err != nil{
		c.logger.Log(logger.LogMsg{
			Time:     time.Now(),
			LogLevel: logger.LogLevelError,
			Msg:      "Internal enable otp authentication for user error: " + logicReturned.Err.Error(),
		})

		returned.Err = ErrServiceInternal
		return
	}

	returned.OtpKey = logicReturned.OtpKey
	returned.OtpUrl = logicReturned.OtpUrl

	return
}