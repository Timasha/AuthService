package jwt

import (
	"errors"
	"time"

	"github.com/Timasha/AuthService/pkg/errlist"
	jwtErrs "github.com/Timasha/AuthService/utils/tokens/jwt/errs"

	"github.com/golang-jwt/jwt/v5"
)

type IntermediateTokenClaims struct {
	jwt.RegisteredClaims
	Login string `json:"login"`
}

func (tp *TokensProvider) CreateIntermediateToken(login string) (string, error) {
	var claims = IntermediateTokenClaims{
		Login: login,
	}

	claims.ExpiresAt = jwt.NewNumericDate(time.Now().Add(tp.cfg.IntermediateTokenLifeTime.Duration))

	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)

	return token.SignedString([]byte(tp.cfg.IntermediateTokenKey))
}

func (tp *TokensProvider) ValidIntermediateToken(strToken string) (string, error) {
	token, parseErr := jwt.ParseWithClaims(
		strToken,
		&IntermediateTokenClaims{},
		func(t *jwt.Token) (interface{}, error) {
			if t.Method != jwt.SigningMethodHS512 {
				return nil, jwtErrs.ErrWrongSingingMethod
			}

			return []byte(tp.cfg.RefreshTokenKey), nil
		},
	)

	if errors.Is(parseErr, jwtErrs.ErrWrongSingingMethod) ||
		errors.Is(parseErr, jwt.ErrTokenMalformed) ||
		errors.Is(parseErr, jwt.ErrTokenSignatureInvalid) {
		return "", errlist.ErrInvalidIntermediateToken
	} else if parseErr != nil && !errors.Is(parseErr, jwt.ErrTokenExpired) {
		return "", parseErr
	}

	claims, ok := token.Claims.(*IntermediateTokenClaims)
	if !ok {
		return "", errlist.ErrInvalidIntermediateToken
	}

	if errors.Is(parseErr, jwt.ErrTokenExpired) {
		return claims.Login, errlist.ErrExpiredIntermediateToken
	}

	return claims.Login, nil
}
