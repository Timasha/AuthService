package logic

import (
	"auth/internal/logic/models"
	"context"
)

type AddRoleArgs struct {
	Ctx  context.Context
	Role models.Role
}

type AddRoleReturned struct {
	Err error
}

func (l *LogicProvider) AddRole(args AddRoleArgs) (returned AddRoleReturned) {
	createRoleErr := l.rolesStorage.CreateRole(args.Ctx, args.Role)
	returned.Err = createRoleErr
	return
}
