package usecase

import (
	"context"

	"github.com/Timasha/AuthService/pkg/errlist"
)

type DisableOtpAuthenticationRequest struct {
	UserID int64
}

func (c *Provider) DisableOtpAuthentication(
	ctx context.Context,
	req DisableOtpAuthenticationRequest,
) (err error) {
	gettedUser, err := c.userStorage.GetUserByUserID(ctx, req.UserID)
	if err != nil {
		return err
	}

	if !gettedUser.OtpEnabled {
		return errlist.ErrOtpAlreadyDisabled
	}

	gettedUser.OtpEnabled = false
	gettedUser.OtpKey = ""

	err = c.userStorage.UpdateUserByLogin(ctx, gettedUser.Login, gettedUser)

	return err
}
