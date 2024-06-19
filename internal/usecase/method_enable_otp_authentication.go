package usecase

import (
	"context"

	"github.com/Timasha/AuthService/pkg/errlist"
)

type EnableOtpAuthenticationRequest struct {
	UserID int64
}

type EnableOtpAuthenticationResponse struct {
	OtpKey string
	OtpURL string
}

func (c *Provider) EnableOtpAuthentication(
	ctx context.Context,
	req EnableOtpAuthenticationRequest,
) (resp EnableOtpAuthenticationResponse, err error) {
	user, err := c.userStorage.GetUserByUserID(ctx, req.UserID)
	if err != nil {
		return resp, err
	}

	if user.OtpEnabled {
		return resp, errlist.ErrOtpAlreadyEnabled
	}

	otpKey, otpURL, err := c.otpGenerator.GenerateKeys(user.Login)
	if err != nil {
		return resp, err
	}

	user.OtpEnabled = true
	user.OtpKey = otpKey

	err = c.userStorage.UpdateUserByLogin(ctx, user.Login, user)
	if err != nil {
		return resp, err
	}

	return EnableOtpAuthenticationResponse{
		OtpKey: otpKey,
		OtpURL: otpURL,
	}, nil
}
