package jwt

import (
	"errors"
	"strings"
	"time"

	"github.com/Timasha/AuthService/utils/consts"

	"github.com/Timasha/AuthService/internal/entities"
	"github.com/Timasha/AuthService/pkg/errlist"
	jwtErrs "github.com/Timasha/AuthService/utils/tokens/jwt/errs"

	"github.com/golang-jwt/jwt/v5"
)

type RefreshTokenClaims struct {
	jwt.RegisteredClaims
	AccessPart string `json:"accessPart"`
}

func (tp *TokensProvider) CreateRefreshToken(accessToken string) (string, error) {
	accessParts := strings.Split(accessToken, ".")
	var claims = RefreshTokenClaims{
		AccessPart: accessParts[2][:tp.cfg.AccessPartLen],
	}

	claims.ExpiresAt = jwt.NewNumericDate(time.Now().Add(tp.cfg.RefreshTokenLifeTime.Duration))
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)

	return token.SignedString([]byte(tp.cfg.RefreshTokenKey))
}

func (tp *TokensProvider) ValidRefreshToken(tokenPair entities.TokenPair) error {
	token, parseErr := jwt.ParseWithClaims(
		tokenPair.RefreshToken,
		&RefreshTokenClaims{},
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
		return errlist.ErrInvalidRefreshToken
	}

	if errors.Is(parseErr, jwt.ErrTokenExpired) {
		return errlist.ErrExpiredRefreshToken
	} else if parseErr != nil {
		return parseErr
	}

	accessParts := strings.Split(tokenPair.AccessToken, ".")

	if len(accessParts) != consts.JWTPartLen {
		return errlist.ErrInvalidAccessToken
	}

	claims, ok := token.Claims.(*RefreshTokenClaims)
	if !ok {
		return errlist.ErrInvalidRefreshToken
	}

	if strings.Compare(claims.AccessPart, accessParts[2][:tp.cfg.AccessPartLen]) != 0 {
		return errlist.ErrInvalidAccessToken
	}

	return nil
}
