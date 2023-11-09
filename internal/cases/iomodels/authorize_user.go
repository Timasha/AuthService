package iomodels

import "context"

type AuthorizeUserArgs struct {
	Ctx         context.Context
	AccessToken string
	Login       string
}

type AuthorizeUserReturned struct {
	UserId string
	Err    error
}
