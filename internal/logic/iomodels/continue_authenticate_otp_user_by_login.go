package iomodels

import "context"

type ContinueAuthenticateOtpUserByLoginArgs struct {
	Ctx context.Context

	Login             string
	IntermediateToken string
	OtpCode           string
}

type ContinueAuthenticateOtpUserByLoginReturned struct {
	AuthInfo struct {
		AccessToken  string
		RefreshToken string
	}
	Err error
}
