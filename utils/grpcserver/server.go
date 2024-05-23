package grpcserver

import (
	"context"

	api "github.com/Timasha/AuthService/internal/api/grpc"
	grpcapi "github.com/Timasha/AuthService/pkg/api"
	"github.com/Timasha/AuthService/utils/consts"
	"google.golang.org/grpc"

	"net"
	"time"

	"google.golang.org/grpc/keepalive"

	"google.golang.org/grpc/reflection"
)

type Server struct {
	cfg        Config
	grpcServer *grpc.Server
	api        *api.API
	middleware *api.Middleware
}

func New(
	cfg Config,
	api *api.API,
	middleware *api.Middleware,
) *Server {
	return &Server{
		cfg:        cfg,
		api:        api,
		middleware: middleware,
	}
}

func (s *Server) Start(_ context.Context) (err error) {
	unaryInterceptor := grpc.ChainUnaryInterceptor(
		s.middleware.Auth,
	)
	streamInterceptor := grpc.ChainStreamInterceptor()

	s.grpcServer = grpc.NewServer(
		grpc.KeepaliveParams(keepalive.ServerParameters{
			MaxConnectionIdle: s.cfg.MaxConnectionIdle.Duration,
			MaxConnectionAge:  s.cfg.MaxConnectionAge.Duration,
			Timeout:           s.cfg.Timeout.Duration,
			Time:              s.cfg.Time.Duration,
		}),
		// grpc.StatsHandler(otelgrpc.NewServerHandler()),
		unaryInterceptor,
		streamInterceptor,
	)

	grpcapi.RegisterAuthServer(s.grpcServer, s.api)

	reflection.Register(s.grpcServer)

	listener, err := net.Listen(consts.TCP, s.cfg.Host)
	if err != nil {
		return err
	}
	errCh := make(chan error)

	go func() {
		if err := s.grpcServer.Serve(listener); err != nil {
			errCh <- err
		}
	}()
	select {
	case err := <-errCh:
		return err
	case <-time.After(s.cfg.StartTimeout.Duration):
		return nil
	}
}

func (s *Server) Stop(_ context.Context) error {
	stopCh := make(chan any)
	go func() {
		s.grpcServer.GracefulStop()
		stopCh <- nil
	}()
	select {
	case <-time.After(s.cfg.StopTimeout.Duration):
		return nil
	case <-stopCh:
		return nil
	}
}
