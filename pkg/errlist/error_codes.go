package errlist

const (
	SuccessCode int64 = iota
	ErrServiceInternalCode
	ErrServiceNotAvaliableCode
)
const (
	ErrExpiredAccessTokenCode int64 = iota + 101
	ErrExpiredRefreshTokenCode
	ErrExpiredIntermediateTokenCode
	ErrInvalidPasswordCode
	ErrInvalidAccessTokenCode
	ErrInvalidRefreshTokenCode
	ErrInvalidIntermediateTokenCode
	ErrUserAlreadyExistsCode
	ErrUserNotExistsCode
	ErrInvalidOtpCode
	ErrRoleHasNoAccessCode
	ErrRoleAlreadyExistsCode
	ErrRoleNotExistsCode
	ErrOtpAlreadyEnabledCode
	ErrOtpAlreadyDisabledCode
	ErrCantScanRoleIDCode
	ErrWrongAuthorizationMethodCode
)
const (
	ErrInvalidLoginOrPasswordCode int64 = iota + 201
	ErrTooShortLoginOrPasswordCode
)
const (
	ErrInvalidInputCode int64 = iota + 301
)
