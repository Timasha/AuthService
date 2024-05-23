package convert

import (
	"github.com/Timasha/AuthService/internal/usecase"
	"github.com/Timasha/AuthService/pkg/api"
)

func ContinueAuthenticateOtpUserByLoginRequestFromProto(
	m *api.ContinueAuthenticateOtpUserByLoginRequest,
) (args usecase.ContinueAuthenticateOtpUserByLoginRequest) {
	args = usecase.ContinueAuthenticateOtpUserByLoginRequest{
		IntermediateToken: m.IntermediateToken,
		OtpCode:           m.OtpCode,
	}

	return args
}

func ContinueAuthenticateOtpUserByLoginResponseToProto(
	m usecase.ContinueAuthenticateOtpUserByLoginResponse,
) (resp *api.ContinueAuthenticateOtpUserByLoginResponse) {
	resp = &api.ContinueAuthenticateOtpUserByLoginResponse{
		TokenPair: TokenPairToProto(m.TokenPair),
	}

	return resp
}
