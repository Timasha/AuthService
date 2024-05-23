package grpc

import (
	"context"
	"slices"
	"strings"

	"github.com/Timasha/AuthService/internal/usecase"
	"github.com/Timasha/AuthService/pkg/errlist"
	"github.com/Timasha/AuthService/utils/consts"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type MiddlewareConfig struct {
	AuthForMethods []string
}

type Middleware struct {
	cfg MiddlewareConfig
	uc  *usecase.Provider
}

func NewMiddleware(cfg MiddlewareConfig, uc *usecase.Provider) *Middleware {
	cfg.AuthForMethods = append(cfg.AuthForMethods, consts.AuthForMethods...)

	return &Middleware{
		cfg: cfg,
		uc:  uc,
	}
}

func (m *Middleware) Auth(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (resp interface{}, err error) {
	if !slices.Contains(m.cfg.AuthForMethods, info.FullMethod) {
		return handler(ctx, req)
	}

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		//TODO: modify errs package to return grpc errs
		return resp, status.Errorf(codes.InvalidArgument, "Retrieving metadata is failed")
	}

	authHeader, ok := md["Authorization"]
	if !ok {
		//TODO: modify errs package to return grpc errs
		return resp, status.Errorf(codes.Unauthenticated, "Authorization token is not supplied")
	}

	authorizeHeaderParts := strings.Split(authHeader[0], " ")
	if len(authorizeHeaderParts) != 2 || authorizeHeaderParts[0] != "Bearer" {
		return resp, errlist.ErrWrongAuthorizationMethod
	}

	ret, err := m.uc.AuthorizeUser(ctx, usecase.AuthorizeUserRequest{
		AccessToken:        authorizeHeaderParts[1],
		RequiredRoleAccess: nil,
	})
	if err != nil {
		return resp, err
	}

	ctx = context.WithValue(ctx, consts.UserIDCtxKey, ret.UserID)

	return handler(ctx, req)
}
