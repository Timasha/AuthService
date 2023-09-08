package cases_test

import (
	"auth/internal/cases"
	casesErrs "auth/internal/cases/errs"
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

	type args struct {
		ctx  context.Context
		user models.User
	}
	tests := []struct {
		name    string
		args    args
		wantErr error
	}{
		{
			name: "success1",
			args: args{
				ctx: context.Background(),
				user: models.User{
					Login:    "SampleLogin3",
					Password: "SamplePassword3",
				},
			},
			wantErr: nil,
		},
		{
			name: "user_already_exists",
			args: args{
				ctx: context.Background(),
				user: models.User{
					Login:    "SampleLogin",
					Password: "SamplePasswordDiff",
				},
			},
			wantErr: logicErrs.ErrUserAlreadyExists{},
		},
		{
			name: "too_short_login",
			args: args{
				ctx: context.Background(),
				user: models.User{
					Login:    "asd",
					Password: "asdasdasdasd",
				},
			},
			wantErr: casesErrs.ErrTooShortLoginOrPassword{},
		},
		{
			name: "too_short_password",
			args: args{
				ctx: context.Background(),
				user: models.User{
					Login:    "asdasdasdasd",
					Password: "asd",
				},
			},
			wantErr: casesErrs.ErrTooShortLoginOrPassword{},
		},
	}
	for _, tCase := range tests {
		t.Run(tCase.name, func(t *testing.T) {
			err := casesProvider.RegisterUser(tCase.args.ctx, tCase.args.user)
			assert.Equal(t, tCase.wantErr, err)
		})
	}
}

func TestCasesProvider_AuthenticateUserByLogin(t *testing.T) {
	casesProvider := CasesMocks(t)

	type args struct {
		ctx      context.Context
		login    string
		password string
	}
	tests := []struct {
		name        string
		args        args
		wantAccess  string
		wantRefresh string
		wantErr     error
	}{
		{
			name: "success1",
			args: args{
				ctx:      context.Background(),
				login:    "SampleLogin",
				password: "SamplePassword",
			},
			wantAccess:  "access.SampleLogin.true",
			wantRefresh: "refresh.acces.true",
			wantErr:     nil,
		},
		{
			name: "success2",
			args: args{
				ctx:      context.Background(),
				login:    "SampleLogin2",
				password: "SamplePassword2",
			},
			wantAccess:  "access.SampleLogin2.true",
			wantRefresh: "refresh.acces.true",
			wantErr:     nil,
		},
		{
			name: "user_not_exists",
			args: args{
				ctx:      context.Background(),
				login:    "SomeLoginExample",
				password: "SomePassword",
			},
			wantAccess:  "",
			wantRefresh: "",
			wantErr:     casesErrs.ErrInvalidLoginOrPassword{},
		},
		{
			name: "too_short_login",
			args: args{
				ctx:      context.Background(),
				login:    "ShortL",
				password: "SamplePassword",
			},
			wantAccess:  "",
			wantRefresh: "",
			wantErr:     casesErrs.ErrTooShortLoginOrPassword{},
		},
		{
			name: "too_short_password",
			args: args{
				ctx:      context.Background(),
				login:    "SampleLogin",
				password: "ShortP",
			},
			wantAccess:  "",
			wantRefresh: "",
			wantErr:     casesErrs.ErrTooShortLoginOrPassword{},
		},
	}
	for _, tCase := range tests {
		t.Run(tCase.name, func(t *testing.T) {
			accessToken, refreshToken, err := casesProvider.AuthenticateUserByLogin(tCase.args.ctx, tCase.args.login, tCase.args.password)
			assert.Equal(t, tCase.wantAccess, accessToken)
			assert.Equal(t, tCase.wantRefresh, refreshToken)
			assert.Equal(t, tCase.wantErr, err)
		})
	}
}

func TestCasesProvider_AuthorizeUser(t *testing.T) {

	casesProvider := CasesMocks(t)

	type args struct {
		ctx         context.Context
		accessToken string
		login       string
	}
	tests := []struct {
		name    string
		args    args
		wantUuid string
		wantErr error
	}{
		{
			name: "success1",
			args: args{
				ctx:         context.Background(),
				accessToken: "access.SampleLogin.true",
				login:       "SampleLogin",
			},
			wantUuid: "Some_UUID",
			wantErr: nil,
		},
		{
			name: "invalid_token1",
			args: args{
				ctx:         context.Background(),
				accessToken: "refresh.SampleLogin.true",
				login:       "SampleLogin",
			},
			wantErr: logicErrs.ErrInvalidAccessToken{},
		},
		{
			name: "invalid_token2",
			args: args{
				ctx:         context.Background(),
				accessToken: "access.SampleLogin2.true",
				login:       "SampleLogin",
			},
			wantErr: logicErrs.ErrInvalidAccessToken{},
		},
		{
			name: "not_existing_user1",
			args: args{
				ctx:         context.Background(),
				accessToken: "access.SampleLogin.true",
				login:       "SampleNotExistingLogin",
			},
			wantErr: logicErrs.ErrUserNotExists{},
		},
		{
			name: "too_short_login",
			args: args{
				ctx:         context.Background(),
				accessToken: "access.ShortL.true",
				login:       "ShortL",
			},
			wantErr: casesErrs.ErrTooShortLoginOrPassword{},
		},
		{
			name: "expired_token1",
			args: args{
				ctx:         context.Background(),
				accessToken: "access.SampleLogin.false",
				login:       "SampleLogin",
			},
			wantErr: logicErrs.ErrExpiredAccessToken{},
		},
	}
	for _, tCase := range tests {
		t.Run(tCase.name, func(t *testing.T) {
			uuid, err := casesProvider.AuthorizeUser(tCase.args.ctx, tCase.args.accessToken, tCase.args.login)
			assert.Equal(t, tCase.wantErr, err)
			assert.Equal(t,tCase.wantUuid,uuid)
		})
	}
}

func TestCasesProvider_RefreshTokens(t *testing.T) {
	casesProvider := CasesMocks(t)
	type args struct {
		ctx          context.Context
		refreshToken string
		accessToken  string
		login        string
	}
	tests := []struct {
		name        string
		args        args
		wantAccess  string
		wantRefresh string
		wantErr     error
	}{
		{
			name: "success1",
			args: args{
				ctx:          context.Background(),
				refreshToken: "refresh.acces.true",
				accessToken:  "access.SampleLogin.true",
				login:        "SampleLogin",
			},
			wantAccess:  "access.SampleLogin.true",
			wantRefresh: "refresh.acces.true",
			wantErr:     nil,
		},
		{
			name: "success2",
			args: args{
				ctx:          context.Background(),
				refreshToken: "refresh.acces.true",
				accessToken:  "access.SampleLogin.false",
				login:        "SampleLogin",
			},
			wantAccess:  "access.SampleLogin.true",
			wantRefresh: "refresh.acces.true",
			wantErr:     nil,
		},
		{
			name: "too_short_login",
			args: args{
				ctx:          context.Background(),
				refreshToken: "refresh.acces.true",
				accessToken:  "access.ShortL.true",
				login:        "ShortL",
			},
			wantAccess:  "",
			wantRefresh: "",
			wantErr:     casesErrs.ErrTooShortLoginOrPassword{},
		},
		{
			name: "invalid_refresh1",
			args: args{
				ctx:          context.Background(),
				refreshToken: "access.acces.true",
				accessToken:  "access.SampleLogin.true",
				login:        "SampleLogin",
			},
			wantAccess:  "",
			wantRefresh: "",
			wantErr:     logicErrs.ErrInvalidRefreshToken{},
		},
		{
			name: "expired_refresh1",
			args: args{
				ctx:          context.Background(),
				refreshToken: "refresh.acces.false",
				accessToken:  "access.SampleLogin.true",
				login:        "SampleLogin",
			},
			wantAccess:  "",
			wantRefresh: "",
			wantErr:     logicErrs.ErrExpiredRefreshToken{},
		},
		{
			name: "invalid_refresh2",
			args: args{
				ctx:          context.Background(),
				refreshToken: "refresh.acce.true",
				accessToken:  "access.SampleLogin.true",
				login:        "SampleLogin",
			},
			wantAccess:  "",
			wantRefresh: "",
			wantErr:     logicErrs.ErrInvalidRefreshToken{},
		},
		{
			name: "invalid_access1",
			args: args{
				ctx:          context.Background(),
				refreshToken: "refresh.refre.true",
				accessToken:  "refresh.SampleLogin.true",
				login:        "SampleLogin",
			},
			wantAccess:  "",
			wantRefresh: "",
			wantErr:     logicErrs.ErrInvalidAccessToken{},
		},
		{
			name: "invalid_access2",
			args: args{
				ctx:          context.Background(),
				refreshToken: "refresh.refre.true",
				accessToken:  "access.SampleLogin.true",
				login:        "SampleLogin2",
			},
			wantAccess:  "",
			wantRefresh: "",
			wantErr:     logicErrs.ErrInvalidAccessToken{},
		},
		{
			name: "user_not_exists",
			args: args{
				ctx:          context.Background(),
				refreshToken: "refresh.refre.true",
				accessToken:  "access.SampleLogin3.true",
				login:        "SampleLogin3",
			},
			wantAccess:  "",
			wantRefresh: "",
			wantErr:     logicErrs.ErrUserNotExists{},
		},
	}
	for _, tCase := range tests {
		t.Run(tCase.name, func(t *testing.T) {
			access, refresh, err := casesProvider.RefreshTokens(tCase.args.ctx, tCase.args.refreshToken, tCase.args.accessToken, tCase.args.login)
			assert.Equal(t, tCase.wantAccess, access)
			assert.Equal(t, tCase.wantRefresh, refresh)
			assert.Equal(t, tCase.wantErr, err)
		})
	}
}
