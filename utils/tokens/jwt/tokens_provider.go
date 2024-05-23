package jwt

import "github.com/Timasha/AuthService/internal/entities"

type TokensProvider struct {
	cfg Config
}

func New(
	cfg Config,
) (t *TokensProvider) {
	t = &TokensProvider{
		cfg: cfg,
	}

	return
}

func (tp *TokensProvider) CreateTokens(login string) (tokenPair entities.TokenPair, err error) {
	accessToken, err := tp.CreateAccessToken(login)
	if err != nil {
		return tokenPair, err
	}

	refreshToken, err := tp.CreateRefreshToken(accessToken)
	if err != nil {
		return tokenPair, err
	}

	tokenPair = entities.TokenPair{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	return tokenPair, nil
}
