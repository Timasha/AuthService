package cases

import (
	"auth/internal/logic/errs"
	"auth/internal/logic/models"
	"context"
	"log"
	"strings"
)

type UserStorageMock map[string]models.User

func (t UserStorageMock) CreateUser(ctx context.Context, user models.User) error {
	if _, ok := t[user.Login]; ok {
		return errs.ErrUserAlreadyExists{}
	}
	t[user.Login] = user
	return nil
}
func (t UserStorageMock) GetUserByLogin(ctx context.Context, login string) (models.User, error) {
	if _, ok := t[login]; !ok {
		return models.User{}, errs.ErrUserNotExists{}
	}
	return t[login], nil
}
func (t UserStorageMock) UpdateUserByLogin(ctx context.Context, login string, user models.User) error {
	if _, ok := t[login]; !ok {
		return errs.ErrUserNotExists{}
	}
	t[login] = user
	return nil
}
func (t UserStorageMock) DeleteUserByLogin(ctx context.Context, login string) error {
	if _, ok := t[login]; !ok {
		return errs.ErrUserNotExists{}
	}
	delete(t, login)
	return nil
}

type TokensProviderMock struct {
}

func (t *TokensProviderMock) CreateTokens(login string) (string, string, error) {
	access := "access." + login + ".true"
	refresh := "refresh." + access[:5] + ".true"
	return access, refresh, nil
}

func (t *TokensProviderMock) ValidAccessToken(token string, login string) error {
	tokensPart := strings.Split(token, ".")

	log.Println(len(tokensPart))

	if len(tokensPart) != 3 {
		return errs.ErrInvalidAccessToken{}
	}
	log.Println(tokensPart[0])
	log.Println(tokensPart[1], login)

	if tokensPart[0] != "access" || tokensPart[1] != login {
		return errs.ErrInvalidAccessToken{}
	} else if tokensPart[2] != "true" {
		return errs.ErrExpiredAccessToken{}
	}
	return nil
}

func (t *TokensProviderMock) ValidRefreshToken(refreshToken string, accessToken string) error {
	tokensPart := strings.Split(refreshToken, ".")

	if len(tokensPart) != 3 {
		return errs.ErrInvalidRefreshToken{}
	}

	if tokensPart[0] != "refresh" || tokensPart[1] != accessToken[:5] {
		return errs.ErrInvalidRefreshToken{}
	} else if tokensPart[2] != "true" {
		return errs.ErrExpiredRefreshToken{}
	}
	return nil
}

func (t *TokensProviderMock) CreateIntermediateToken(login string) (string, error) {
	return "intermediate." + login + ".true", nil
}

func (t *TokensProviderMock) ValidIntermediateToken(strToken, login string) error {
	tokensPart := strings.Split(strToken, ".")

	log.Println(len(tokensPart))

	if len(tokensPart) != 3 {
		return errs.ErrInvalidIntermediateToken{}
	}
	log.Println(tokensPart[0])
	log.Println(tokensPart[1], login)

	if tokensPart[0] != "intermediate" || tokensPart[1] != login {
		return errs.ErrInvalidIntermediateToken{}
	} else if tokensPart[2] != "true" {
		return errs.ErrExpiredIntermediateToken{}
	}
	return nil
}
