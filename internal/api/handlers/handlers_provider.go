package handlers

import (
	"auth/internal/cases"
	"context"
)

type HandlersProvider struct {
	ctx           context.Context
	casesProvider *cases.CasesProvider
}

func (h *HandlersProvider) Init(ctx context.Context, casesProvider *cases.CasesProvider) {
	h.ctx = ctx
	h.casesProvider = casesProvider
}
