package cases

import (
	"AuthService/internal/logic"
	"AuthService/internal/logic/models"
	"AuthService/internal/utils/logger"
	"context"
	"time"
)

type AuthorizeUserArgs struct {
	Ctx            context.Context
	AccessToken    string
	RequiredRoleId models.RoleId
}

type AuthorizeUserReturned struct {
	UserId string
	Err    error
}

func (c *CasesProvider) AuthorizeUser(args AuthorizeUserArgs) (returned AuthorizeUserReturned) {
	var logicArgs logic.AuthorizeUserArgs = logic.AuthorizeUserArgs{
		Ctx:         args.Ctx,
		AccessToken: args.AccessToken,
	}

	logicReturned := c.logic.AuthorizeUser(logicArgs)

	if logicReturned.Err == logic.ErrUserNotExists || logicReturned.Err == logic.ErrExpiredAccessToken || logicReturned.Err == logic.ErrInvalidAccessToken || logicReturned.Err == logic.ErrRoleHasNoAccess {
		returned.Err = logicReturned.Err
		return
	} else if logicReturned.Err != nil {
		// ToDo: add tracing by token
		c.logger.Log(logger.LogMsg{
			Time:     time.Now(),
			LogLevel: logger.LogLevelError,
			Msg:      "Internal authorize user error: " + logicReturned.Err.Error(),
		})
		returned.Err = ErrServiceInternal
		return
	}
	returned.UserId = logicReturned.UserId
	return
}