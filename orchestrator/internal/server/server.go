package server

import (
	"context"
	"errors"
	"fmt"
	"github.com/tredoc/go-messaging-platform/orchestrator/internal/config"
	"github.com/tredoc/go-messaging-platform/orchestrator/pb"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"log/slog"
	"net"
)

func Run(ctx context.Context, waitGroup *errgroup.Group, cfg config.Config, grpcHandler GRPCHandler) {
	log := slog.Default()

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.Port))
	if err != nil {
		log.Error("Orchestrator failed to listen on port", slog.Int("port", cfg.Port), slog.Any("error", err.Error()))
	}

	grpcLogger := grpc.UnaryInterceptor(GRPCLogger)
	grpcServer := grpc.NewServer(grpcLogger)
	pb.RegisterOrchestratorServiceServer(grpcServer, grpcHandler)

	waitGroup.Go(func() error {
		log.Info("Starting orchestrator server", slog.Any("environment", cfg.Env), slog.Int("port", cfg.Port))

		err = grpcServer.Serve(lis)
		if err != nil {
			if errors.Is(err, grpc.ErrServerStopped) {
				return nil
			}
			log.Error("Orchestrator gRPC server failed to serve")
			return err
		}

		return nil
	})

	waitGroup.Go(func() error {
		<-ctx.Done()
		log.Info("Graceful shutdown orchestrator gRPC server")

		grpcServer.GracefulStop()
		log.Info("Orchestrator gRPC server is stopped")

		return nil
	})
}
