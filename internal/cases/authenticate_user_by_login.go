package cases

import (
	"AuthService/internal/logic"
	"AuthService/internal/utils/logger"
	"context"
	"time"
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

func (c *CasesProvider) AuthenticateUserByLogin(args AuthenticateUserByLoginArgs) (returned AuthenticateUserByLoginReturned) {
	select {
	case <-args.Ctx.Done():
		{
			returned.Err = ErrServiceNotAvaliable
			return
		}
	default:
		{
			if len(args.Login) < c.config.GetMinLoginLen() || len(args.Password) < c.config.GetMinPasswordLen() {
				returned.Err = ErrTooShortLoginOrPassword
				return
			}

			var logicArgs logic.AuthenticateUserByLoginArgs = logic.AuthenticateUserByLoginArgs{
				Ctx:      args.Ctx,
				Login:    args.Login,
				Password: args.Password,
			}

			logicReturned := c.logic.AuthenticateUserByLogin(logicArgs)

			if logicReturned.Err == logic.ErrUserNotExists || logicReturned.Err == logic.ErrInvalidPassword {
				returned.Err = ErrInvalidLoginOrPassword
				return
			} else if logicReturned.Err != nil {
				// ToDo: add tracing by token
				c.logger.Log(logger.LogMsg{
					Time:     time.Now(),
					LogLevel: logger.LogLevelError,
					Msg:      "Internal authenticate user by login error: " + logicReturned.Err.Error(),
				})
				returned.Err = ErrServiceInternal
				return
			}
			if logicReturned.OtpEnabled {
				returned.IntermediateToken = logicReturned.IntermediateToken
				returned.OtpEnabled = logicReturned.OtpEnabled
				return
			}
			returned.AuthInfo.AccessToken = logicReturned.AuthInfo.AccessToken
			returned.AuthInfo.RefreshToken = logicReturned.AuthInfo.RefreshToken
			return
		}
	}
}
