package app

import (
	"context"
	"os"
	"time"

	"github.com/Timasha/AuthService/internal/api/grpc"
	"github.com/Timasha/AuthService/internal/usecase"
	"github.com/Timasha/AuthService/utils/config"
	"github.com/Timasha/AuthService/utils/grpcserver"
	"github.com/Timasha/AuthService/utils/password"
	"github.com/Timasha/AuthService/utils/storage"
	"github.com/Timasha/AuthService/utils/tokens/jwt"
	"github.com/Timasha/AuthService/utils/twofa"
	"github.com/Timasha/AuthService/utils/uuid"
	"github.com/gofiber/fiber/v2/log"
	"github.com/rs/zerolog"
)

type Lifecycle interface {
	Start(ctx context.Context) error
	Stop(ctx context.Context) error
}

type (
	App struct {
		cfg  config.Config
		cmps []cmp
	}
	cmp struct {
		Service Lifecycle
		Name    string
	}
)

func New(cfg config.Config) *App {
	return &App{cfg: cfg}
}

func (a *App) Start(ctx context.Context) error {
	log.Info("Start application")

	tokensProvider := jwt.New(a.cfg.Tokens)
	postgresStorage := storage.NewPostgresStorage(a.cfg.Postgres)
	otp := twofa.New(a.cfg.TwoFA)

	uc := usecase.New(
		a.cfg.UseCase,
		zerolog.New(os.Stdout),
		postgresStorage,
		postgresStorage,
		tokensProvider,
		&password.BcryptPasswordHasher{},
		&uuid.GoogleUUIDProvider{},
		otp,
	)

	api := grpc.NewAPI(uc)
	middleware := grpc.NewMiddleware(a.cfg.Middleware, uc)

	server := grpcserver.New(a.cfg.Server, api, middleware)

	a.cmps = []cmp{
		{postgresStorage, "postgres storage"},
		{server, "grpc server"},
	}

	for _, cmp := range a.cmps {
		log.Infof("Starting component: %s", cmp.Name)

		err := cmp.Service.Start(ctx)
		if err != nil {
			log.Fatalf("Cant start component: %s", cmp.Name)
		}

		log.Infof("Component started: %s", cmp.Name)
	}

	log.Info("Application started\n")

	return nil
}

func (a *App) Stop(ctx context.Context) error {
	for _, cmp := range a.cmps {
		log.Infof("Stopping component: %s", cmp.Name)

		err := cmp.Service.Stop(ctx)
		if err != nil {
			log.Errorf("Cant stop component: %s", cmp.Name)
		}

		log.Infof("Component stopped: %s", cmp.Name)
	}

	return nil
}

func (a *App) GetStartTimeout() time.Duration {
	return time.Minute
}

func (a *App) GetStopTimeout() time.Duration {
	return time.Minute
}
