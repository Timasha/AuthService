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
	//bodySerializer BodySerializer
	logger logger.Logger
}

func New(ctx context.Context, casesProvider *cases.CasesProvider, apiConfig ApiConfig /*bodySerializer BodySerializer,*/, logger logger.Logger) (a *Auth) {
	a = &Auth{
		ctx:           ctx,
		casesProvider: casesProvider,
		apiConfig:     apiConfig,
		//bodySerializer: bodySerializer,
		logger: logger,
	}
	return
}

func (a *Auth) Start() {
	app := fiber.New()
	app.Group("/", a.GetJsonMiddleware())
	app.Post("/authenticate", a.GetAuthenticateUserByLoginHandler())
	app.Post("/register", a.GetRegisterUserHandler())
	app.Post("/authorize", a.GetAuthorizeUserHandler())
	app.Post("/refresh", a.GetRefreshTokensHandler())
	defer func() {
		app.ShutdownWithTimeout(time.Minute)
	}()

	a.logger.Log(logger.LogMsg{
		Time:     time.Now(),
		LogLevel: logger.LogLevelFatal,
		Msg:      app.Listen(":" + a.apiConfig.GetApiPort()).Error(),
	})
}

var ErrWrongAuthMethod errsutil.AuthErr = errsutil.AuthErr{
	Msg:     "wrong auth method",
	ErrCode: errsutil.ErrWrongAuthMethodCode,
}

type ApiConfig interface {
	GetApiPort() string
}

// type BodySerializer interface {
// 	Unmarshal[T any](data []byte, serializableObject T) error
// 	Marshal[T](serializableObject T) ([]byte, error)
// }

type BaseResponse struct {
	Err     string               `json:"error"`
	ErrCode errsutil.AuthErrCode `json:"errorCode"`
}
