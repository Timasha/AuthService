package usecase

import (
	"context"

	"github.com/Timasha/AuthService/internal/entities"
	"github.com/Timasha/AuthService/pkg/errlist"
)

type AuthorizeUserRequest struct {
	AccessToken        string
	RequiredRoleAccess entities.RoleAccess
}

type AuthorizeUserResponse struct {
	UserID int64
}

func (c *Provider) AuthorizeUser(
	ctx context.Context,
	req AuthorizeUserRequest,
) (resp AuthorizeUserResponse, err error) {
	login, err := c.tokensProvider.ValidAccessToken(req.AccessToken)
	if err != nil {
		return resp, err
	}

	user, err := c.userStorage.GetUserByLogin(ctx, login)
	if err != nil {
		return resp, err
	}

	if req.RequiredRoleAccess != nil && !user.HaveAccess(req.RequiredRoleAccess) {
		return resp, errlist.ErrRoleHasNoAccess
	}

	return AuthorizeUserResponse{
		UserID: user.ID,
	}, nil
}
