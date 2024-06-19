package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"

	"github.com/Timasha/AuthService/internal/storage"

	api "github.com/Timasha/AuthService/internal/api/grpc"
	"github.com/Timasha/AuthService/internal/usecase"
	"github.com/Timasha/AuthService/utils/grpcserver"
	"github.com/Timasha/AuthService/utils/tokens/jwt"
	"github.com/Timasha/AuthService/utils/twofa"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2/log"
)

type Config struct {
	Server     grpcserver.Config    `validate:"required"`
	Middleware api.MiddlewareConfig `validate:"required"`

	UseCase usecase.Config `validate:"required"`

	Tokens jwt.Config `validate:"required"`

	TwoFA twofa.Config `validate:"required"`

	Postgres storage.PostgresStorageConfig `validate:"required"`
}

func ReadConfigJSON(path string) (cfg *Config, err error) {
	cfg = new(Config)

	file, openErr := os.Open(path)
	if openErr != nil {
		return nil, openErr
	}

	fileData, readErr := io.ReadAll(file)

	if readErr != nil {
		return nil, readErr
	}

	unmarshalErr := json.Unmarshal(fileData, cfg)

	if unmarshalErr != nil {
		return nil, unmarshalErr
	}

	valid := validator.New()

	err = valid.Struct(cfg)
	if err != nil {
		var validErrs validator.ValidationErrors

		ok := errors.As(err, &validErrs)
		if !ok {
			return cfg, nil
		}

		return cfg, fmt.Errorf("%w", validErrs)
	}

	log.Infof("Config loaded: %v", *cfg)

	return cfg, nil
}
