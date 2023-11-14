package cases

import (
	"auth/internal/logic"
	"auth/internal/logic/models"
	"auth/internal/utils/logger"
	"context"
	"time"
)

type AddRoleArgs struct {
	Ctx  context.Context
	Role models.Role
}

type AddRoleReturned struct {
	Err error
}

func (c *CasesProvider) AddRole(args AddRoleArgs) (returned AddRoleReturned) {
	select {
	case <-args.Ctx.Done():
		{
			returned.Err = ErrServiceNotAvaliable
			return
		}
	default:
		{
			var logicArgs logic.AddRoleArgs = logic.AddRoleArgs{
				Ctx:  args.Ctx,
				Role: args.Role,
			}
			logicReturned := c.logic.AddRole(logicArgs)
			if logicReturned.Err == logic.ErrRoleAlreadyExists {
				returned.Err = logicReturned.Err
				return
			} else if logicReturned.Err != nil {
				c.logger.Log(logger.LogMsg{
					Time:     time.Now(),
					LogLevel: logger.LogLevelError,
					Msg:      "Internal add role error: " + logicReturned.Err.Error(),
				})
				returned.Err = ErrServiceInternal
				return
			}

			return
		}
	}
}
