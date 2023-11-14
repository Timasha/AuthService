package main

import (
	"auth/internal/api"
	"auth/internal/cases"
	"auth/internal/logic"
	"auth/internal/utils/config"
	"auth/internal/utils/logger"
	"auth/internal/utils/logger/logdrivers"
	"auth/internal/utils/password"
	"auth/internal/utils/storage"
	"auth/internal/utils/tokens/jwt"
	"auth/internal/utils/twofa"
	"auth/internal/utils/uuid"
	"context"
	"io"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	var (
		log       logger.Logger = logdrivers.NewZerologDriver([]io.Writer{os.Stdout})
		appConfig *config.JSONConfig

		unitedStorage  *storage.PostgresStorage = storage.NewPostgresStorage()
		tokensProvider *jwt.TokensProvider
		passwordHasher *password.BcryptPasswordHasher = &password.BcryptPasswordHasher{}
		uuidProvider   *uuid.GoogleUUIDProvider       = &uuid.GoogleUUIDProvider{}
		otpGenerator   *twofa.DefaultOtp

		//bodySerializer *body.JsonBodySerializer

		logicProvider *logic.LogicProvider
		casesProvider *cases.CasesProvider
		authApi       *api.Auth
	)

	appConfig, readConfigErr := config.ReadJsonConfig("./config.json", log)
	if readConfigErr != nil {
		log.Log(logger.LogMsg{
			Time:     time.Now(),
			LogLevel: logger.LogLevelFatal,
			Msg:      "Read config error: " + readConfigErr.Error(),
		})
		return
	}

	connectStorageContext, connectStorageClose := context.WithTimeout(context.Background(), time.Minute)
	defer connectStorageClose()

	connectErr := unitedStorage.Connect(connectStorageContext, appConfig.PostgresConfig)

	if connectErr != nil {
		log.Log(logger.LogMsg{
			Time:     time.Now(),
			LogLevel: logger.LogLevelFatal,
			Msg:      "Connect to db error: " + connectErr.Error(),
		})
		return
	}

	migrateCtx, migrateClose := context.WithTimeout(context.Background(), time.Minute)
	defer migrateClose()

	migrateErr := unitedStorage.MigrateUp(migrateCtx, appConfig.MigrationsPath)

	if migrateErr != nil {
		log.Log(logger.LogMsg{
			Time:     time.Now(),
			LogLevel: logger.LogLevelFatal,
			Msg:      "Migrate database error: " + migrateErr.Error(),
		})
		return
	}

	tokensProvider = jwt.New(appConfig.AccessTokenKey, appConfig.RefreshTokenKey, appConfig.AccessTokenLifeTime, appConfig.RefreshTokenLifeTime, appConfig.AccessPartLen)

	logicProvider = logic.New(unitedStorage, unitedStorage, tokensProvider, passwordHasher, uuidProvider, otpGenerator)

	casesProvider = cases.New(appConfig, log, logicProvider)

	rolesCreatingTimeoutCtx, _ := context.WithTimeout(context.Background(), time.Second*15)

	for i := 0; i < len(appConfig.Roles); i++ {
		casesProvider.AddRole(cases.AddRoleArgs{
			Ctx:  rolesCreatingTimeoutCtx,
			Role: appConfig.Roles[i],
		})
	}

	ctx, ctxClose := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer unitedStorage.Close()
	defer ctxClose()

	authApi = api.New(ctx, casesProvider, appConfig /*bodySerializer,*/, log)

	authApi.Start()
}
