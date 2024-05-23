package convert

import (
	"github.com/Timasha/AuthService/internal/usecase"
	"github.com/Timasha/AuthService/pkg/api"
)

func RefreshTokensRequestFromProto(m *api.RefreshTokensRequest) (args usecase.RefreshTokensRequest) {
	args = usecase.RefreshTokensRequest{
		TokenPair: TokenPairFromProto(m.GetTokenPair()),
	}

	return args
}

func RefreshTokensResponseToProto(m usecase.RefreshTokensResponse) (resp *api.RefreshTokensResponse) {
	resp = &api.RefreshTokensResponse{
		TokenPair: TokenPairToProto(m.TokenPair),
	}

	return resp
}
