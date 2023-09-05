package cases

import (
	"auth/internal/cases/dependencies"
	"auth/internal/cases/errs"
	"auth/internal/logic"
	logicErrs "auth/internal/logic/errs"

	"auth/internal/logic/models"
	"context"

	"github.com/rs/zerolog"
)

type CasesProvider struct {
	config dependencies.UseCasesConfig
	logger zerolog.Logger

	logic *logic.LogicProvider
}

func (c *CasesProvider) Init(config dependencies.UseCasesConfig, logger zerolog.Logger, logic *logic.LogicProvider) {
	c.config = config
	c.logger = logger
	c.logic = logic
}


func (c *CasesProvider) RegisterUser(ctx context.Context, user models.User) error {
	select {
	case <-ctx.Done():
		{
			return errs.ErrServiceNotAvaliable{}
		}
	default:
		{
			if len(user.Login) < c.config.GetMinLoginLen() || len(user.Password) < c.config.GetMinPasswordLen() {
				return errs.ErrTooShortLoginOrPassword{}
			}

			regErr := c.logic.RegisterUser(ctx, user)

			if regErr == (logicErrs.ErrUserAlreadyExists{}) {
				return regErr
			} else if regErr != nil {
				// ToDo: add tracing by token
				c.logger.Error().Msg("Internal register user error: " + regErr.Error())
				return errs.ErrServiceInternal{}
			}
			return nil
		}
	}
}

func (c *CasesProvider) AuthenticateUserByLogin(ctx context.Context, login, password string) (string, string, error) {
	select {
	case <-ctx.Done():
		{
			return "", "", errs.ErrServiceNotAvaliable{}
		}
	default:
		{
			if len(login) < c.config.GetMinLoginLen() || len(password) < c.config.GetMinPasswordLen() {
				return "", "", errs.ErrTooShortLoginOrPassword{}
			}

			accessToken, refreshToken, err := c.logic.AuthenticateUserByLogin(ctx, login, password)

			if err == (logicErrs.ErrUserNotExists{}) || err == (logicErrs.ErrInvalidPassword{}) {
				return "", "", errs.ErrInvalidLoginOrPassword{}
			} else if err != nil {
				// ToDo: add tracing by token
				c.logger.Error().Msg("Internal authenticate user by login error: " + err.Error())
				return "", "", errs.ErrServiceInternal{}
			}
			return accessToken, refreshToken, nil
		}
	}
}

func (c *CasesProvider) AuthorizeUser(ctx context.Context, accessToken, login string) error {
	select {
	case <-ctx.Done():
		{
			return errs.ErrServiceNotAvaliable{}
		}
	default:
		{
			if len(login) < c.config.GetMinLoginLen() {
				return errs.ErrTooShortLoginOrPassword{}
			}
			err := c.logic.AuthorizeUser(ctx, accessToken, login)

			if err == (logicErrs.ErrUserNotExists{}) || err == (logicErrs.ErrExpiredAccessToken{}) || err == (logicErrs.ErrInvalidAccessToken{}) {
				return err
			} else if err != nil {
				c.logger.Error().Msg("Internal authorize user error: " + err.Error())
				return errs.ErrServiceInternal{}
			}
			return nil
		}
	}
}

func (c *CasesProvider) RefreshTokens(ctx context.Context, refreshToken, accessToken, login string) (string, string, error) {
	select {
	case <-ctx.Done():
		{
			return "", "", errs.ErrServiceNotAvaliable{}
		}
	default:
		{
			if len(login) < c.config.GetMinLoginLen() {
				return "", "", errs.ErrTooShortLoginOrPassword{}
			}

			newAccess, newRefresh, err := c.logic.RefreshTokens(ctx, accessToken, refreshToken, login)

			if err == (logicErrs.ErrExpiredRefreshToken{}) || err == (logicErrs.ErrInvalidRefreshToken{}) || err == (logicErrs.ErrUserNotExists{}) || err == (logicErrs.ErrInvalidAccessToken{}) {
				return "", "", err
			} else if err != nil {
				c.logger.Error().Msg("Internal authorize user error: " + err.Error())
				return "", "", errs.ErrServiceInternal{}
			}
			return newAccess, newRefresh, nil
		}
	}
}
