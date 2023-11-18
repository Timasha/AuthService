package cases

import (
	"AuthService/internal/logic"
	"AuthService/internal/logic/models"
	"AuthService/internal/utils/logger"
	"context"
	"time"
)

type RegisterUserArgs struct {
	Ctx  context.Context
	User models.User
}

type RegisterUserReturned struct {
	Err error
}

func (c *CasesProvider) RegisterUser(args RegisterUserArgs) (returned RegisterUserReturned) {

	if len(args.User.Login) < c.config.GetMinLoginLen() || len(args.User.Password) < c.config.GetMinPasswordLen() {
		returned.Err = ErrTooShortLoginOrPassword
		return
	}
	args.User.Role.RoleId = c.config.GetDefaultUserRoleId()

	var logicArgs logic.RegisterUserArgs = logic.RegisterUserArgs{
		Ctx:  args.Ctx,
		User: args.User,
	}
	logicReturned := c.logic.RegisterUser(logicArgs)

	if logicReturned.Err == logic.ErrUserAlreadyExists {
		returned.Err = logicReturned.Err
		return
	} else if logicReturned.Err != nil {
		// ToDo: add tracing by token
		c.logger.Log(logger.LogMsg{
			Time:     time.Now(),
			LogLevel: logger.LogLevelError,
			Msg:      "Internal register user error: " + logicReturned.Err.Error(),
		})
		returned.Err = ErrServiceInternal
		return
	}
	return
}
