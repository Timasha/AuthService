package jwt

import (
	"AuthService/internal/logic"
	jwtErrs "AuthService/internal/utils/tokens/jwt/errs"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type IntermediateTokenClaims struct {
	jwt.RegisteredClaims
	Login string `json:"login"`
}

func (j *TokensProvider) CreateIntermediateToken(login string) (string, error) {
	var claims IntermediateTokenClaims = IntermediateTokenClaims{
		Login: login,
	}

	claims.ExpiresAt = jwt.NewNumericDate(time.Now().Add(time.Duration(j.IntermediateTokenLifeTime) * time.Minute))

	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)

	return token.SignedString([]byte(j.IntermediateTokenKey))
}

func (j *TokensProvider) ValidIntermediateToken(strToken string) (string, error) {
	token, parseErr := jwt.ParseWithClaims(strToken, &IntermediateTokenClaims{}, func(t *jwt.Token) (interface{}, error) {
		if t.Method != jwt.SigningMethodHS512 {
			return nil, jwtErrs.ErrWrongSingingMethod
		}
		return []byte(j.RefreshTokenKey), nil
	})

	if parseErr == jwtErrs.ErrWrongSingingMethod || errors.Is(parseErr, jwt.ErrTokenMalformed) || errors.Is(parseErr, jwt.ErrTokenSignatureInvalid) {
		return "", logic.ErrInvalidIntermediateToken
	}else if parseErr != nil && !errors.Is(parseErr, jwt.ErrTokenExpired) {
		return "", parseErr
	}

	claims, ok := token.Claims.(*IntermediateTokenClaims)
	if !ok {
		return "", logic.ErrInvalidIntermediateToken
	}

	if errors.Is(parseErr, jwt.ErrTokenExpired) {
		return claims.Login, logic.ErrExpiredIntermediateToken
	}

	return claims.Login, nil

}
