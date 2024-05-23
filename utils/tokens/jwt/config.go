package jwt

import "github.com/Timasha/AuthService/utils/duration"

type Config struct {
	// Lifetime in minutes
	AccessTokenKey      string
	AccessTokenLifeTime duration.Minutes

	// Lifetime in hours
	RefreshTokenKey      string
	RefreshTokenLifeTime duration.Hours
	AccessPartLen        int

	// Lifetime in minutes
	IntermediateTokenKey      string
	IntermediateTokenLifeTime duration.Minutes
}
