package cases

import (
	"AuthService/internal/logic"
	"AuthService/internal/utils/logger"
	"context"
	"time"
)

type DisableOtpAuthenticationForUserArgs struct{
	Ctx context.Context
	UserId string
}

type DisableOtpAuthenticationForUserReturned struct{
	Err error
}

func (c *CasesProvider)DisableOtpAuthenticationForUser(args DisableOtpAuthenticationForUserArgs) (returned DisableOtpAuthenticationForUserReturned){

	logicArgs := logic.DisableOtpAuthenticationForUserArgs{
		Ctx: args.Ctx,
		UserId: args.UserId,
	}
	logicReturned := c.logic.DisableOtpAuthenticationForUser(logicArgs)
	if logicReturned.Err == logic.ErrUserNotExists || logicReturned.Err == logic.ErrOtpAlreadyDisabled{
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

	return
}