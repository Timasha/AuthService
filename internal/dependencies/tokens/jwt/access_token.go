package jwt

import (
	jwtErrs "auth/internal/dependencies/tokens/jwt/errs"
	"auth/internal/logic/errs"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type AccessTokenClaims struct {
	jwt.RegisteredClaims
	Login string `json:"login"`
}

func (j *TokensProvider) CreateAccessToken(login string) (string, error) {
	var claims AccessTokenClaims

	claims.Login = login
	claims.ExpiresAt.Time = time.Now().Add(time.Duration(j.AccessTokenLifeTime) * time.Minute)

	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)

	return token.SignedString([]byte(j.AccessTokenKey))
}

func (j *TokensProvider) ValidAccessToken(strToken, login string) error {
	token, parseErr := jwt.Parse(strToken, func(t *jwt.Token) (interface{}, error) {
		if t.Method != jwt.SigningMethodHS512 {
			return nil, jwtErrs.ErrWrongSingingMethod
		}
		return []byte(j.AccessTokenKey), nil
	})

	if claims, ok := token.Claims.(AccessTokenClaims); !ok || claims.Login != login || parseErr == jwtErrs.ErrWrongSingingMethod || errors.Is(parseErr, jwt.ErrTokenMalformed) || errors.Is(parseErr, jwt.ErrTokenSignatureInvalid) || !token.Valid {
		return errs.ErrInvalidAccessToken{}
	} else if errors.Is(parseErr, jwt.ErrTokenExpired) {
		return errs.ErrExpiredAccessToken{}
	} else if parseErr != nil {
		return parseErr
	}

	return nil
}
