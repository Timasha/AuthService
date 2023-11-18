package jwt

import (
	"AuthService/internal/logic"
	jwtErrs "AuthService/internal/utils/tokens/jwt/errs"
	"errors"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type RefreshTokenClaims struct {
	jwt.RegisteredClaims
	AccessPart string `json:"accessPart"`
}

func (j *TokensProvider) CreateRefreshToken(accessToken string) (string, error) {
	accessParts := strings.Split(accessToken,".")
	var claims RefreshTokenClaims = RefreshTokenClaims{
		AccessPart: accessParts[2][:j.AccessPartLen],
	}

	claims.ExpiresAt = jwt.NewNumericDate(time.Now().Add(time.Hour * time.Duration(j.RefreshTokenLifeTime)))
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)

	return token.SignedString([]byte(j.RefreshTokenKey))
}

func (j *TokensProvider) ValidRefreshToken(refreshToken, accessToken string) error {
	token, parseErr := jwt.ParseWithClaims(refreshToken, &RefreshTokenClaims{}, func(t *jwt.Token) (interface{}, error) {
		if t.Method != jwt.SigningMethodHS512 {
			return nil, jwtErrs.ErrWrongSingingMethod
		}
		return []byte(j.RefreshTokenKey), nil
	})

	if parseErr == jwtErrs.ErrWrongSingingMethod || errors.Is(parseErr, jwt.ErrTokenMalformed) || errors.Is(parseErr, jwt.ErrTokenSignatureInvalid) {
		return logic.ErrInvalidRefreshToken
	}

	if errors.Is(parseErr, jwt.ErrTokenExpired) {
		return logic.ErrExpiredRefreshToken
	} else if parseErr != nil {
		return parseErr
	}

	accessParts := strings.Split(accessToken,".")

	if(len(accessParts) != 3){
		return logic.ErrInvalidAccessToken
	}

	claims, ok := token.Claims.(*RefreshTokenClaims)
	if !ok {
		return logic.ErrInvalidRefreshToken
	}

	if strings.Compare(claims.AccessPart, accessParts[2][:j.AccessPartLen]) != 0{
		return logic.ErrInvalidAccessToken
	}

	return nil
}
