package server

import (
	"context"
	"fmt"
	"github.com/tredoc/go-messaging-platform/message/internal/config"
	"github.com/tredoc/go-messaging-platform/message/pb"
	"google.golang.org/grpc"
	"log/slog"
	"net"
)

func Run(_ context.Context, cfg config.Config, grpcHandler GRPCHandler, log *slog.Logger) error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.Port))
	if err != nil {
		return fmt.Errorf("failed to listen on port %d: %w", cfg.Port, err)
	}

	grpcLogger := grpc.UnaryInterceptor(GRPCLogger(log))
	grpcServer := grpc.NewServer(grpcLogger)
	pb.RegisterMessageServiceServer(grpcServer, grpcHandler)

	log.Info("Starting message server", slog.Any("environment", cfg.Env), slog.Int("port", cfg.Port))

	return grpcServer.Serve(lis)
}
