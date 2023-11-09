package iomodels

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
