package main

import (
	"context"
	"github.com/tredoc/go-messaging-platform/orchestrator/internal/client"
	"github.com/tredoc/go-messaging-platform/orchestrator/internal/config"
	"github.com/tredoc/go-messaging-platform/orchestrator/internal/server"
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

	msgClient, err := client.RunMessageClient(ctx, cfg)
	if err != nil {
		log.Fatal(err)
	}

	templateClient, err := client.RunTemplateClient(ctx, cfg)
	if err != nil {
		log.Fatal(err)
	}

	grpcHandler := server.NewGRPCHandler(msgClient, templateClient)

	err = server.Run(ctx, cfg, grpcHandler)
	if err != nil {
		log.Fatal(err)
	}
}
