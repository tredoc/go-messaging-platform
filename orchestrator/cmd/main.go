package main

import (
	"context"
	"github.com/tredoc/go-messaging-platform/orchestrator/internal/client"
	"github.com/tredoc/go-messaging-platform/orchestrator/internal/config"
	"github.com/tredoc/go-messaging-platform/orchestrator/internal/logger"
	"github.com/tredoc/go-messaging-platform/orchestrator/internal/server"
	"golang.org/x/sync/errgroup"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
)

var interruptSignals = []os.Signal{
	os.Interrupt,
	syscall.SIGTERM,
	syscall.SIGINT,
}

func main() {
	cfg, err := config.GetConfig()
	if err != nil {
		panic("configuration error: " + err.Error())
	}

	logger.SetupLogger(cfg.Env)
	log := slog.Default()

	ctx, stop := signal.NotifyContext(context.Background(), interruptSignals...)
	defer stop()

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

	waitGroup, ctx := errgroup.WithContext(ctx)

	server.Run(ctx, waitGroup, cfg, grpcHandler)

	err = waitGroup.Wait()
	if err != nil {
		log.Error("error from wait group: " + err.Error())
	}
}
