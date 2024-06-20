package main

import (
	"context"
	"github.com/tredoc/go-messaging-platform/message/internal/command"
	"github.com/tredoc/go-messaging-platform/message/internal/config"
	"github.com/tredoc/go-messaging-platform/message/internal/logger"
	"github.com/tredoc/go-messaging-platform/message/internal/query"
	"github.com/tredoc/go-messaging-platform/message/internal/repository"
	"github.com/tredoc/go-messaging-platform/message/internal/server"
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

	log := logger.GetLogger(cfg.Env)

	ctx, stop := signal.NotifyContext(context.Background(), interruptSignals...)
	defer stop()

	mng, err := repository.RunMongo(ctx, cfg)
	if err != nil {
		log.Error("mongo connection error: "+err.Error(), slog.Any("env", cfg.Env))
		os.Exit(1)
	}

	log.Info("Connected to MongoDB", slog.Any("environment", cfg.Env))

	msgRepo := repository.NewMessageRepository(mng, log)
	stsRepo := repository.NewStatusRepository(mng, log)

	msgQueries := query.NewMessageQuery(msgRepo, log)
	msgCommands := command.NewMessageCommand(msgRepo, log)

	stsQueries := query.NewStatusQuery(stsRepo, log)
	stsCommands := command.NewStatusCommand(stsRepo, log)

	handler := server.NewGRPCHandler(msgQueries, msgCommands, stsQueries, stsCommands)

	waitGroup, ctx := errgroup.WithContext(ctx)

	server.Run(ctx, waitGroup, cfg, handler, log)

	err = waitGroup.Wait()
	if err != nil {
		log.Error("error from wait group")
	}
}
