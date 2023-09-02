package logic

import (
	"auth/internal/logic/dependencies"
	"auth/internal/logic/errs"
	"auth/internal/logic/models"
	"context"
)

type LogicProvider struct {
	userStorage    dependencies.UserStorage
	tokensProvider dependencies.TokensProvider
	passwordHasher dependencies.PasswordHasher
	uuidProvider   dependencies.UUIDProvider
}

func (l *LogicProvider) Init(userStorage dependencies.UserStorage, tokensProvider dependencies.TokensProvider,
	passwordHasher dependencies.PasswordHasher, uuidProvider dependencies.UUIDProvider) {
	l.userStorage = userStorage
	l.tokensProvider = tokensProvider
	l.passwordHasher = passwordHasher
	l.uuidProvider = uuidProvider
}

func (l *LogicProvider) RegisterUser(ctx context.Context, user models.User) error {

	user.UserID = l.uuidProvider.GenerateUUID()
	var hashErr error
	user.Password, hashErr = l.passwordHasher.Hash(user.Password)
	if hashErr != nil {
		return hashErr
	}
	createErr := l.userStorage.CreateUser(ctx, user)
	return createErr
}

func (l *LogicProvider) AuthenticateUserByLogin(ctx context.Context, login string, password string) (string, string, error) {
	user, getErr := l.userStorage.GetUserByLogin(ctx, login)

	if getErr != nil {
		return "", "", getErr
	}

	if !l.passwordHasher.Compare(password, user.Password) {
		return "", "", errs.ErrInvalidPassword{}
	}

	access, refresh, createTokenErr := l.tokensProvider.CreateTokens(login)

	return access, refresh, createTokenErr
}

func (l *LogicProvider) AuthorizeUser(ctx context.Context, accessToken string, login string) error {
	_, getErr := l.userStorage.GetUserByLogin(ctx, login)
	if getErr != nil {
		return getErr
	}

	validErr := l.tokensProvider.ValidAccessToken(accessToken, login)
	return validErr

}

func (l *LogicProvider) RefreshTokens(ctx context.Context, accessToken, refreshToken, login string) (string, string, error) {
	_, getErr := l.userStorage.GetUserByLogin(ctx, login)

	if getErr != nil {
		return "", "", getErr
	}

	accessValidErr := l.tokensProvider.ValidAccessToken(accessToken, login)

	if accessValidErr != nil && accessValidErr != (errs.ErrExpiredAccessToken{}) {
		return "", "", accessValidErr
	}

	refreshValidErr := l.tokensProvider.ValidRefreshToken(refreshToken, accessToken)

	if refreshValidErr != nil {
		return "", "", refreshValidErr
	}

	access, refresh, createTokenErr := l.tokensProvider.CreateTokens(login)

	return access, refresh, createTokenErr
}
