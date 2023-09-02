package jwt

import (
	jwtErrs "auth/internal/dependencies/tokens/jwt/errs"
	"auth/internal/logic/errs"
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
	var claims RefreshTokenClaims

	claims.ExpiresAt.Time = time.Now().Add(time.Duration(j.RefreshTokenLifeTime) * time.Hour)
	claims.AccessPart = accessToken[:j.AccessPartLen]

	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)

	return token.SignedString([]byte(j.RefreshTokenKey))
}

func (j *TokensProvider) ValidRefreshToken(refreshToken, accessToken string) error {
	token, parseErr := jwt.Parse(refreshToken, func(t *jwt.Token) (interface{}, error) {
		if t.Method != jwt.SigningMethodHS512 {
			return nil, jwtErrs.ErrWrongSingingMethod
		}
		return j.RefreshTokenKey, nil
	})
	if parseErr == jwtErrs.ErrWrongSingingMethod || errors.Is(parseErr, jwt.ErrTokenMalformed) || errors.Is(parseErr, jwt.ErrTokenSignatureInvalid) || !token.Valid {
		return errs.ErrInvalidRefreshToken{}
	} else if errors.Is(parseErr, jwt.ErrTokenExpired) {
		return errs.ErrExpiredRefreshToken{}
	} else if parseErr != nil {
		return parseErr
	}

	if claims, ok := token.Claims.(RefreshTokenClaims); !ok || strings.Compare(claims.AccessPart, accessToken[:j.AccessPartLen]) != 0 {
		return errs.ErrInvalidRefreshToken{}
	}
	return nil
}
