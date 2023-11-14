package errsutil

type AuthErrCode int

const (
	ErrServiceInternalCode AuthErrCode = iota
	SuccessCode
	ErrServiceNotAvaliableCode
)
const (
	ErrExpiredAccessTokenCode AuthErrCode = iota + 101
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
)
const (
	ErrInvalidLoginOrPasswordCode AuthErrCode = iota + 201
	ErrTooShortLoginOrPasswordCode
)
const (
	ErrInputCode AuthErrCode = iota + 301
	ErrWrongAuthMethodCode
)
