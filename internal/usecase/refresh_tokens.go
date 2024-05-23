package usecase

import (
	"context"
	"errors"

	"github.com/Timasha/AuthService/internal/entities"
	"github.com/Timasha/AuthService/pkg/errlist"
)

type RefreshTokensRequest struct {
	TokenPair entities.TokenPair
}

type RefreshTokensResponse struct {
	TokenPair entities.TokenPair
}

func (c *Provider) RefreshTokens(
	ctx context.Context,
	req RefreshTokensRequest,
) (resp RefreshTokensResponse, err error) {
	login, err := c.tokensProvider.ValidAccessToken(req.TokenPair.AccessToken)
	if err != nil && !errors.Is(err, errlist.ErrExpiredAccessToken) {
		return resp, err
	}

	err = c.tokensProvider.ValidRefreshToken(req.TokenPair)
	if err != nil {
		return resp, err
	}

	_, err = c.userStorage.GetUserByLogin(ctx, login)
	if err != nil {
		return resp, err
	}

	tokenPair, err := c.tokensProvider.CreateTokens(login)

	return RefreshTokensResponse{
		TokenPair: tokenPair,
	}, err
}
