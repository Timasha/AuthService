package cases

import (
	"auth/internal/logic"
	"auth/internal/logic/models"
	"auth/internal/utils/errsutil"
	"auth/internal/utils/logger"
)

type CasesProvider struct {
	config UseCasesConfig
	logger logger.Logger

	logic *logic.LogicProvider
}

func New(config UseCasesConfig, logger logger.Logger, logic *logic.LogicProvider) (c *CasesProvider) {
	c = &CasesProvider{
		config: config,
		logger: logger,
		logic:  logic,
	}
	return
}

var (
	ErrServiceInternal errsutil.AuthErr = errsutil.AuthErr{
		Msg:     "internal service error",
		ErrCode: errsutil.ErrServiceInternalCode,
	}
	ErrInvalidLoginOrPassword errsutil.AuthErr = errsutil.AuthErr{
		Msg:     "invalid login or password",
		ErrCode: errsutil.ErrInvalidLoginOrPasswordCode,
	}
	ErrServiceNotAvaliable errsutil.AuthErr = errsutil.AuthErr{
		Msg:     "service not avaliable",
		ErrCode: errsutil.ErrServiceNotAvaliableCode,
	}
	ErrTooShortLoginOrPassword errsutil.AuthErr = errsutil.AuthErr{
		Msg:     "too short login or password",
		ErrCode: errsutil.ErrTooShortLoginOrPasswordCode,
	}
)

type UseCasesConfig interface {
	GetDefaultUserRoleId() models.RoleId
	GetMinLoginLen() int
	GetMinPasswordLen() int
}
