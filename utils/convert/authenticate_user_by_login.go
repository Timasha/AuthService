package convert

import (
	"github.com/Timasha/AuthService/internal/usecase"
	"github.com/Timasha/AuthService/pkg/api"
)

func AuthenticateUserByLoginRequestFromProto(
	m *api.AuthenticateUserByLoginRequest,
) (args usecase.AuthenticateUserByLoginRequest) {
	args = usecase.AuthenticateUserByLoginRequest{
		Login:    m.GetLogin(),
		Password: m.GetPassword(),
	}

	return args
}

func AuthenticateUserByLoginResponseToProto(
	m usecase.AuthenticateUserByLoginResponse,
) (resp *api.AuthenticateUserByLoginResponse) {
	resp = &api.AuthenticateUserByLoginResponse{
		OtpEnabled:        m.OtpEnabled,
		IntermediateToken: m.IntermediateToken,
	}
	if m.TokenPair != nil {
		resp.TokenPair = TokenPairToProto(*m.TokenPair)
	}

	return resp
}
