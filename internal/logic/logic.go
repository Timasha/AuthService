package logic

import (
	"AuthService/internal/logic/models"
	"AuthService/internal/utils/errsutil"
	"context"
)

type LogicProvider struct {
	userStorage  UserStorage
	rolesStorage RolesStorage

	tokensProvider TokensProvider
	passwordHasher PasswordHasher
	uuidProvider   UUIDProvider
	otpGenerator   OtpGenerator
}

func New(userStorage UserStorage, rolesStorage RolesStorage, tokensProvider TokensProvider,
	passwordHasher PasswordHasher, uuidProvider UUIDProvider, otpGenerator OtpGenerator) (l *LogicProvider) {
	l = &LogicProvider{
		userStorage: userStorage,
		rolesStorage: rolesStorage,

		tokensProvider: tokensProvider,
		passwordHasher: passwordHasher,
		uuidProvider: uuidProvider,
		otpGenerator: otpGenerator,
	}
	return
}

type UserStorage interface {
	CreateUser(ctx context.Context, user models.User) error
	GetUserByLogin(ctx context.Context, login string) (models.User, error)
	UpdateUserByLogin(ctx context.Context, login string, user models.User) error
	DeleteUserByLogin(ctx context.Context, login string) error
}

var (
	ErrUserAlreadyExists errsutil.AuthErr = errsutil.AuthErr{
		Msg:     "user already exists",
		ErrCode: errsutil.ErrUserAlreadyExistsCode,
	}
	ErrUserNotExists errsutil.AuthErr = errsutil.AuthErr{
		Msg:     "user not exists",
		ErrCode: errsutil.ErrUserNotExistsCode,
	}
)

type RolesStorage interface {
	CreateRole(ctx context.Context, role models.Role) error
	GetRoleById(ctx context.Context, roleId models.RoleId) (models.Role, error)
	UpdateRoleById(ctx context.Context, roleId models.RoleId, role models.Role) error
	DeleteRoleById(ctx context.Context, roleId models.RoleId) error
}

var(
	ErrRoleAlreadyExists errsutil.AuthErr = errsutil.AuthErr {
		Msg: "role already exists",
		ErrCode: errsutil.ErrRoleAlreadyExistsCode,
	}
	ErrRoleNotExists errsutil.AuthErr = errsutil.AuthErr {
		Msg: "role not exists",
		ErrCode: errsutil.ErrRoleNotExistsCode,
	}	
)

type TokensProvider interface {
	/* Implimentation should create access token based on login,
	create refresh token based on access,
	generate error if it was caused by something
	and return these values */
	CreateTokens(login string) (string, string, error)

	/*Implementation should valid access token by checking if thats is expired or invalid by structure and return login and error*/
	ValidAccessToken(token string) (string, error)

	/*Implementation should valid refresh token by checking if it is expired,
	check relation with access token,
	and return error if it was caused by something*/
	ValidRefreshToken(refreshToken string, accessToken string) error

	CreateIntermediateToken(login string) (string, error)

	ValidIntermediateToken(strToken string) (string, error)
}

var (
	ErrExpiredAccessToken errsutil.AuthErr = errsutil.AuthErr{
		Msg:     "access token is expired",
		ErrCode: errsutil.ErrExpiredAccessTokenCode,
	}
	ErrExpiredIntermediateToken errsutil.AuthErr = errsutil.AuthErr{
		Msg:     "intermediate token is expired",
		ErrCode: errsutil.ErrExpiredIntermediateTokenCode,
	}
	ErrExpiredRefreshToken errsutil.AuthErr = errsutil.AuthErr{
		Msg:     "refresh token is expired",
		ErrCode: errsutil.ErrExpiredRefreshTokenCode,
	}

	ErrInvalidAccessToken errsutil.AuthErr = errsutil.AuthErr{
		Msg:     "invalid access token",
		ErrCode: errsutil.ErrInvalidAccessTokenCode,
	}
	ErrInvalidIntermediateToken errsutil.AuthErr = errsutil.AuthErr{
		Msg:     "invalid intermediate token",
		ErrCode: errsutil.ErrInvalidIntermediateTokenCode,
	}
	ErrInvalidRefreshToken errsutil.AuthErr = errsutil.AuthErr{
		Msg:     "invalid refresh token",
		ErrCode: errsutil.ErrInvalidRefreshTokenCode,
	}
)

type PasswordHasher interface {
	Hash(password string) (string, error)
	Compare(password, hash string) bool
}

type UUIDProvider interface {
	GenerateUUID() string
}

type OtpGenerator interface {
	GenerateKeys(login string) (string, string, error)
	ValidOtp(passcode string, key string) bool
}
