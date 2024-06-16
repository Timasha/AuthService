package usecase

import (
	"context"

	"github.com/Timasha/AuthService/internal/entities"
)

type (
	UserStorage interface {
		CreateUser(ctx context.Context, user entities.User) (err error)
		GetUserByLogin(ctx context.Context, login string) (user entities.User, err error)
		GetUserByUserID(ctx context.Context, userID int64) (user entities.User, err error)
		UpdateUserByLogin(ctx context.Context, login string, user entities.User) (err error)
		DeleteUserByLogin(ctx context.Context, login string) (err error)
	}

	RolesStorage interface {
		CreateRole(ctx context.Context, role entities.Role) (err error)
		GetRoleByID(ctx context.Context, roleID int64) (role entities.Role, err error)
		UpdateRoleByID(ctx context.Context, roleID int64, role entities.Role) (err error)
		DeleteRoleByID(ctx context.Context, roleID int64) (err error)
	}

	TokensProvider interface {
		/*CreateTokens should create access token based on login,
		create refresh token based on access,
		generate error if it was caused by something
		and return these values */
		CreateTokens(login string) (tokenPair TokenPair, err error)

		/*ValidAccessToken should valid access token by checking if that is expired
		or invalid by structure and return login and error*/
		ValidAccessToken(token string) (login string, err error)

		/*ValidRefreshToken should valid refresh token by checking if it is expired,
		check relation with access token,
		and return error if it was caused by something*/
		ValidRefreshToken(tokenPair TokenPair) (err error)

		CreateIntermediateToken(login string) (token string, err error)

		ValidIntermediateToken(strToken string) (login string, err error)
	}

	PasswordHasher interface {
		Hash(password string) (hashedPassword string, err error)
		Compare(password, hash string) (isEqual bool)
	}

	UUIDProvider interface {
		GenerateUUID() string
	}

	OtpGenerator interface {
		GenerateKeys(login string) (secret string, link string, err error)
		ValidOtp(passcode string, key string) (valid bool)
	}
)
