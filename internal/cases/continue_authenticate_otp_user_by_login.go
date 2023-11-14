package cases

import (
	"auth/internal/logic"
	"auth/internal/utils/logger"
	"context"
	"time"
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

func (c *CasesProvider) ContinueAuthenticateOtpUserByLogin(args ContinueAuthenticateOtpUserByLoginArgs) (returned ContinueAuthenticateOtpUserByLoginReturned) {
	select {
	case <-args.Ctx.Done():
		{
			returned.Err = ErrServiceNotAvaliable
			return
		}
	default:
		{
			var logicArgs logic.ContinueAuthenticateOtpUserByLoginArgs = logic.ContinueAuthenticateOtpUserByLoginArgs{
				Ctx:               args.Ctx,
				IntermediateToken: args.IntermediateToken,
				OtpCode:           args.OtpCode,
			}

			logicReturned := c.logic.ContinueAuthenticateOtpUserByLogin(logicArgs)

			if logicReturned.Err == logic.ErrInvalidOtp || logicReturned.Err == logic.ErrExpiredIntermediateToken || logicReturned.Err == logic.ErrInvalidIntermediateToken {
				returned.Err = logicReturned.Err
				return
			} else if logicReturned.Err != nil {

				c.logger.Log(logger.LogMsg{
					Time:     time.Now(),
					LogLevel: logger.LogLevelError,
					Msg:      "Internal continue authenticate otp user by login error: " + logicReturned.Err.Error(),
				})

				returned.Err = ErrServiceInternal
				return
			}

			returned.AuthInfo.AccessToken = logicReturned.AuthInfo.AccessToken
			returned.AuthInfo.RefreshToken = logicReturned.AuthInfo.RefreshToken
			return
		}
	}
}
