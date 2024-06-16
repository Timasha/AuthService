package usecase

import (
	"context"

	"github.com/Timasha/AuthService/pkg/errlist"
)

type ContinueAuthenticateOtpUserByLoginRequest struct {
	IntermediateToken string
	OtpCode           string
}

type ContinueAuthenticateOtpUserByLoginResponse struct {
	TokenPair TokenPair
}

func (c *Provider) ContinueAuthenticateOtpUserByLogin(
	ctx context.Context,
	req ContinueAuthenticateOtpUserByLoginRequest,
) (resp ContinueAuthenticateOtpUserByLoginResponse, err error) {
	login, err := c.tokensProvider.ValidIntermediateToken(req.IntermediateToken)
	if err != nil {
		return resp, err
	}

	user, err := c.userStorage.GetUserByLogin(ctx, login)
	if err != nil {
		return resp, err
	}

	if !c.otpGenerator.ValidOtp(req.OtpCode, user.OtpKey) {
		return resp, errlist.ErrInvalidOtp
	}

	tokenPair, err := c.tokensProvider.CreateTokens(login)

	return ContinueAuthenticateOtpUserByLoginResponse{
		TokenPair: tokenPair,
	}, err
}
