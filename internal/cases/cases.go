package cases

import (
	"auth/internal/cases/dependencies"
	"auth/internal/cases/errs"
	"auth/internal/cases/iomodels"
	"auth/internal/logic"
	logicErrs "auth/internal/logic/errs"
	logicIomodels "auth/internal/logic/iomodels"

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

func (c *CasesProvider) RegisterUser(args iomodels.RegisterUserArgs) (returned iomodels.RegisterUserReturned) {
	select {
	case <-args.Ctx.Done():
		{
			returned.Err = errs.ErrServiceNotAvaliable{}
			return
		}
	default:
		{
			if len(args.User.Login) < c.config.GetMinLoginLen() || len(args.User.Password) < c.config.GetMinPasswordLen() {
				returned.Err = errs.ErrTooShortLoginOrPassword{}
				return
			}

			var registerUserArgs logicIomodels.RegisterUserArgs = logicIomodels.RegisterUserArgs{
				Ctx:  args.Ctx,
				User: args.User,
			}
			registerUserReturned := c.logic.RegisterUser(registerUserArgs)

			if registerUserReturned.Err == (logicErrs.ErrUserAlreadyExists{}) {
				returned.Err = registerUserReturned.Err
				return
			} else if registerUserReturned.Err != nil {
				// ToDo: add tracing by token
				c.logger.Error().Msg("Internal register user error: " + registerUserReturned.Err.Error())
				returned.Err = errs.ErrServiceInternal{}
				return
			}
			return
		}
	}
}

func (c *CasesProvider) AuthenticateUserByLogin(args iomodels.AuthenticateUserByLoginArgs) (returned iomodels.AuthenticateUserByLoginReturned) {
	select {
	case <-args.Ctx.Done():
		{
			returned.Err = errs.ErrServiceNotAvaliable{}
			return
		}
	default:
		{
			if len(args.Login) < c.config.GetMinLoginLen() || len(args.Password) < c.config.GetMinPasswordLen() {
				returned.Err = errs.ErrTooShortLoginOrPassword{}
				return
			}

			var authenticateUserByLoginArgs logicIomodels.AuthenticateUserByLoginArgs = logicIomodels.AuthenticateUserByLoginArgs{
				Ctx:      args.Ctx,
				Login:    args.Login,
				Password: args.Password,
			}

			authenticateUserByLoginReturned := c.logic.AuthenticateUserByLogin(authenticateUserByLoginArgs)

			if authenticateUserByLoginReturned.Err == (logicErrs.ErrUserNotExists{}) || authenticateUserByLoginReturned.Err == (logicErrs.ErrInvalidPassword{}) {
				returned.Err = errs.ErrInvalidLoginOrPassword{}
				return
			} else if authenticateUserByLoginReturned.Err != nil {
				// ToDo: add tracing by token
				c.logger.Error().Msg("Internal authenticate user by login error: " + authenticateUserByLoginReturned.Err.Error())
				returned.Err = errs.ErrServiceInternal{}
				return
			}
			if authenticateUserByLoginReturned.OtpEnabled {
				returned.IntermediateToken = authenticateUserByLoginReturned.IntermediateToken
				returned.OtpEnabled = authenticateUserByLoginReturned.OtpEnabled
			}
			returned.AuthInfo.AccessToken = authenticateUserByLoginReturned.AuthInfo.AccessToken
			returned.AuthInfo.RefreshToken = authenticateUserByLoginReturned.AuthInfo.RefreshToken
			return
		}
	}
}

func (c *CasesProvider) AuthorizeUser(args iomodels.AuthorizeUserArgs) (returned iomodels.AuthorizeUserReturned) {
	select {
	case <-args.Ctx.Done():
		{
			returned.Err = errs.ErrServiceNotAvaliable{}
			return
		}
	default:
		{
			if len(args.Login) < c.config.GetMinLoginLen() {
				returned.Err = errs.ErrTooShortLoginOrPassword{}
				return
			}

			var authorizeUserArgs logicIomodels.AuthorizeUserArgs = logicIomodels.AuthorizeUserArgs{
				Ctx:         args.Ctx,
				AccessToken: args.AccessToken,
				Login:       args.Login,
			}

			authorizeUserReturned := c.logic.AuthorizeUser(authorizeUserArgs)

			if authorizeUserReturned.Err == (logicErrs.ErrUserNotExists{}) || authorizeUserReturned.Err == (logicErrs.ErrExpiredAccessToken{}) || authorizeUserReturned.Err == (logicErrs.ErrInvalidAccessToken{}) {
				returned.Err = authorizeUserReturned.Err
				return
			} else if authorizeUserReturned.Err != nil {
				c.logger.Error().Msg("Internal authorize user error: " + authorizeUserReturned.Err.Error())
				returned.Err = errs.ErrServiceInternal{}
				return
			}
			returned.UserId = authorizeUserReturned.UserId
			return
		}
	}
}

func (c *CasesProvider) RefreshTokens(args iomodels.RefreshTokensArgs) (returned iomodels.RefreshTokensReturned) {
	select {
	case <-args.Ctx.Done():
		{
			returned.Err = errs.ErrServiceNotAvaliable{}
			return
		}
	default:
		{
			if len(args.Login) < c.config.GetMinLoginLen() {
				returned.Err = errs.ErrTooShortLoginOrPassword{}
				return
			}

			var refreshTokensArgs logicIomodels.RefreshTokensArgs = logicIomodels.RefreshTokensArgs{
				Ctx:          args.Ctx,
				AccessToken:  args.AccessToken,
				RefreshToken: args.RefreshToken,
				Login:        args.Login,
			}

			refreshTokensReturned := c.logic.RefreshTokens(refreshTokensArgs)

			if refreshTokensReturned.Err == (logicErrs.ErrExpiredRefreshToken{}) || refreshTokensReturned.Err == (logicErrs.ErrInvalidRefreshToken{}) || refreshTokensReturned.Err == (logicErrs.ErrUserNotExists{}) || refreshTokensReturned.Err == (logicErrs.ErrInvalidAccessToken{}) {
				returned.Err = refreshTokensReturned.Err
				return
			} else if refreshTokensReturned.Err != nil {
				returned.Err = errs.ErrServiceInternal{}
				c.logger.Error().Msg("Internal authorize user error: " + refreshTokensReturned.Err.Error())
				return
			}
			returned.AccessToken = refreshTokensReturned.AccessToken
			returned.RefreshToken = refreshTokensReturned.RefreshToken
			return
		}
	}
}
