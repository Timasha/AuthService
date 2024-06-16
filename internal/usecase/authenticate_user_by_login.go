package usecase

import (
	"context"

	"github.com/Timasha/AuthService/internal/entities"
	"github.com/Timasha/AuthService/pkg/errlist"
)

type AuthenticateUserByLoginRequest struct {
	Login    string
	Password string
}

func (a AuthenticateUserByLoginRequest) ToUser() *entities.User {
	return &entities.User{
		Login:    a.Login,
		Password: a.Password,
	}
}

type AuthenticateUserByLoginResponse struct {
	OtpEnabled bool

	IntermediateToken *string

	TokenPair *TokenPair
}

func (c *Provider) AuthenticateUserByLogin(
	ctx context.Context,
	req AuthenticateUserByLoginRequest,
) (resp AuthenticateUserByLoginResponse, err error) {
	req.Password, err = c.passwordHasher.Hash(req.Password)
	if err != nil {
		return resp, err
	}

	user, err := c.userStorage.GetUserByLogin(ctx, req.Login)
	if err != nil {
		return resp, err
	}

	if !user.Check(req.ToUser()) {
		return resp, errlist.ErrInvalidLoginOrPassword
	}

	if user.OtpEnabled {
		intermediateToken, err := c.tokensProvider.CreateIntermediateToken(req.Login)

		resp = AuthenticateUserByLoginResponse{
			OtpEnabled:        true,
			IntermediateToken: &intermediateToken,
		}

		return resp, err
	}

	tokensPair, err := c.tokensProvider.CreateTokens(req.Login)
	if err != nil {
		return resp, err
	}

	resp = AuthenticateUserByLoginResponse{
		OtpEnabled: false,

		TokenPair: &tokensPair,
	}

	return resp, nil
}
