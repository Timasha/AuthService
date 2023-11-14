package logic

import (
	"context"
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

func (l *LogicProvider) RefreshTokens(args RefreshTokensArgs) (returned RefreshTokensReturned) {

	login, accessValidErr := l.tokensProvider.ValidAccessToken(args.AccessToken)

	if accessValidErr != nil && accessValidErr != ErrExpiredAccessToken {
		returned.Err = accessValidErr
		return
	}

	_, getErr := l.userStorage.GetUserByLogin(args.Ctx, login)

	if getErr != nil {
		returned.Err = getErr
		return
	}

	refreshValidErr := l.tokensProvider.ValidRefreshToken(args.RefreshToken, args.AccessToken)

	if refreshValidErr != nil {
		returned.Err = refreshValidErr
		return
	}

	returned.AccessToken, returned.RefreshToken, returned.Err = l.tokensProvider.CreateTokens(login)

	return
}
