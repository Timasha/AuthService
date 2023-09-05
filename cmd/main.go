package main

import (
	"auth/internal/api"
	"auth/internal/api/handlers"
	"auth/internal/cases"
	"auth/internal/dependencies/config"
	"auth/internal/dependencies/password"
	"auth/internal/dependencies/storage"
	"auth/internal/dependencies/tokens/jwt"
	"auth/internal/dependencies/uuid"
	"auth/internal/logic"
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/rs/zerolog"
)

func main() {
	var (
		logger    zerolog.Logger
		appConfig *config.JSONConfig

		userStorage    *storage.PostgresUserStorage   = &storage.PostgresUserStorage{}
		tokensProvider *jwt.TokensProvider            = &jwt.TokensProvider{}
		passwordHasher *password.BcryptPasswordHasher = &password.BcryptPasswordHasher{}
		uuidProvider   *uuid.GoogleUUIDProvider       = &uuid.GoogleUUIDProvider{}

		logicProvider *logic.LogicProvider = &logic.LogicProvider{}

		casesProvider *cases.CasesProvider = &cases.CasesProvider{}

		handlersProvider *handlers.HandlersProvider = &handlers.HandlersProvider{}
	)
	logger = zerolog.New(os.Stdout)

	appConfig, readConfigErr := config.ReadJsonConfig("./config.json", logger)
	if readConfigErr != nil {
		logger.Fatal().Msg("Read config error: " + readConfigErr.Error())
	}



	connectStorageContext, connectStorageClose := context.WithTimeout(context.Background(), time.Minute)
	defer connectStorageClose()

	userStorage, connectErr := storage.Connect(connectStorageContext, appConfig.PostgresConfig)

	if connectErr != nil {
		logger.Fatal().Msg("Read config error: " + connectErr.Error())
	}



	migrateCtx, migrateClose := context.WithTimeout(context.Background(), time.Minute)
	defer migrateClose()

	migrateErr := userStorage.MigrateUp(migrateCtx, appConfig.MigrationsPath)
	
	if migrateErr != nil {
		logger.Fatal().Msg("Migrate database error: " + migrateErr.Error())
	}



	tokensProvider.Init(appConfig.AccessTokenKey, appConfig.RefreshTokenKey, appConfig.AccessTokenLifeTime, appConfig.RefreshTokenLifeTime, appConfig.AccessPartLen)

	logicProvider.Init(userStorage, tokensProvider, passwordHasher, uuidProvider)

	casesProvider.Init(appConfig, logger, logicProvider)

	ctx, ctxClose := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer ctxClose()

	handlersProvider.Init(ctx, casesProvider)

	api.Startup(ctx, handlersProvider, appConfig, logger)
}
