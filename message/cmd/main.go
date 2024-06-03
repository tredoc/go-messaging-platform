package main

import (
	"context"
	"github.com/tredoc/go-messaging-platform/message/internal/config"
	"github.com/tredoc/go-messaging-platform/message/internal/handler"
	"github.com/tredoc/go-messaging-platform/message/internal/server"
	"log"
)

func main() {
	cfg, err := config.GetConfig()
	if err != nil {
		log.Fatalf("configuration error: %v", err)
	}

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	grpcHandler := handler.NewGRPCMessageHandler()
	GRPCServer := server.NewGRPCServer(grpcHandler)

	err = server.Run(ctx, GRPCServer, cfg)
	if err != nil {
		log.Fatal(err)
	}
}
