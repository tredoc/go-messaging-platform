package server

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log/slog"
	"time"
)

func GRPCLogger(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
	log := slog.Default()

	start := time.Now()
	result, err := handler(ctx, req)
	duration := time.Since(start)

	statusCode := codes.Unknown
	if sts, ok := status.FromError(err); ok {
		statusCode = sts.Code()
	}

	logger := log.Info
	if err != nil {
		logger = log.Error
	}

	logger("GRPC call",
		slog.String("protocol", "grpc"),
		slog.String("method", info.FullMethod),
		slog.Int("status_code", int(statusCode)),
		slog.String("status_text", statusCode.String()),
		slog.Duration("duration", duration),
	)

	return result, err
}
