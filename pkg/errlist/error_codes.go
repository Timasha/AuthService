package errlist

const (
	SuccessCode int64 = iota
	ErrServiceInternalCode
	ErrServiceNotAvaliableCode
)
const (
	ErrExpiredAccessTokenCode int64 = iota + 1001
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
	ErrUnauthorizedCode
)
const (
	ErrInvalidLoginOrPasswordCode int64 = iota + 2001
	ErrTooShortLoginOrPasswordCode
)
const (
	ErrInvalidInputCode int64 = iota + 3001
)
