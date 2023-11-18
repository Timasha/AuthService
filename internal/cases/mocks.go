package cases

import (
	"AuthService/internal/logic"
	"AuthService/internal/logic/models"
	"context"
	"strings"
)

type UserStorageMock map[string]models.User

func (t UserStorageMock) CreateUser(ctx context.Context, user models.User) error {
	if _, ok := t[user.Login]; ok {
		return logic.ErrUserAlreadyExists
	}
	t[user.Login] = user
	return nil
}
func (t UserStorageMock) GetUserByLogin(ctx context.Context, login string) (models.User, error) {
	if _, ok := t[login]; !ok {
		return models.User{}, logic.ErrUserNotExists
	}
	return t[login], nil
}

func (t UserStorageMock) GetUserByUserId(ctx context.Context,userId string)(models.User,error){
	for _,user := range(t){
		if user.UserID == userId{
			return user,nil
		}
	}
	return models.User{},logic.ErrUserNotExists
}

func (t UserStorageMock) UpdateUserByLogin(ctx context.Context, login string, user models.User) error {
	if _, ok := t[login]; !ok {
		return logic.ErrUserNotExists
	}
	t[login] = user
	return nil
}
func (t UserStorageMock) DeleteUserByLogin(ctx context.Context, login string) error {
	if _, ok := t[login]; !ok {
		return logic.ErrUserNotExists
	}
	delete(t, login)
	return nil
}

type RolesStorageMock map[models.RoleId]models.Role

func (t RolesStorageMock) CreateRole(ctx context.Context, role models.Role) error {
	if _, ok := t[role.RoleId]; ok {
		return logic.ErrRoleAlreadyExists
	}
	t[role.RoleId] = role
	return nil
}
func (t RolesStorageMock) GetRoleById(ctx context.Context, roleId models.RoleId) (models.Role, error) {
	role, ok := t[roleId]
	if !ok {
		return models.Role{}, logic.ErrRoleNotExists
	}
	return role, nil
}
func (t RolesStorageMock) UpdateRoleById(ctx context.Context, roleId models.RoleId, role models.Role) error {
	if _, ok := t[roleId]; !ok {
		return logic.ErrRoleNotExists
	}
	t[roleId] = role
	return nil
}
func (t RolesStorageMock) DeleteRoleById(ctx context.Context, roleId models.RoleId) error {
	if _, ok := t[roleId]; !ok {
		return logic.ErrRoleNotExists
	}
	delete(t, roleId)
	return nil
}

type TokensProviderMock struct {
}

func (t *TokensProviderMock) CreateTokens(login string) (string, string, error) {
	access := "access." + login + ".true"
	refresh := "refresh." + access[:5] + ".true"
	return access, refresh, nil
}

func (t *TokensProviderMock) ValidAccessToken(token string) (string, error) {
	tokensPart := strings.Split(token, ".")

	if len(tokensPart) != 3 {
		return "", logic.ErrInvalidAccessToken
	}

	if tokensPart[0] != "access" {
		return "", logic.ErrInvalidAccessToken
	} else if tokensPart[2] != "true" {
		return tokensPart[1], logic.ErrExpiredAccessToken
	}
	return tokensPart[1], nil
}

func (t *TokensProviderMock) ValidRefreshToken(refreshToken string, accessToken string) error {
	tokensPart := strings.Split(refreshToken, ".")

	if len(tokensPart) != 3 {
		return logic.ErrInvalidRefreshToken
	}

	if tokensPart[0] != "refresh" || tokensPart[1] != accessToken[:5] {
		return logic.ErrInvalidRefreshToken
	} else if tokensPart[2] != "true" {
		return logic.ErrExpiredRefreshToken
	}
	return nil
}

func (t *TokensProviderMock) CreateIntermediateToken(login string) (string, error) {
	return "intermediate." + login + ".true", nil
}

func (t *TokensProviderMock) ValidIntermediateToken(strToken string) (string, error) {
	tokensPart := strings.Split(strToken, ".")

	if len(tokensPart) != 3 {
		return "", logic.ErrInvalidIntermediateToken
	}

	if tokensPart[0] != "intermediate" {
		return "", logic.ErrInvalidIntermediateToken
	} else if tokensPart[2] != "true" {
		return "", logic.ErrExpiredIntermediateToken
	}
	return tokensPart[1], nil
}

type OtpGeneratorMock struct {
}

func (o *OtpGeneratorMock) GenerateKeys(login string) (string, string, error) {
	return login, login, nil
}
func (o *OtpGeneratorMock) ValidOtp(passcode string, key string) bool {
	return passcode == key
}
