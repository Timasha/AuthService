package cases_test

import (
	"AuthService/internal/cases"
	"AuthService/internal/logic"
	"AuthService/internal/logic/models"
	"AuthService/internal/utils/config"
	"AuthService/internal/utils/logger/logdrivers"
	"AuthService/internal/utils/password"
	"AuthService/internal/utils/uuid"
	"context"
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func CasesMocks(t *testing.T) *cases.CasesProvider {
	var logicProvider *logic.LogicProvider
	var userStorage logic.UserStorage = make(cases.UserStorageMock)
	var rolesStorage logic.RolesStorage = make(cases.RolesStorageMock)

	hasher := &password.BcryptPasswordHasher{}
	hashedPassword, err := hasher.Hash("SamplePassword")

	if err != nil {
		t.Fatalf("Cannot hash password: %v", err)
	}

	userStorage.CreateUser(context.Background(), models.User{
		UserID:   "Some_UUID",
		Login:    "SampleLogin",
		Password: hashedPassword,
		Role: models.Role{
			RoleId:   1,
			RoleName: "user",
		},
	})
	hashedPassword, err = hasher.Hash("SamplePassword2")
	userStorage.CreateUser(context.Background(), models.User{
		UserID:   "Some_UUID2",
		Login:    "SampleLogin2",
		Password: hashedPassword,
		Role: models.Role{
			RoleId:   1,
			RoleName: "user",
		},
	})

	logicProvider = logic.New(userStorage, rolesStorage, &cases.TokensProviderMock{}, &password.BcryptPasswordHasher{}, &uuid.GoogleUUIDProvider{}, &cases.OtpGeneratorMock{})

	var casesProvider *cases.CasesProvider = &cases.CasesProvider{}

	casesProvider = cases.New(&config.JSONConfig{
		MinLoginLen:    10,
		MinPasswordLen: 10,
	}, logdrivers.NewZerologDriver([]io.Writer{os.Stdout}), logicProvider)
	return casesProvider
}

func TestCasesProvider_RegisterUser(t *testing.T) {
	casesProvider := CasesMocks(t)

	tests := []struct {
		name         string
		args         cases.RegisterUserArgs
		wantReturned cases.RegisterUserReturned
	}{
		{
			name: "success1",
			args: cases.RegisterUserArgs{
				Ctx: context.Background(),
				User: models.User{
					Login:    "SampleLogin3",
					Password: "SamplePassword3",
				},
			},
			wantReturned: cases.RegisterUserReturned{
				Err: nil,
			},
		},
		{
			name: "user_already_exists",
			args: cases.RegisterUserArgs{
				Ctx: context.Background(),
				User: models.User{
					Login:    "SampleLogin",
					Password: "SamplePasswordDiff",
				},
			},
			wantReturned: cases.RegisterUserReturned{
				Err: logic.ErrUserAlreadyExists,
			},
		},
		{
			name: "too_short_login",
			args: cases.RegisterUserArgs{
				Ctx: context.Background(),
				User: models.User{
					Login:    "asd",
					Password: "asdasdasdasd",
				},
			},
			wantReturned: cases.RegisterUserReturned{
				Err: cases.ErrTooShortLoginOrPassword,
			},
		},
		{
			name: "too_short_password",
			args: cases.RegisterUserArgs{
				Ctx: context.Background(),
				User: models.User{
					Login:    "asdasdasdasd",
					Password: "asd",
				},
			},
			wantReturned: cases.RegisterUserReturned{
				Err: cases.ErrTooShortLoginOrPassword,
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
		args         cases.AuthenticateUserByLoginArgs
		wantReturned cases.AuthenticateUserByLoginReturned
	}{
		{
			name: "success1",
			args: cases.AuthenticateUserByLoginArgs{
				Ctx:      context.Background(),
				Login:    "SampleLogin",
				Password: "SamplePassword",
			},
			wantReturned: cases.AuthenticateUserByLoginReturned{
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
			args: cases.AuthenticateUserByLoginArgs{
				Ctx:      context.Background(),
				Login:    "SampleLogin2",
				Password: "SamplePassword2",
			},
			wantReturned: cases.AuthenticateUserByLoginReturned{
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
			args: cases.AuthenticateUserByLoginArgs{
				Ctx:      context.Background(),
				Login:    "SomeLoginExample",
				Password: "SomePassword",
			},
			wantReturned: cases.AuthenticateUserByLoginReturned{
				OtpEnabled:        false,
				IntermediateToken: "",

				AuthInfo: struct {
					AccessToken  string
					RefreshToken string
				}{
					AccessToken:  "",
					RefreshToken: "",
				},

				Err: cases.ErrInvalidLoginOrPassword,
			},
		},
		{
			name: "too_short_login",
			args: cases.AuthenticateUserByLoginArgs{
				Ctx:      context.Background(),
				Login:    "ShortL",
				Password: "SamplePassword",
			},
			wantReturned: cases.AuthenticateUserByLoginReturned{
				OtpEnabled:        false,
				IntermediateToken: "",

				AuthInfo: struct {
					AccessToken  string
					RefreshToken string
				}{
					AccessToken:  "",
					RefreshToken: "",
				},

				Err: cases.ErrTooShortLoginOrPassword,
			},
		},
		{
			name: "too_short_password",
			args: cases.AuthenticateUserByLoginArgs{
				Ctx:      context.Background(),
				Login:    "SampleLogin",
				Password: "ShortP",
			},
			wantReturned: cases.AuthenticateUserByLoginReturned{
				OtpEnabled:        false,
				IntermediateToken: "",

				AuthInfo: struct {
					AccessToken  string
					RefreshToken string
				}{
					AccessToken:  "",
					RefreshToken: "",
				},

				Err: cases.ErrTooShortLoginOrPassword,
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
		args         cases.AuthorizeUserArgs
		wantReturned cases.AuthorizeUserReturned
	}{
		{
			name: "success1",

			args: cases.AuthorizeUserArgs{
				Ctx:         context.Background(),
				AccessToken: "access.SampleLogin.true",
			},
			wantReturned: cases.AuthorizeUserReturned{
				UserId: "Some_UUID",
				Err:    nil,
			},
		},
		{
			name: "invalid_token1",
			args: cases.AuthorizeUserArgs{
				Ctx:         context.Background(),
				AccessToken: "refresh.SampleLogin.true",
			},
			wantReturned: cases.AuthorizeUserReturned{
				UserId: "",
				Err:    logic.ErrInvalidAccessToken,
			},
		},
		{
			name: "not_existing_user1",
			args: cases.AuthorizeUserArgs{
				Ctx:         context.Background(),
				AccessToken: "access.SampleLogin3w.true",
			},
			wantReturned: cases.AuthorizeUserReturned{
				UserId: "",
				Err:    logic.ErrUserNotExists,
			},
		},
		{
			name: "expired_token1",
			args: cases.AuthorizeUserArgs{
				Ctx:         context.Background(),
				AccessToken: "access.SampleLogin.false",
			},
			wantReturned: cases.AuthorizeUserReturned{
				UserId: "",
				Err:    logic.ErrExpiredAccessToken,
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
		args         cases.RefreshTokensArgs
		wantReturned cases.RefreshTokensReturned
	}{
		{
			name: "success1",
			args: cases.RefreshTokensArgs{
				Ctx:          context.Background(),
				RefreshToken: "refresh.acces.true",
				AccessToken:  "access.SampleLogin.true",
			},
			wantReturned: cases.RefreshTokensReturned{
				AccessToken:  "access.SampleLogin.true",
				RefreshToken: "refresh.acces.true",
				Err:          nil,
			},
		},
		{
			name: "success2",
			args: cases.RefreshTokensArgs{
				Ctx:          context.Background(),
				RefreshToken: "refresh.acces.true",
				AccessToken:  "access.SampleLogin.false",
			},
			wantReturned: cases.RefreshTokensReturned{
				AccessToken:  "access.SampleLogin.true",
				RefreshToken: "refresh.acces.true",
				Err:          nil,
			},
		},
		{
			name: "invalid_refresh1",
			args: cases.RefreshTokensArgs{
				Ctx:          context.Background(),
				RefreshToken: "access.acces.true",
				AccessToken:  "access.SampleLogin.true",
			},
			wantReturned: cases.RefreshTokensReturned{
				AccessToken:  "",
				RefreshToken: "",
				Err:          logic.ErrInvalidRefreshToken,
			},
		},
		{
			name: "expired_refresh1",
			args: cases.RefreshTokensArgs{
				Ctx:          context.Background(),
				RefreshToken: "refresh.acces.false",
				AccessToken:  "access.SampleLogin.true",
			},
			wantReturned: cases.RefreshTokensReturned{
				AccessToken:  "",
				RefreshToken: "",
				Err:          logic.ErrExpiredRefreshToken,
			},
		},
		{
			name: "invalid_refresh2",
			args: cases.RefreshTokensArgs{
				Ctx:          context.Background(),
				RefreshToken: "refresh.acce.true",
				AccessToken:  "access.SampleLogin.true",
			},
			wantReturned: cases.RefreshTokensReturned{
				AccessToken:  "",
				RefreshToken: "",
				Err:          logic.ErrInvalidRefreshToken,
			},
		},
		{
			name: "invalid_access1",
			args: cases.RefreshTokensArgs{
				Ctx:          context.Background(),
				RefreshToken: "refresh.refre.true",
				AccessToken:  "refresh.SampleLogin.true",
			},
			wantReturned: cases.RefreshTokensReturned{
				AccessToken:  "",
				RefreshToken: "",
				Err:          logic.ErrInvalidAccessToken,
			},
		},
		{
			name: "user_not_exists",
			args: cases.RefreshTokensArgs{
				Ctx:          context.Background(),
				RefreshToken: "refresh.refre.true",
				AccessToken:  "access.SampleLogin3.true",
			},
			wantReturned: cases.RefreshTokensReturned{
				AccessToken:  "",
				RefreshToken: "",
				Err:          logic.ErrUserNotExists,
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
