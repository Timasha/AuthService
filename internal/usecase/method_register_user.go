package usecase

import (
	"context"

	"github.com/Timasha/AuthService/internal/entities"
	"github.com/Timasha/AuthService/utils/consts"
)

type RegisterUserRequest struct {
	Login    string
	Password string
}

func (c *Provider) RegisterUser(
	ctx context.Context,
	req RegisterUserRequest,
) (err error) {
	var user entities.User

	user.Role = consts.DefaultRole

	user.Password, err = c.passwordHasher.Hash(req.Password)
	if err != nil {
		return err
	}

	err = c.userStorage.CreateUser(ctx, user)

	return err
}
