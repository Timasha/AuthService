package usecase

import "github.com/rs/zerolog"

type Provider struct {
	cfg    Config
	logger zerolog.Logger

	userStorage  UserStorage
	rolesStorage RolesStorage

	tokensProvider TokensProvider
	passwordHasher PasswordHasher
	uuidProvider   UUIDProvider
	otpGenerator   OtpGenerator
}

func New(cfg Config,
	logger zerolog.Logger,
	userStorage UserStorage,
	rolesStorage RolesStorage,
	tokensProvider TokensProvider,
	passwordHasher PasswordHasher,
	uuidProvider UUIDProvider,
	otpGenerator OtpGenerator,
) (c *Provider) {
	c = &Provider{
		cfg:    cfg,
		logger: logger,

		userStorage:  userStorage,
		rolesStorage: rolesStorage,

		tokensProvider: tokensProvider,
		passwordHasher: passwordHasher,
		uuidProvider:   uuidProvider,
		otpGenerator:   otpGenerator,
	}

	return
}

type Config struct {
	LoginRegexp string `json:"loginRegexp"`

	PasswordRegexp string `json:"passwordRegexp"`
}
