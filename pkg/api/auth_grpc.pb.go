// protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative auth.proto

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.2
// source: pkg/api/auth.proto

package api

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	Auth_AuthenticateUserByLogin_FullMethodName     = "/auth.Auth/AuthenticateUserByLogin"
	Auth_ContinueAuthenticateOtpUser_FullMethodName = "/auth.Auth/ContinueAuthenticateOtpUser"
	Auth_Register_FullMethodName                    = "/auth.Auth/Register"
	Auth_Authorize_FullMethodName                   = "/auth.Auth/Authorize"
	Auth_RefreshTokens_FullMethodName               = "/auth.Auth/RefreshTokens"
	Auth_EnableOtpAuthentication_FullMethodName     = "/auth.Auth/EnableOtpAuthentication"
	Auth_DisableOtpAuthentication_FullMethodName    = "/auth.Auth/DisableOtpAuthentication"
)

// AuthClient is the client API for Auth service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AuthClient interface {
	AuthenticateUserByLogin(ctx context.Context, in *AuthenticateUserByLoginRequest, opts ...grpc.CallOption) (*AuthenticateUserByLoginResponse, error)
	ContinueAuthenticateOtpUser(ctx context.Context, in *ContinueAuthenticateOtpUserByLoginRequest, opts ...grpc.CallOption) (*ContinueAuthenticateOtpUserByLoginResponse, error)
	Register(ctx context.Context, in *RegisterUserRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	Authorize(ctx context.Context, in *AuthorizeRequest, opts ...grpc.CallOption) (*AuthorizeResponse, error)
	RefreshTokens(ctx context.Context, in *RefreshTokensRequest, opts ...grpc.CallOption) (*RefreshTokensResponse, error)
	EnableOtpAuthentication(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*EnableOtpAuthenticationResponse, error)
	DisableOtpAuthentication(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type authClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthClient(cc grpc.ClientConnInterface) AuthClient {
	return &authClient{cc}
}

func (c *authClient) AuthenticateUserByLogin(ctx context.Context, in *AuthenticateUserByLoginRequest, opts ...grpc.CallOption) (*AuthenticateUserByLoginResponse, error) {
	out := new(AuthenticateUserByLoginResponse)
	err := c.cc.Invoke(ctx, Auth_AuthenticateUserByLogin_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) ContinueAuthenticateOtpUser(ctx context.Context, in *ContinueAuthenticateOtpUserByLoginRequest, opts ...grpc.CallOption) (*ContinueAuthenticateOtpUserByLoginResponse, error) {
	out := new(ContinueAuthenticateOtpUserByLoginResponse)
	err := c.cc.Invoke(ctx, Auth_ContinueAuthenticateOtpUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) Register(ctx context.Context, in *RegisterUserRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, Auth_Register_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) Authorize(ctx context.Context, in *AuthorizeRequest, opts ...grpc.CallOption) (*AuthorizeResponse, error) {
	out := new(AuthorizeResponse)
	err := c.cc.Invoke(ctx, Auth_Authorize_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) RefreshTokens(ctx context.Context, in *RefreshTokensRequest, opts ...grpc.CallOption) (*RefreshTokensResponse, error) {
	out := new(RefreshTokensResponse)
	err := c.cc.Invoke(ctx, Auth_RefreshTokens_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) EnableOtpAuthentication(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*EnableOtpAuthenticationResponse, error) {
	out := new(EnableOtpAuthenticationResponse)
	err := c.cc.Invoke(ctx, Auth_EnableOtpAuthentication_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) DisableOtpAuthentication(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, Auth_DisableOtpAuthentication_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthServer is the server API for Auth service.
// All implementations must embed UnimplementedAuthServer
// for forward compatibility
type AuthServer interface {
	AuthenticateUserByLogin(context.Context, *AuthenticateUserByLoginRequest) (*AuthenticateUserByLoginResponse, error)
	ContinueAuthenticateOtpUser(context.Context, *ContinueAuthenticateOtpUserByLoginRequest) (*ContinueAuthenticateOtpUserByLoginResponse, error)
	Register(context.Context, *RegisterUserRequest) (*emptypb.Empty, error)
	Authorize(context.Context, *AuthorizeRequest) (*AuthorizeResponse, error)
	RefreshTokens(context.Context, *RefreshTokensRequest) (*RefreshTokensResponse, error)
	EnableOtpAuthentication(context.Context, *emptypb.Empty) (*EnableOtpAuthenticationResponse, error)
	DisableOtpAuthentication(context.Context, *emptypb.Empty) (*emptypb.Empty, error)
	mustEmbedUnimplementedAuthServer()
}

// UnimplementedAuthServer must be embedded to have forward compatible implementations.
type UnimplementedAuthServer struct {
}

func (UnimplementedAuthServer) AuthenticateUserByLogin(context.Context, *AuthenticateUserByLoginRequest) (*AuthenticateUserByLoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AuthenticateUserByLogin not implemented")
}
func (UnimplementedAuthServer) ContinueAuthenticateOtpUser(context.Context, *ContinueAuthenticateOtpUserByLoginRequest) (*ContinueAuthenticateOtpUserByLoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ContinueAuthenticateOtpUser not implemented")
}
func (UnimplementedAuthServer) Register(context.Context, *RegisterUserRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (UnimplementedAuthServer) Authorize(context.Context, *AuthorizeRequest) (*AuthorizeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Authorize not implemented")
}
func (UnimplementedAuthServer) RefreshTokens(context.Context, *RefreshTokensRequest) (*RefreshTokensResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RefreshTokens not implemented")
}
func (UnimplementedAuthServer) EnableOtpAuthentication(context.Context, *emptypb.Empty) (*EnableOtpAuthenticationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EnableOtpAuthentication not implemented")
}
func (UnimplementedAuthServer) DisableOtpAuthentication(context.Context, *emptypb.Empty) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DisableOtpAuthentication not implemented")
}
func (UnimplementedAuthServer) mustEmbedUnimplementedAuthServer() {}

// UnsafeAuthServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AuthServer will
// result in compilation errors.
type UnsafeAuthServer interface {
	mustEmbedUnimplementedAuthServer()
}

func RegisterAuthServer(s grpc.ServiceRegistrar, srv AuthServer) {
	s.RegisterService(&Auth_ServiceDesc, srv)
}

func _Auth_AuthenticateUserByLogin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthenticateUserByLoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).AuthenticateUserByLogin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Auth_AuthenticateUserByLogin_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).AuthenticateUserByLogin(ctx, req.(*AuthenticateUserByLoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_ContinueAuthenticateOtpUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ContinueAuthenticateOtpUserByLoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).ContinueAuthenticateOtpUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Auth_ContinueAuthenticateOtpUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).ContinueAuthenticateOtpUser(ctx, req.(*ContinueAuthenticateOtpUserByLoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Auth_Register_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).Register(ctx, req.(*RegisterUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_Authorize_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthorizeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).Authorize(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Auth_Authorize_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).Authorize(ctx, req.(*AuthorizeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_RefreshTokens_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RefreshTokensRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).RefreshTokens(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Auth_RefreshTokens_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).RefreshTokens(ctx, req.(*RefreshTokensRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_EnableOtpAuthentication_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).EnableOtpAuthentication(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Auth_EnableOtpAuthentication_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).EnableOtpAuthentication(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_DisableOtpAuthentication_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).DisableOtpAuthentication(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Auth_DisableOtpAuthentication_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).DisableOtpAuthentication(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// Auth_ServiceDesc is the grpc.ServiceDesc for Auth service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Auth_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "auth.Auth",
	HandlerType: (*AuthServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AuthenticateUserByLogin",
			Handler:    _Auth_AuthenticateUserByLogin_Handler,
		},
		{
			MethodName: "ContinueAuthenticateOtpUser",
			Handler:    _Auth_ContinueAuthenticateOtpUser_Handler,
		},
		{
			MethodName: "Register",
			Handler:    _Auth_Register_Handler,
		},
		{
			MethodName: "Authorize",
			Handler:    _Auth_Authorize_Handler,
		},
		{
			MethodName: "RefreshTokens",
			Handler:    _Auth_RefreshTokens_Handler,
		},
		{
			MethodName: "EnableOtpAuthentication",
			Handler:    _Auth_EnableOtpAuthentication_Handler,
		},
		{
			MethodName: "DisableOtpAuthentication",
			Handler:    _Auth_DisableOtpAuthentication_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/api/auth.proto",
}
