package convert

import (
	"github.com/Timasha/AuthService/internal/usecase"
	"github.com/Timasha/AuthService/pkg/api"
)

func RegisterUserRequestFromProto(m *api.RegisterUserRequest) (args usecase.RegisterUserRequest) {
	args = usecase.RegisterUserRequest{
		Login:    m.GetLogin(),
		Password: m.GetPassword(),
	}

	return args
}
