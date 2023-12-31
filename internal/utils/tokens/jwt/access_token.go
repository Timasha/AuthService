package jwt

import (
	"AuthService/internal/logic"
	jwtErrs "AuthService/internal/utils/tokens/jwt/errs"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type AccessTokenClaims struct {
	jwt.RegisteredClaims
	Login string `json:"login"`
}

func (j *TokensProvider) CreateAccessToken(login string) (string, error) {
	var claims AccessTokenClaims = AccessTokenClaims{
		Login: login,
	}

	claims.ExpiresAt = jwt.NewNumericDate(time.Now().Add(time.Duration(j.AccessTokenLifeTime) * time.Minute))

	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)

	return token.SignedString([]byte(j.AccessTokenKey))
}

func (j *TokensProvider) ValidAccessToken(strToken string) (string, error) {
	token, parseErr := jwt.ParseWithClaims(strToken, &AccessTokenClaims{}, func(t *jwt.Token) (interface{}, error) {
		if t.Method != jwt.SigningMethodHS512 {
			return nil, jwtErrs.ErrWrongSingingMethod
		}
		return []byte(j.AccessTokenKey), nil
	})

	if parseErr == jwtErrs.ErrWrongSingingMethod || errors.Is(parseErr, jwt.ErrTokenMalformed) || errors.Is(parseErr, jwt.ErrTokenSignatureInvalid) {
		return "", logic.ErrInvalidAccessToken
	}else if parseErr != nil && !errors.Is(parseErr, jwt.ErrTokenExpired) {
		return "", parseErr
	}
	

	claims, ok := token.Claims.(*AccessTokenClaims)
	if !ok {
		return "", logic.ErrInvalidAccessToken
	}
	
	if errors.Is(parseErr, jwt.ErrTokenExpired) {
		return claims.Login, logic.ErrExpiredAccessToken
	}

	

	return claims.Login, nil
}
