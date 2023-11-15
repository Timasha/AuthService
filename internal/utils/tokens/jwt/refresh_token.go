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
	var claims RefreshTokenClaims = RefreshTokenClaims{
		AccessPart: accessToken[:j.AccessPartLen],
	}

	claims.ExpiresAt = jwt.NewNumericDate(time.Now().Add(time.Duration(j.RefreshTokenLifeTime) * time.Hour))

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

	claims, ok := token.Claims.(*RefreshTokenClaims)
	if !ok || strings.Compare(claims.AccessPart, accessToken[:j.AccessPartLen]) != 0 || !token.Valid {
		return logic.ErrInvalidRefreshToken
	}

	if errors.Is(parseErr, jwt.ErrTokenExpired) {
		return logic.ErrExpiredRefreshToken
	} else if parseErr != nil {
		return parseErr
	}

	return nil
}
