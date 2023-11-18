package errsutil

type AuthErrCode int

const (
	SuccessCode AuthErrCode = iota
	ErrServiceInternalCode 
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
	ErrOtpAlreadyEnabledCode
	ErrOtpAlreadyDisabledCode
)
const (
	ErrInvalidLoginOrPasswordCode AuthErrCode = iota + 201
	ErrTooShortLoginOrPasswordCode
)
const (
	ErrInvalidInputCode AuthErrCode = iota + 301
	ErrWrongAuthorizeMethodCode
)
