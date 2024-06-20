package server

import (
	"context"
	"errors"
	"fmt"
	"github.com/tredoc/go-messaging-platform/message/internal/config"
	"github.com/tredoc/go-messaging-platform/message/pb"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"log/slog"
	"net"
)

func Run(ctx context.Context, waitGroup *errgroup.Group, cfg config.Config, grpcHandler GRPCHandler, log *slog.Logger) {
	grpcLogger := grpc.UnaryInterceptor(GRPCLogger(log))
	grpcServer := grpc.NewServer(grpcLogger)
	pb.RegisterMessageServiceServer(grpcServer, grpcHandler)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.Port))
	if err != nil {
		log.Error("failed to listen on port", slog.Int("port", cfg.Port), slog.Any("error", err.Error()))
	}

	waitGroup.Go(func() error {
		log.Info("Starting message server", slog.Any("environment", cfg.Env), slog.Int("port", cfg.Port))

		err = grpcServer.Serve(lis)
		if err != nil {
			if errors.Is(err, grpc.ErrServerStopped) {
				return nil
			}
			log.Error("Message gRPC server failed to serve")
			return err
		}

		return nil
	})

	waitGroup.Go(func() error {
		<-ctx.Done()
		log.Info("Graceful shutdown message gRPC server")

		grpcServer.GracefulStop()
		log.Info("Message gRPC server is stopped")

		return nil
	})
}
