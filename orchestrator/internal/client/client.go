package client

import (
	"context"
	"github.com/tredoc/go-messaging-platform/orchestrator/internal/config"
	"github.com/tredoc/go-messaging-platform/orchestrator/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func RunMessageClient(_ context.Context, cfg config.Config) (pb.MessageServiceClient, error) {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	conn, err := grpc.NewClient(cfg.MessageAddr, opts...)
	if err != nil {
		return nil, err
	}

	client := pb.NewMessageServiceClient(conn)

	return client, nil
}

func RunTemplateClient(_ context.Context, cfg config.Config) (pb.TemplateServiceClient, error) {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	conn, err := grpc.NewClient(cfg.TemplateAddr, opts...)
	if err != nil {
		return nil, err
	}

	client := pb.NewTemplateServiceClient(conn)

	return client, nil
}
