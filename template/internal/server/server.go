package server

import (
	"context"
	"fmt"
	"github.com/tredoc/go-messaging-platform/template/internal/config"
	"github.com/tredoc/go-messaging-platform/template/pb"
	"google.golang.org/grpc"
	"log"
	"net"
)

func Run(_ context.Context, cfg config.Config, grpcHandler GRPCHandler) error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", cfg.Port))
	if err != nil {
		return fmt.Errorf("failed to listen on port %s: %w", cfg.Port, err)
	}

	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	pb.RegisterTemplateServiceServer(grpcServer, grpcHandler)

	log.Printf("Starting template server on port: %s\n", cfg.Port)

	return grpcServer.Serve(lis)
}
