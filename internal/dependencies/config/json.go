package config

import (
	"auth/internal/dependencies/storage"
	"encoding/json"
	"io"
	"os"
	"time"

	"github.com/rs/zerolog"
)

type JSONConfig struct {
	ApiPort string `json:"apiPort"`

	MinLoginLen    int `json:"minLoginLen"`
	MinPasswordLen int `json:"minPasswordLen"`

	AccessTokenKey string `json:"accessTokenKey"`
	// Lifetime in minutes
	AccessTokenLifeTime int64 `json:"accessTokenLifeTime"`

	RefreshTokenKey string `json:"refreshTokenKey"`
	// Lifetime in hours
	RefreshTokenLifeTime int64 `json:"refreshTokenLifeTime"`
	AccessPartLen        int   `json:"accessPartLen"`

	PostgresConfig storage.PostgresUserStorageConfig `json:"postgresConfig,omitempty"`
}

func (j *JSONConfig) GetMinLoginLen() int {
	return j.MinLoginLen
}

func (j *JSONConfig) GetMinPasswordLen() int {
	return j.MinPasswordLen
}
func (j *JSONConfig) GetApiPort() string {
	return j.ApiPort
}

func ReadJsonConfig(path string, logger zerolog.Logger) (*JSONConfig, error) {
	file, openErr := os.Open(path)

	if openErr != nil {
		return nil, openErr
	}

	fileData, readErr := io.ReadAll(file)

	if readErr != nil {
		return nil, readErr
	}
	var config *JSONConfig = &JSONConfig{}

	unmarshalErr := json.Unmarshal(fileData, config)

	if unmarshalErr != nil {
		return nil, unmarshalErr
	}

	CheckConfigValues(config, logger)

	return config, nil
}

func CheckConfigValues(config *JSONConfig, logger zerolog.Logger) {
	if len(config.AccessTokenKey) < 5 {
		logger.Warn().Msg("Too short access token sign key. Not safe usage.")
	}
	if len(config.RefreshTokenKey) < 5 {
		logger.Warn().Msg("Too short refresh token sign key. Not safe usage.")
	}

	if time.Duration(config.AccessTokenLifeTime)*time.Minute < time.Hour {
		logger.Warn().Msg("Too short access token lifetime. Using default values.")
		config.AccessTokenLifeTime = 60
	}
	if time.Duration(config.RefreshTokenLifeTime)*time.Hour < time.Hour*24 {
		logger.Warn().Msg("Too short refresh token lifetime. Using default values.")
		config.RefreshTokenLifeTime = 24
	}

	if config.AccessPartLen < 5 {
		logger.Warn().Msg("Too short access part length for creating refresh token. Using deault values.")
		config.AccessPartLen = 5
	}
}
