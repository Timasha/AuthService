package jwt

type TokensProvider struct {
	AccessTokenKey      string
	AccessTokenLifeTime int64

	RefreshTokenKey      string
	RefreshTokenLifeTime int64
	AccessPartLen        int

	IntermediateTokenKey      string
	IntermediateTokenLifeTime int64
}

func New(accessTokenKey, refreshTokenKey string, accessTokenLifeTime, refreshTokenLifeTime int64, accessPartLen int) (t *TokensProvider) {
	t = &TokensProvider{
		AccessTokenKey: accessTokenKey,
		AccessTokenLifeTime: accessTokenLifeTime,

		RefreshTokenKey: refreshTokenKey,
		RefreshTokenLifeTime: refreshTokenLifeTime,
		AccessPartLen: accessPartLen,
	}
	return
}

func (t *TokensProvider) CreateTokens(login string) (string, string, error) {
	var (
		accessToken          string
		accessTokenCreateErr error

		refreshToken          string
		refreshTokenCreateErr error
	)

	accessToken, accessTokenCreateErr = t.CreateAccessToken(login)
	if accessTokenCreateErr != nil {
		return "", "", accessTokenCreateErr
	}

	refreshToken, refreshTokenCreateErr = t.CreateRefreshToken(accessToken)
	if refreshTokenCreateErr != nil {
		return "", "", refreshTokenCreateErr
	}

	return accessToken, refreshToken, nil
}
