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

func (t *TokensProvider) Init(accessTokenKey, refreshTokenKey string, accessTokenLifeTime, refreshTokenLifeTime int64, accessPartLen int) {
	t.AccessTokenKey = accessTokenKey
	t.RefreshTokenKey = refreshTokenKey
	t.AccessTokenLifeTime = accessTokenLifeTime
	t.RefreshTokenLifeTime = refreshTokenLifeTime
	t.AccessPartLen = accessPartLen
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
