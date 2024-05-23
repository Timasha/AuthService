package jwt

import (
	"github.com/Timasha/AuthService/pkg/errlist"
	jwtErrs "github.com/Timasha/AuthService/utils/tokens/jwt/errs"

	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type AccessTokenClaims struct {
	jwt.RegisteredClaims
	Login string `json:"login"`
}

func (tp *TokensProvider) CreateAccessToken(login string) (string, error) {
	claims := AccessTokenClaims{
		Login: login,
	}

	claims.ExpiresAt = jwt.NewNumericDate(time.Now().Add(tp.cfg.AccessTokenLifeTime.Duration))

	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)

	return token.SignedString([]byte(tp.cfg.AccessTokenKey))
}

func (tp *TokensProvider) ValidAccessToken(strToken string) (login string, err error) {
	token, err := jwt.ParseWithClaims(
		strToken,
		&AccessTokenClaims{},
		func(t *jwt.Token) (interface{}, error) {
			if t.Method != jwt.SigningMethodHS512 {
				return nil, jwtErrs.ErrWrongSingingMethod
			}

			return []byte(tp.cfg.AccessTokenKey), nil
		},
	)

	if errors.Is(err, jwtErrs.ErrWrongSingingMethod) ||
		errors.Is(err, jwt.ErrTokenMalformed) ||
		errors.Is(err, jwt.ErrTokenSignatureInvalid) {
		return "", errlist.ErrInvalidAccessToken
	} else if err != nil &&
		!errors.Is(err, jwt.ErrTokenExpired) {
		return "", err
	}

	claims, ok := token.Claims.(*AccessTokenClaims)
	if !ok {
		return "", errlist.ErrInvalidAccessToken
	}

	if errors.Is(err, jwt.ErrTokenExpired) {
		return claims.Login, errlist.ErrExpiredAccessToken
	}

	return claims.Login, nil
}
