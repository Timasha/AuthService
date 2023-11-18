package api

import (
	"AuthService/internal/cases"
	"AuthService/internal/utils/errsutil"
	"AuthService/internal/utils/logger"
	"context"
	"time"

	"github.com/gofiber/fiber/v2"
)

type Auth struct {
	ctx context.Context

	casesProvider *cases.CasesProvider
	apiConfig     ApiConfig
	bodySerializer BodySerializer
	logger logger.Logger
}

func New(ctx context.Context, casesProvider *cases.CasesProvider, apiConfig ApiConfig,bodySerializer BodySerializer, logger logger.Logger) (a *Auth) {
	a = &Auth{
		ctx:           ctx,
		casesProvider: casesProvider,
		apiConfig:     apiConfig,
		bodySerializer: bodySerializer,
		logger: logger,
	}
	return
}

func (a *Auth) Start() {
	app := fiber.New()
	app.Group("/", a.GetJsonMiddleware())
	app.Post("/authenticate", a.GetAuthenticateUserByLoginHandler())
	app.Post("/authenticate/continue", a.GetContinueAuthenticateOtpUserByLoginHandler())
	app.Post("/register", a.GetRegisterUserHandler())
	app.Post("/authorize", a.GetAuthorizeUserHandler())
	app.Post("/refresh", a.GetRefreshTokensHandler())

	app.Group("/otp",a.GetAuthorizeMiddleware())

	app.Post("/otp/enable",a.GetEnableOtpAuthenticationForUserHandler())
	app.Post("/otp/disable",a.GetDisableOtpAuthenticationForUserHandler())
	defer func() {
		<-a.ctx.Done()
		app.Shutdown()
	}()

	a.logger.Log(logger.LogMsg{
		Time:     time.Now(),
		LogLevel: logger.LogLevelFatal,
		Msg:      app.Listen(":" + a.apiConfig.GetApiPort()).Error(),
	})
}

var (
	ErrWrongAuthorizationMethod errsutil.AuthErr = errsutil.AuthErr{
		Msg:     "wrong authorization method",
		ErrCode: errsutil.ErrWrongAuthorizeMethodCode,
	}
	ErrInvalidInput  errsutil.AuthErr = errsutil.AuthErr{
		Msg: "cannot read input: ",
		ErrCode: errsutil.ErrInvalidInputCode,
	}
)

type ApiConfig interface {
	GetApiPort() string
}

type BodySerializer interface {
	Unmarshal(data []byte, serializableObject interface{}) error
	Marshal(serializableObject interface{}) ([]byte, error)
}

type BaseResponse struct {
	Err     string               `json:"error"`
	ErrCode errsutil.AuthErrCode `json:"errorCode"`
}
