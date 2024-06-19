package main

import (
	"context"
	"github.com/tredoc/go-messaging-platform/orchestrator/internal/client"
	"github.com/tredoc/go-messaging-platform/orchestrator/internal/config"
	"github.com/tredoc/go-messaging-platform/orchestrator/internal/logger"
	"github.com/tredoc/go-messaging-platform/orchestrator/internal/server"
	"log/slog"
	"os"
)

func main() {
	cfg, err := config.GetConfig()
	if err != nil {
		panic("configuration error: " + err.Error())
	}

	logger.SetupLogger(cfg.Env)
	log := slog.Default()

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	msgClient, err := client.RunMessageClient(ctx, cfg)
	if err != nil {
		log.Error(err.Error())
		os.Exit(1)
	}

	templateClient, err := client.RunTemplateClient(ctx, cfg)
	if err != nil {
		log.Error(err.Error())
		os.Exit(1)
	}

	grpcHandler := server.NewGRPCHandler(msgClient, templateClient)

	err = server.Run(ctx, cfg, grpcHandler)
	if err != nil {
		log.Error(err.Error())
		os.Exit(1)
	}
}
