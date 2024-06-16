package consts

import (
	"github.com/Timasha/AuthService/internal/entities"
	"github.com/Timasha/AuthService/pkg/api"
)

const (
	SqlxConnectFmt = `postgres://%s:%s@%s:%s/auth`
	TCP            = "tcp"
	UserIDCtxKey   = "user_id"
	OTPSecretSize  = 15
	JWTPartLen     = 3
)

var (
	AuthForMethods = []string{
		api.Auth_EnableOtpAuthentication_FullMethodName,
		api.Auth_DisableOtpAuthentication_FullMethodName,
	}
)

var (
	RootRole = entities.Role{
		ID:     0,
		Access: nil,
		Name:   "root",
	}
	DefaultRole = entities.Role{
		ID:     1,
		Access: entities.RoleAccess{0},
		Name:   "user",
	}
)
