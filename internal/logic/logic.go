package logic

import (
	"auth/internal/logic/dependencies"
	"auth/internal/logic/errs"
	"auth/internal/logic/iomodels"
)

type LogicProvider struct {
	userStorage    dependencies.UserStorage
	tokensProvider dependencies.TokensProvider
	passwordHasher dependencies.PasswordHasher
	uuidProvider   dependencies.UUIDProvider
	otpGenerator   dependencies.OtpGenerator
}

func (l *LogicProvider) Init(userStorage dependencies.UserStorage, tokensProvider dependencies.TokensProvider,
	passwordHasher dependencies.PasswordHasher, uuidProvider dependencies.UUIDProvider) {
	l.userStorage = userStorage
	l.tokensProvider = tokensProvider
	l.passwordHasher = passwordHasher
	l.uuidProvider = uuidProvider
}

func (l *LogicProvider) RegisterUser(args iomodels.RegisterUserArgs) (returned iomodels.RegisterUserReturned) {
	args.User.UserID = l.uuidProvider.GenerateUUID()
	var hashErr error
	args.User.Password, hashErr = l.passwordHasher.Hash(args.User.Password)
	if hashErr != nil {
		returned.Err = hashErr
		return
	}
	createErr := l.userStorage.CreateUser(args.Ctx, args.User)
	returned.Err = createErr
	return
}

func (l *LogicProvider) AuthenticateUserByLogin(args iomodels.AuthenticateUserByLoginArgs) (returned iomodels.AuthenticateUserByLoginReturned) {
	user, getErr := l.userStorage.GetUserByLogin(args.Ctx, args.Login)

	if getErr != nil {
		returned.Err = getErr
		return
	}

	if !l.passwordHasher.Compare(args.Password, user.Password) {
		returned.Err = errs.ErrInvalidPassword{}
		return
	}

	if user.OtpEnabled {
		returned.OtpEnabled = true

		returned.IntermediateToken, returned.Err = l.tokensProvider.CreateIntermediateToken(args.Login)

		return
	}

	returned.AuthInfo.AccessToken, returned.AuthInfo.RefreshToken, returned.Err = l.tokensProvider.CreateTokens(args.Login)

	return
}

func (l *LogicProvider) ContinueAuthenticateOtpUserByLogin(args iomodels.ContinueAuthenticateOtpUserByLoginArgs) (returned iomodels.ContinueAuthenticateOtpUserByLoginReturned) {
	user, getErr := l.userStorage.GetUserByLogin(args.Ctx, args.Login)

	if getErr != nil {
		returned.Err = getErr
		return
	}

	validErr := l.tokensProvider.ValidIntermediateToken(args.IntermediateToken, args.Login)

	if validErr != nil {
		returned.Err = validErr
		return
	}

	if !l.otpGenerator.ValidOtp(args.OtpCode, user.OtpKey) {
		returned.Err = errs.ErrInvalidOtp{}
	}

	returned.AuthInfo.AccessToken, returned.AuthInfo.RefreshToken, returned.Err = l.tokensProvider.CreateTokens(args.Login)

	return
}

func (l *LogicProvider) AuthorizeUser(args iomodels.AuthorizeUserArgs) (returned iomodels.AuthorizeUserReturned) {
	user, getErr := l.userStorage.GetUserByLogin(args.Ctx, args.Login)
	if getErr != nil {
		returned.Err = getErr
		return
	}

	validErr := l.tokensProvider.ValidAccessToken(args.AccessToken, args.Login)
	if validErr != nil {
		returned.Err = validErr
		return
	}
	returned.UserId = user.UserID
	return

}

func (l *LogicProvider) RefreshTokens(args iomodels.RefreshTokensArgs) (returned iomodels.RefreshTokensReturned) {
	_, getErr := l.userStorage.GetUserByLogin(args.Ctx, args.Login)

	if getErr != nil {
		returned.Err = getErr
		return
	}

	accessValidErr := l.tokensProvider.ValidAccessToken(args.AccessToken, args.Login)

	if accessValidErr != nil && accessValidErr != (errs.ErrExpiredAccessToken{}) {
		returned.Err = accessValidErr
		return
	}

	refreshValidErr := l.tokensProvider.ValidRefreshToken(args.RefreshToken, args.AccessToken)

	if refreshValidErr != nil {
		returned.Err = refreshValidErr
		return
	}

	returned.AccessToken, returned.RefreshToken, returned.Err = l.tokensProvider.CreateTokens(args.Login)

	return
}
