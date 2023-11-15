package logic

import (
	"AuthService/internal/logic/models"
	"AuthService/internal/utils/errsutil"
	"context"
)

type AuthorizeUserArgs struct {
	Ctx            context.Context
	AccessToken    string
	RequiredRoleId models.RoleId
}

type AuthorizeUserReturned struct {
	UserId string
	Err    error
}

func (l *LogicProvider) AuthorizeUser(args AuthorizeUserArgs) (returned AuthorizeUserReturned) {

	login, validErr := l.tokensProvider.ValidAccessToken(args.AccessToken)
	if validErr != nil {
		returned.Err = validErr
		return
	}

	user, getErr := l.userStorage.GetUserByLogin(args.Ctx, login)
	if getErr != nil {
		returned.Err = getErr
		return
	}

	if !user.Role.RoleId.HavaAccess(args.RequiredRoleId) {
		returned.Err = ErrRoleHasNoAccess
		return
	}

	returned.UserId = user.UserID
	return

}

var ErrRoleHasNoAccess errsutil.AuthErr = errsutil.AuthErr{
	Msg:     "role has no access",
	ErrCode: errsutil.ErrRoleHasNoAccessCode,
}
