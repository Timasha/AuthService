package dependencies

import (
	"auth/internal/logic/models"
	"context"
)

type UserStorage interface {
	CreateUser(ctx context.Context, user models.User) error
	GetUserByLogin(ctx context.Context, login string) (models.User, error)
	UpdateUserByLogin(ctx context.Context, login string, user models.User) error
	DeleteUserByLogin(ctx context.Context, login string) error
}
