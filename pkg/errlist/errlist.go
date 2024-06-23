package errlist

import (
	"github.com/Timasha/customErrs/pkg/errs"
)

var (
	ErrInvalidLoginOrPassword = errs.NewHttp(
		"invalid login or password",
		ErrInvalidLoginOrPasswordCode,
		errs.StatusBadRequest,
	)

	ErrServiceNotAvaliable = errs.NewHttp(
		"service not avaliable",
		ErrServiceNotAvaliableCode,
		errs.StatusBadRequest,
	)

	ErrTooShortLoginOrPassword = errs.NewHttp(
		"too short login or password",
		ErrTooShortLoginOrPasswordCode,
		errs.StatusBadRequest,
	)

	ErrExpiredAccessToken = errs.NewHttp(
		"access token is expired",
		ErrExpiredAccessTokenCode,
		errs.StatusBadRequest,
	)

	ErrExpiredIntermediateToken = errs.NewHttp(
		"intermediate token is expired",
		ErrExpiredIntermediateTokenCode,
		errs.StatusBadRequest,
	)

	ErrExpiredRefreshToken = errs.NewHttp(
		"refresh token is expired",
		ErrExpiredRefreshTokenCode,
		errs.StatusBadRequest,
	)

	ErrInvalidAccessToken = errs.NewHttp(
		"invalid access token",
		ErrInvalidAccessTokenCode,
		errs.StatusBadRequest,
	)

	ErrInvalidIntermediateToken = errs.NewHttp(
		"invalid intermediate token",
		ErrInvalidIntermediateTokenCode,
		errs.StatusBadRequest,
	)

	ErrInvalidRefreshToken = errs.NewHttp(
		"invalid refresh token",
		ErrInvalidRefreshTokenCode,
		errs.StatusBadRequest,
	)

	ErrRoleAlreadyExists = errs.NewHttp(
		"role already exists",
		ErrRoleAlreadyExistsCode,
		errs.StatusBadRequest,
	)

	ErrRoleNotExists = errs.NewHttp(
		"role not exists",
		ErrRoleNotExistsCode,
		errs.StatusBadRequest,
	)

	ErrUserAlreadyExists = errs.NewHttp(
		"user already exists",
		ErrUserAlreadyExistsCode,
		errs.StatusBadRequest,
	)

	ErrUserNotExists = errs.NewHttp(
		"user not exists",
		ErrUserNotExistsCode,
		errs.StatusBadRequest,
	)

	ErrRoleHasNoAccess = errs.NewHttp(
		"role has no access",
		ErrRoleHasNoAccessCode,
		401,
	)
	ErrInvalidOtp = errs.NewHttp(
		"invalid otp code",
		ErrInvalidOtpCode,
		errs.StatusBadRequest,
	)

	ErrOtpAlreadyDisabled = errs.NewHttp(
		"otp is already disabled",
		ErrOtpAlreadyDisabledCode,
		errs.StatusBadRequest,
	)

	ErrOtpAlreadyEnabled = errs.NewHttp(
		"otp is already enabled",
		ErrOtpAlreadyEnabledCode,
		errs.StatusBadRequest,
	)

	ErrCantScanRoleID = errs.NewHttp(
		"invalid role id",
		ErrCantScanRoleIDCode,
		errs.StatusInternal,
	)

	ErrWrongAuthorizationMethod = errs.NewHttp(
		"wrong authorization method",
		ErrWrongAuthorizationMethodCode,
		errs.StatusBadRequest,
	)

	ErrUnauthorized = errs.New(
		"user unauthorized",
		ErrUnauthorizedCode,
	)
)
