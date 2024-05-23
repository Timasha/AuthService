package grpcserver

import "github.com/Timasha/AuthService/utils/duration"

type Config struct {
	Host              string `validate:"required"`
	SecureConnection  bool
	StartTimeout      duration.Seconds `validate:"required"`
	StopTimeout       duration.Seconds `validate:"required"`
	MaxConnectionIdle duration.Seconds `validate:"required"`
	Timeout           duration.Seconds `validate:"required"`
	MaxConnectionAge  duration.Seconds `validate:"required"`
	Time              duration.Seconds `validate:"required"`
}
