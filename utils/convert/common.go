package convert

import (
	"github.com/Timasha/AuthService/internal/usecase"
	"github.com/Timasha/AuthService/pkg/api"
)

func TokenPairFromProto(m *api.TokenPair) usecase.TokenPair {
	return usecase.TokenPair{
		AccessToken:  m.GetAccessToken(),
		RefreshToken: m.GetRefreshToken(),
	}
}

func TokenPairToProto(m usecase.TokenPair) *api.TokenPair {
	return &api.TokenPair{
		AccessToken:  m.AccessToken,
		RefreshToken: m.RefreshToken,
	}
}
