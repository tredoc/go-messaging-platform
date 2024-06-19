package server

import (
	"context"
	"fmt"
	"github.com/tredoc/go-messaging-platform/template/internal/config"
	"github.com/tredoc/go-messaging-platform/template/pb"
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
	pb.RegisterTemplateServiceServer(grpcServer, grpcHandler)

	log.Info("Starting template server", slog.Any("environment", cfg.Env), slog.Int("port", cfg.Port))

	return grpcServer.Serve(lis)
}
