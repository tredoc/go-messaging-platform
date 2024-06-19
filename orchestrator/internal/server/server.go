package server

import (
	"context"
	"fmt"
	"github.com/tredoc/go-messaging-platform/orchestrator/internal/config"
	"github.com/tredoc/go-messaging-platform/orchestrator/pb"
	"google.golang.org/grpc"
	"log/slog"
	"net"
)

func Run(_ context.Context, cfg config.Config, grpcHandler GRPCHandler) error {
	log := slog.Default()

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.Port))
	if err != nil {
		return fmt.Errorf("failed to listen on port %d: %w", cfg.Port, err)
	}

	grpcLogger := grpc.UnaryInterceptor(GRPCLogger)
	grpcServer := grpc.NewServer(grpcLogger)
	pb.RegisterOrchestratorServiceServer(grpcServer, grpcHandler)

	log.Info("Starting orchestrator server", slog.Any("environment", cfg.Env), slog.Int("port", cfg.Port))

	return grpcServer.Serve(lis)
}
