package grpc

import (
	"context"
	"github.com/Timasha/AuthService/pkg/errlist"
	"github.com/Timasha/AuthService/utils/consts"

	"github.com/Timasha/AuthService/internal/usecase"
	"github.com/Timasha/AuthService/pkg/api"
	"github.com/Timasha/AuthService/utils/convert"
	"google.golang.org/protobuf/types/known/emptypb"
)

type API struct {
	api.UnimplementedAuthServer
	uc *usecase.Provider
}

func NewAPI(uc *usecase.Provider) *API {
	return &API{
		uc: uc,
	}
}

func (a *API) AuthenticateUserByLogin(
	ctx context.Context,
	req *api.AuthenticateUserByLoginRequest,
) (resp *api.AuthenticateUserByLoginResponse, err error) {
	resp = new(api.AuthenticateUserByLoginResponse)

	ucResp, err := a.uc.AuthenticateUserByLogin(ctx, convert.AuthenticateUserByLoginRequestFromProto(req))
	if err != nil {
		return resp, err
	}

	resp = convert.AuthenticateUserByLoginResponseToProto(ucResp)

	return resp, nil
}

func (a *API) ContinueAuthenticateOtpUser(
	ctx context.Context,
	req *api.ContinueAuthenticateOtpUserByLoginRequest,
) (resp *api.ContinueAuthenticateOtpUserByLoginResponse, err error) {
	resp = new(api.ContinueAuthenticateOtpUserByLoginResponse)

	ucResp, err := a.uc.ContinueAuthenticateOtpUserByLogin(
		ctx,
		convert.ContinueAuthenticateOtpUserByLoginRequestFromProto(req),
	)
	if err != nil {
		return resp, err
	}

	resp = convert.ContinueAuthenticateOtpUserByLoginResponseToProto(ucResp)

	return resp, nil
}

func (a *API) Register(ctx context.Context, req *api.RegisterUserRequest) (resp *emptypb.Empty, err error) {
	resp = new(emptypb.Empty)

	err = a.uc.RegisterUser(ctx, convert.RegisterUserRequestFromProto(req))

	return resp, err
}

func (a *API) Authorize(ctx context.Context, req *api.AuthorizeRequest) (resp *api.AuthorizeResponse, err error) {
	resp = new(api.AuthorizeResponse)

	ucResp, err := a.uc.AuthorizeUser(ctx, convert.AuthorizeUserRequestFromProto(req))
	if err != nil {
		return resp, err
	}

	resp = convert.AuthorizeUserResponseToProto(ucResp)

	return resp, nil
}

func (a *API) RefreshTokens(
	ctx context.Context,
	req *api.RefreshTokensRequest,
) (resp *api.RefreshTokensResponse, err error) {
	resp = new(api.RefreshTokensResponse)

	ucResp, err := a.uc.RefreshTokens(ctx, convert.RefreshTokensRequestFromProto(req))
	if err != nil {
		return resp, err
	}

	resp = convert.RefreshTokensResponseToProto(ucResp)

	return resp, nil
}

func (a *API) EnableOtpAuthentication(
	ctx context.Context,
	_ *emptypb.Empty,
) (resp *api.EnableOtpAuthenticationResponse, err error) {
	resp = new(api.EnableOtpAuthenticationResponse)

	userID, ok := ctx.Value(consts.UserIDCtxKey).(*int64)
	if !ok || userID == nil {
		return resp, errlist.ErrUnauthorized
	}

	ucResp, err := a.uc.EnableOtpAuthentication(ctx, usecase.EnableOtpAuthenticationRequest{
		UserID: *userID,
	})
	if err != nil {
		return resp, err
	}

	resp = convert.EnableOtpAuthenticationResponseToProto(ucResp)

	return resp, nil
}

func (a *API) DisableOtpAuthentication(ctx context.Context, _ *emptypb.Empty) (resp *emptypb.Empty, err error) {
	resp = new(emptypb.Empty)

	userID, ok := ctx.Value(consts.UserIDCtxKey).(*int64)
	if !ok || userID == nil {
		return resp, errlist.ErrUnauthorized
	}

	err = a.uc.DisableOtpAuthentication(ctx, usecase.DisableOtpAuthenticationRequest{
		UserID: *userID,
	})
	if err != nil {
		return resp, err
	}

	return resp, nil
}
