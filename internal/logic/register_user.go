package logic

import (
	"auth/internal/logic/models"
	"context"
)

type RegisterUserArgs struct {
	Ctx  context.Context
	User models.User
}

type RegisterUserReturned struct {
	Err error
}

func (l *LogicProvider) RegisterUser(args RegisterUserArgs) (returned RegisterUserReturned) {
	args.User.UserID = l.uuidProvider.GenerateUUID()
	var hashErr error
	args.User.Password, hashErr = l.passwordHasher.Hash(args.User.Password)
	if hashErr != nil {
		returned.Err = hashErr
		return
	}
	createErr := l.userStorage.CreateUser(args.Ctx, args.User)
	returned.Err = createErr
	return
}
