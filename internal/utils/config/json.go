package config

import (
	"AuthService/internal/logic/models"
	"AuthService/internal/utils/logger"
	"AuthService/internal/utils/storage"
	"encoding/json"
	"io"
	"os"
	"time"
)

type JSONConfig struct {
	OrganizationName string `json:"organizationName"`

	ApiPort string `json:"apiPort"`

	MinLoginLen    int `json:"minLoginLen"`
	MinPasswordLen int `json:"minPasswordLen"`


	// Lifetime in minutes
	AccessTokenLifeTime int64 `json:"accessTokenLifeTime"`
	AccessTokenKey string `json:"accessTokenKey"`


	// Lifetime in hours
	RefreshTokenLifeTime int64 `json:"refreshTokenLifeTime"`
	RefreshTokenKey string `json:"refreshTokenKey"`
	AccessPartLen        int   `json:"accessPartLen"`

	// Lifetime in minute
	IntermediateTokenLifeTime int64 `json:"intermediateTokenLifetime"`
	IntermediateTokenKey string `json:"intermediateTokenKey"`


	PostgresConfig storage.PostgresStorageConfig `json:"postgresConfig,omitempty"`
	MigrationsPath string                        `json:"migrationsPath"`
	Roles          []models.Role                 `json:"roles"`
	DefaultRoleId  models.RoleId                 `json:"defaultRoleId"`
}

func (j *JSONConfig) GetMinLoginLen() int {
	return j.MinLoginLen
}

func (j *JSONConfig) GetMinPasswordLen() int {
	return j.MinPasswordLen
}

func (j *JSONConfig) GetDefaultUserRoleId() models.RoleId {
	return j.DefaultRoleId
}

func (j *JSONConfig) GetApiPort() string {
	return j.ApiPort
}

func ReadJsonConfig(path string, log logger.Logger) (*JSONConfig, error) {
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

	CheckConfigValues(config, log)

	return config, nil
}

func CheckConfigValues(config *JSONConfig, log logger.Logger) {
	if len(config.AccessTokenKey) < 10 {
		log.Log(logger.LogMsg{
			Time:     time.Now(),
			LogLevel: logger.LogLevelWarn,
			Msg:      "Too short access token sign key. Not safe usage.",
		})
	}
	if len(config.RefreshTokenKey) < 10 {
		log.Log(logger.LogMsg{
			Time:     time.Now(),
			LogLevel: logger.LogLevelWarn,
			Msg:      "Too short refresh token sign key. Not safe usage.",
		})
	}
	if len(config.IntermediateTokenKey) < 10{
		log.Log(logger.LogMsg{
			Time: time.Now(),
			LogLevel: logger.LogLevelWarn,
			Msg: "Too short intermediate token sign key. Not safe usage.",
		})
	}
	if (config.Roles == nil){
		log.Log(logger.LogMsg{
			Time: time.Now(),
			LogLevel: logger.LogLevelFatal,
			Msg: "Cannot use auth service without any roles. Please use default config.",
		})
	}

	if config.AccessPartLen < 5 {
		log.Log(logger.LogMsg{
			Time:     time.Now(),
			LogLevel: logger.LogLevelWarn,
			Msg:      "Too short access part length for creating refresh token. Not safe usage.",
		})
	}
}
