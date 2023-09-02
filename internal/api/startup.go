package api

import (
	"auth/internal/api/dependencies"
	"auth/internal/api/handlers"
	"context"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
)

func Startup(ctx context.Context, handlerProvider *handlers.HandlersProvider, apiConfig dependencies.ApiConfig, logger zerolog.Logger) {
	app := fiber.New()
	app.Get("/authenticate", handlerProvider.AuthenticateUserByLoginHandler())
	app.Post("register", handlerProvider.RegisterUserHandler())
	app.Get("/authorize", handlerProvider.AuthorizeUserHandler())
	go func() {
		app.ShutdownWithTimeout(time.Minute)
	}()

	logger.Fatal().Msg(app.Listen(":" + apiConfig.GetApiPort()).Error())
}
