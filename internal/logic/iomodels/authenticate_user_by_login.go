package iomodels

import "context"

type AuthenticateUserByLoginArgs struct {
	Ctx      context.Context
	Login    string
	Password string
}

type AuthenticateUserByLoginReturned struct {
	OtpEnabled bool

	IntermediateToken string

	AuthInfo struct {
		AccessToken  string
		RefreshToken string
	}
	Err error
}
