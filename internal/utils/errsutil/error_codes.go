package errsutil

const (
	ErrServiceInternalCode     int = 0
	SuccessCode                int = 1
	ErrServiceNotAvaliableCode int = 2

	ErrExpiredAccessTokenCode  int = 101
	ErrExpiredRefreshTokenCode int = 102
	ErrInvalidPasswordCode     int = 103
	ErrInvalidAccessTokenCode  int = 104
	ErrInvalidRefreshTokenCode int = 105
	ErrUserAlreadyExistsCode   int = 106
	ErrUserNotExistsCode       int = 107

	ErrInvalidLoginOrPasswordCode  int = 201
	ErrTooShortLoginOrPasswordCode int = 202

	ErrInputCode int = 301
)
