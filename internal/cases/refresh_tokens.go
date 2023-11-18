package cases

import (
	"AuthService/internal/logic"
	"AuthService/internal/utils/logger"
	"context"
	"time"
)

type RefreshTokensArgs struct {
	Ctx          context.Context
	AccessToken  string
	RefreshToken string
}

type RefreshTokensReturned struct {
	AccessToken  string
	RefreshToken string
	Err          error
}

func (c *CasesProvider) RefreshTokens(args RefreshTokensArgs) (returned RefreshTokensReturned) {

	var logicArgs logic.RefreshTokensArgs = logic.RefreshTokensArgs{
		Ctx:          args.Ctx,
		AccessToken:  args.AccessToken,
		RefreshToken: args.RefreshToken,
	}

	logicReturned := c.logic.RefreshTokens(logicArgs)

	if logicReturned.Err == logic.ErrExpiredRefreshToken || logicReturned.Err == logic.ErrInvalidRefreshToken || logicReturned.Err == logic.ErrUserNotExists || logicReturned.Err == logic.ErrInvalidAccessToken {
		returned.Err = logicReturned.Err
		return
	} else if logicReturned.Err != nil {
		c.logger.Log(logger.LogMsg{
			Time:     time.Now(),
			LogLevel: logger.LogLevelError,
			Msg:      "Internal authorize user error: " + logicReturned.Err.Error(),
		})
		returned.Err = ErrServiceInternal
		return
	}
	returned.AccessToken = logicReturned.AccessToken
	returned.RefreshToken = logicReturned.RefreshToken
	return
}