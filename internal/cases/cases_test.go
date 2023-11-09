package cases_test

import (
	"auth/internal/cases"
	casesErrs "auth/internal/cases/errs"
	"auth/internal/cases/iomodels"
	"auth/internal/dependencies/config"
	"auth/internal/dependencies/password"
	"auth/internal/dependencies/uuid"
	"auth/internal/logic"
	"auth/internal/logic/dependencies"
	logicErrs "auth/internal/logic/errs"
	"auth/internal/logic/models"

	"context"
	"testing"

	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
)

func CasesMocks(t *testing.T) *cases.CasesProvider {
	var logicProvider *logic.LogicProvider = &logic.LogicProvider{}
	var userStorage dependencies.UserStorage = make(cases.UserStorageMock)

	hasher := &password.BcryptPasswordHasher{}
	hashedPassword, err := hasher.Hash("SamplePassword")

	if err != nil {
		t.Fatalf("Cannot hash password: %v", err)
	}

	userStorage.CreateUser(context.Background(), models.User{
		UserID:   "Some_UUID",
		Login:    "SampleLogin",
		Password: hashedPassword,
	})
	hashedPassword, err = hasher.Hash("SamplePassword2")
	userStorage.CreateUser(context.Background(), models.User{
		UserID:   "Some_UUID2",
		Login:    "SampleLogin2",
		Password: hashedPassword,
	})

	logicProvider.Init(userStorage, &cases.TokensProviderMock{}, &password.BcryptPasswordHasher{}, &uuid.GoogleUUIDProvider{})

	var casesProvider *cases.CasesProvider = &cases.CasesProvider{}

	casesProvider.Init(&config.JSONConfig{
		MinLoginLen:    10,
		MinPasswordLen: 10,
	}, zerolog.New(zerolog.NewConsoleWriter()), logicProvider)
	return casesProvider
}

func TestCasesProvider_RegisterUser(t *testing.T) {
	casesProvider := CasesMocks(t)

	tests := []struct {
		name         string
		args         iomodels.RegisterUserArgs
		wantReturned iomodels.RegisterUserReturned
	}{
		{
			name: "success1",
			args: iomodels.RegisterUserArgs{
				Ctx: context.Background(),
				User: models.User{
					Login:    "SampleLogin3",
					Password: "SamplePassword3",
				},
			},
			wantReturned: iomodels.RegisterUserReturned{
				Err: nil,
			},
		},
		{
			name: "user_already_exists",
			args: iomodels.RegisterUserArgs{
				Ctx: context.Background(),
				User: models.User{
					Login:    "SampleLogin",
					Password: "SamplePasswordDiff",
				},
			},
			wantReturned: iomodels.RegisterUserReturned{
				Err: logicErrs.ErrUserAlreadyExists{},
			},
		},
		{
			name: "too_short_login",
			args: iomodels.RegisterUserArgs{
				Ctx: context.Background(),
				User: models.User{
					Login:    "asd",
					Password: "asdasdasdasd",
				},
			},
			wantReturned: iomodels.RegisterUserReturned{
				Err: casesErrs.ErrTooShortLoginOrPassword{},
			},
		},
		{
			name: "too_short_password",
			args: iomodels.RegisterUserArgs{
				Ctx: context.Background(),
				User: models.User{
					Login:    "asdasdasdasd",
					Password: "asd",
				},
			},
			wantReturned: iomodels.RegisterUserReturned{
				Err: casesErrs.ErrTooShortLoginOrPassword{},
			},
		},
	}
	for _, tCase := range tests {
		t.Run(tCase.name, func(t *testing.T) {
			returned := casesProvider.RegisterUser(tCase.args)
			assert.Equal(t, tCase.wantReturned, returned)
		})
	}
}

func TestCasesProvider_AuthenticateUserByLogin(t *testing.T) {
	casesProvider := CasesMocks(t)

	tests := []struct {
		name         string
		args         iomodels.AuthenticateUserByLoginArgs
		wantReturned iomodels.AuthenticateUserByLoginReturned
	}{
		{
			name: "success1",
			args: iomodels.AuthenticateUserByLoginArgs{
				Ctx:      context.Background(),
				Login:    "SampleLogin",
				Password: "SamplePassword",
			},
			wantReturned: iomodels.AuthenticateUserByLoginReturned{
				OtpEnabled:        false,
				IntermediateToken: "",

				AuthInfo: struct {
					AccessToken  string
					RefreshToken string
				}{
					AccessToken:  "access.SampleLogin.true",
					RefreshToken: "refresh.acces.true",
				},

				Err: nil,
			},
		},
		{
			name: "success2",
			args: iomodels.AuthenticateUserByLoginArgs{
				Ctx:      context.Background(),
				Login:    "SampleLogin2",
				Password: "SamplePassword2",
			},
			wantReturned: iomodels.AuthenticateUserByLoginReturned{
				OtpEnabled:        false,
				IntermediateToken: "",

				AuthInfo: struct {
					AccessToken  string
					RefreshToken string
				}{
					AccessToken:  "access.SampleLogin2.true",
					RefreshToken: "refresh.acces.true",
				},

				Err: nil,
			},
		},
		{
			name: "user_not_exists",
			args: iomodels.AuthenticateUserByLoginArgs{
				Ctx:      context.Background(),
				Login:    "SomeLoginExample",
				Password: "SomePassword",
			},
			wantReturned: iomodels.AuthenticateUserByLoginReturned{
				OtpEnabled:        false,
				IntermediateToken: "",

				AuthInfo: struct {
					AccessToken  string
					RefreshToken string
				}{
					AccessToken:  "",
					RefreshToken: "",
				},

				Err: casesErrs.ErrInvalidLoginOrPassword{},
			},
		},
		{
			name: "too_short_login",
			args: iomodels.AuthenticateUserByLoginArgs{
				Ctx:      context.Background(),
				Login:    "ShortL",
				Password: "SamplePassword",
			},
			wantReturned: iomodels.AuthenticateUserByLoginReturned{
				OtpEnabled:        false,
				IntermediateToken: "",

				AuthInfo: struct {
					AccessToken  string
					RefreshToken string
				}{
					AccessToken:  "",
					RefreshToken: "",
				},

				Err: casesErrs.ErrTooShortLoginOrPassword{},
			},
		},
		{
			name: "too_short_password",
			args: iomodels.AuthenticateUserByLoginArgs{
				Ctx:      context.Background(),
				Login:    "SampleLogin",
				Password: "ShortP",
			},
			wantReturned: iomodels.AuthenticateUserByLoginReturned{
				OtpEnabled:        false,
				IntermediateToken: "",

				AuthInfo: struct {
					AccessToken  string
					RefreshToken string
				}{
					AccessToken:  "",
					RefreshToken: "",
				},

				Err: casesErrs.ErrTooShortLoginOrPassword{},
			},
		},
	}
	for _, tCase := range tests {
		t.Run(tCase.name, func(t *testing.T) {
			returned := casesProvider.AuthenticateUserByLogin(tCase.args)
			assert.Equal(t, tCase.wantReturned, returned)
		})
	}
}

func TestCasesProvider_AuthorizeUser(t *testing.T) {

	casesProvider := CasesMocks(t)

	tests := []struct {
		name         string
		args         iomodels.AuthorizeUserArgs
		wantReturned iomodels.AuthorizeUserReturned
	}{
		{
			name: "success1",

			args: iomodels.AuthorizeUserArgs{
				Ctx:         context.Background(),
				AccessToken: "access.SampleLogin.true",
				Login:       "SampleLogin",
			},
			wantReturned: iomodels.AuthorizeUserReturned{
				UserId: "Some_UUID",
				Err:    nil,
			},
		},
		{
			name: "invalid_token1",
			args: iomodels.AuthorizeUserArgs{
				Ctx:         context.Background(),
				AccessToken: "refresh.SampleLogin.true",
				Login:       "SampleLogin",
			},
			wantReturned: iomodels.AuthorizeUserReturned{
				UserId: "",
				Err:    logicErrs.ErrInvalidAccessToken{},
			},
		},
		{
			name: "invalid_token2",
			args: iomodels.AuthorizeUserArgs{
				Ctx:         context.Background(),
				AccessToken: "access.SampleLogin2.true",
				Login:       "SampleLogin",
			},
			wantReturned: iomodels.AuthorizeUserReturned{
				UserId: "",
				Err:    logicErrs.ErrInvalidAccessToken{},
			},
		},
		{
			name: "not_existing_user1",
			args: iomodels.AuthorizeUserArgs{
				Ctx:         context.Background(),
				AccessToken: "access.SampleLogin.true",
				Login:       "SampleNotExistingLogin",
			},
			wantReturned: iomodels.AuthorizeUserReturned{
				UserId: "",
				Err:    logicErrs.ErrUserNotExists{},
			},
		},
		{
			name: "too_short_login",
			args: iomodels.AuthorizeUserArgs{
				Ctx:         context.Background(),
				AccessToken: "access.ShortL.true",
				Login:       "ShortL",
			},
			wantReturned: iomodels.AuthorizeUserReturned{
				UserId: "",
				Err:    casesErrs.ErrTooShortLoginOrPassword{},
			},
		},
		{
			name: "expired_token1",
			args: iomodels.AuthorizeUserArgs{
				Ctx:         context.Background(),
				AccessToken: "access.SampleLogin.false",
				Login:       "SampleLogin",
			},
			wantReturned: iomodels.AuthorizeUserReturned{
				UserId: "",
				Err:    logicErrs.ErrExpiredAccessToken{},
			},
		},
	}
	for _, tCase := range tests {
		t.Run(tCase.name, func(t *testing.T) {
			returned := casesProvider.AuthorizeUser(tCase.args)
			assert.Equal(t, tCase.wantReturned, returned)
		})
	}
}

func TestCasesProvider_RefreshTokens(t *testing.T) {
	casesProvider := CasesMocks(t)

	tests := []struct {
		name         string
		args         iomodels.RefreshTokensArgs
		wantReturned iomodels.RefreshTokensReturned
	}{
		{
			name: "success1",
			args: iomodels.RefreshTokensArgs{
				Ctx:          context.Background(),
				RefreshToken: "refresh.acces.true",
				AccessToken:  "access.SampleLogin.true",
				Login:        "SampleLogin",
			},
			wantReturned: iomodels.RefreshTokensReturned{
				AccessToken:  "access.SampleLogin.true",
				RefreshToken: "refresh.acces.true",
				Err:          nil,
			},
		},
		{
			name: "success2",
			args: iomodels.RefreshTokensArgs{
				Ctx:          context.Background(),
				RefreshToken: "refresh.acces.true",
				AccessToken:  "access.SampleLogin.false",
				Login:        "SampleLogin",
			},
			wantReturned: iomodels.RefreshTokensReturned{
				AccessToken:  "access.SampleLogin.true",
				RefreshToken: "refresh.acces.true",
				Err:          nil,
			},
		},
		{
			name: "too_short_login",
			args: iomodels.RefreshTokensArgs{
				Ctx:          context.Background(),
				RefreshToken: "refresh.acces.true",
				AccessToken:  "access.ShortL.true",
				Login:        "ShortL",
			},
			wantReturned: iomodels.RefreshTokensReturned{
				AccessToken:  "",
				RefreshToken: "",
				Err:          casesErrs.ErrTooShortLoginOrPassword{},
			},
		},
		{
			name: "invalid_refresh1",
			args: iomodels.RefreshTokensArgs{
				Ctx:          context.Background(),
				RefreshToken: "access.acces.true",
				AccessToken:  "access.SampleLogin.true",
				Login:        "SampleLogin",
			},
			wantReturned: iomodels.RefreshTokensReturned{
				AccessToken:  "",
				RefreshToken: "",
				Err:          logicErrs.ErrInvalidRefreshToken{},
			},
		},
		{
			name: "expired_refresh1",
			args: iomodels.RefreshTokensArgs{
				Ctx:          context.Background(),
				RefreshToken: "refresh.acces.false",
				AccessToken:  "access.SampleLogin.true",
				Login:        "SampleLogin",
			},
			wantReturned: iomodels.RefreshTokensReturned{
				AccessToken:  "",
				RefreshToken: "",
				Err:          logicErrs.ErrExpiredRefreshToken{},
			},
		},
		{
			name: "invalid_refresh2",
			args: iomodels.RefreshTokensArgs{
				Ctx:          context.Background(),
				RefreshToken: "refresh.acce.true",
				AccessToken:  "access.SampleLogin.true",
				Login:        "SampleLogin",
			},
			wantReturned: iomodels.RefreshTokensReturned{
				AccessToken:  "",
				RefreshToken: "",
				Err:          logicErrs.ErrInvalidRefreshToken{},
			},
		},
		{
			name: "invalid_access1",
			args: iomodels.RefreshTokensArgs{
				Ctx:          context.Background(),
				RefreshToken: "refresh.refre.true",
				AccessToken:  "refresh.SampleLogin.true",
				Login:        "SampleLogin",
			},
			wantReturned: iomodels.RefreshTokensReturned{
				AccessToken:  "",
				RefreshToken: "",
				Err:          logicErrs.ErrInvalidAccessToken{},
			},
		},
		{
			name: "invalid_access2",
			args: iomodels.RefreshTokensArgs{
				Ctx:          context.Background(),
				RefreshToken: "refresh.refre.true",
				AccessToken:  "access.SampleLogin.true",
				Login:        "SampleLogin2",
			},
			wantReturned: iomodels.RefreshTokensReturned{
				AccessToken:  "",
				RefreshToken: "",
				Err:          logicErrs.ErrInvalidAccessToken{},
			},
		},
		{
			name: "user_not_exists",
			args: iomodels.RefreshTokensArgs{
				Ctx:          context.Background(),
				RefreshToken: "refresh.refre.true",
				AccessToken:  "access.SampleLogin3.true",
				Login:        "SampleLogin3",
			},
			wantReturned: iomodels.RefreshTokensReturned{
				AccessToken:  "",
				RefreshToken: "",
				Err:          logicErrs.ErrUserNotExists{},
			},
		},
	}
	for _, tCase := range tests {
		t.Run(tCase.name, func(t *testing.T) {
			returned := casesProvider.RefreshTokens(tCase.args)
			assert.Equal(t, tCase.wantReturned, returned)
		})
	}
}
