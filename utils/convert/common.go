package convert

import (
	"github.com/Timasha/AuthService/internal/entities"
	"github.com/Timasha/AuthService/pkg/api"
)

func TokenPairFromProto(m *api.TokenPair) entities.TokenPair {
	return entities.TokenPair{
		AccessToken:  m.GetAccessToken(),
		RefreshToken: m.GetRefreshToken(),
	}
}

func TokenPairToProto(m entities.TokenPair) *api.TokenPair {
	return &api.TokenPair{
		AccessToken:  m.AccessToken,
		RefreshToken: m.RefreshToken,
	}
}
