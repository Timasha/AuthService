package dependencies

type TokensProvider interface {
	/* Implimentation should create access token based on login,
	create refresh token based on access,
	generate error if it was caused by something
	and return these values */
	CreateTokens(login string) (string, string, error)

	/*Implementation should valid access token by checking if thats is expired
	or token login is not equal to requested login.*/
	ValidAccessToken(token string, login string) error

	/*Implementation should valid refresh token by checking if it is expired,
	check relation with access token,
	and return error if it was caused by something*/
	ValidRefreshToken(refreshToken string, accessToken string) error

	CreateIntermediateToken(login string) (string, error)

	ValidIntermediateToken(strToken, login string) error
}
