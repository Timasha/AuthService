package iomodels

import "context"

type RefreshTokensArgs struct {
	Ctx          context.Context
	AccessToken  string
	RefreshToken string
	Login        string
}

type RefreshTokensReturned struct {
	AccessToken  string
	RefreshToken string
	Err          error
}
