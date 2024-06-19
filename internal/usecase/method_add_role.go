package usecase

import (
	"context"

	"github.com/Timasha/AuthService/internal/entities"
)

type AddRoleRequest struct {
	Role entities.Role
}

func (c *Provider) AddRole(
	ctx context.Context,
	req AddRoleRequest,
) (err error) {
	err = c.rolesStorage.CreateRole(ctx, req.Role)

	return err
}
