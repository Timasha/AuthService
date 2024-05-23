package convert

import (
	"github.com/Timasha/AuthService/internal/usecase"
	"github.com/Timasha/AuthService/pkg/api"
)

func EnableOtpAuthenticationResponseToProto(
	m usecase.EnableOtpAuthenticationResponse,
) (resp *api.EnableOtpAuthenticationResponse) {
	resp = &api.EnableOtpAuthenticationResponse{
		OtpKey: m.OtpKey,
		OtpUrl: m.OtpURL,
	}

	return resp
}
