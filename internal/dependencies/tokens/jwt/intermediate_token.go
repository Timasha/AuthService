package jwt

import (
	jwtErrs "auth/internal/dependencies/tokens/jwt/errs"
	"auth/internal/logic/errs"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type IntermediateTokenClaims struct {
	jwt.RegisteredClaims
	Login string `json:"login"`
}

func (j *TokensProvider) CreateIntermediateToken(login string) (string, error) {
	var claims IntermediateTokenClaims

	claims.ExpiresAt.Time = time.Now().Add(time.Duration(j.IntermediateTokenLifeTime) * time.Second)
	claims.Login = login

	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)

	return token.SignedString([]byte(j.IntermediateTokenKey))
}

func (j *TokensProvider) ValidIntermediateToken(strToken, login string) error {
	token, parseErr := jwt.Parse(strToken, func(t *jwt.Token) (interface{}, error) {
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

	if claims, ok := token.Claims.(IntermediateTokenClaims); !ok || claims.Login != login {
		return errs.ErrInvalidRefreshToken{}
	}
	return nil
}
