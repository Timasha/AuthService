package handlers

import (
	"auth/internal/cases"
	"context"
)

type FiberHandlersProvider struct {
	ctx           context.Context
	casesProvider *cases.CasesProvider
}

func (h *FiberHandlersProvider) Init(ctx context.Context, casesProvider *cases.CasesProvider) {
	h.ctx = ctx
	h.casesProvider = casesProvider
}
