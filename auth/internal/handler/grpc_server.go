package handler

import (
	"errors"
	"fmt"
	"net"

	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/recovery"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"

	"auth/internal/config"
	"auth/internal/log"
	"auth/pkg/rpc"
)

var ErrRecovery = errors.New("errRecovery")

func RegisterGRPCServer(srv rpc.AuthServiceServer) (*grpc.Server, net.Listener, error) {
	grpcServer := grpc.NewServer(
		grpc.ChainUnaryInterceptor(recovery.UnaryServerInterceptor(recoveryOpts()...)),
		grpc.StatsHandler(otelgrpc.NewServerHandler()),
	)

	rpc.RegisterAuthServiceServer(grpcServer, srv)
	reflection.Register(grpcServer)

	grpcListener, err := net.Listen("tcp", fmt.Sprintf(":%v", config.Default().GRPC.Port))
	if err != nil {
		return nil, nil, err
	}

	return grpcServer, grpcListener, nil
}

func recoveryOpts() []recovery.Option {
	return []recovery.Option{
		recovery.WithRecoveryHandler(func(p any) error {
			log.Default().Error().Any("panic", p).Msgf("recovered from panic")

			return status.Error(codes.Internal, ErrRecovery.Error())
		}),
	}
}
