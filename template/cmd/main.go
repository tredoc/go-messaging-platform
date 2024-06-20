package main

import (
	"context"
	"github.com/tredoc/go-messaging-platform/template/internal/command"
	"github.com/tredoc/go-messaging-platform/template/internal/config"
	"github.com/tredoc/go-messaging-platform/template/internal/logger"
	"github.com/tredoc/go-messaging-platform/template/internal/query"
	"github.com/tredoc/go-messaging-platform/template/internal/repository"
	"github.com/tredoc/go-messaging-platform/template/internal/server"
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

	repo := repository.NewTemplateRepository(mng, log)

	commands := command.NewCommand(repo, log)
	queries := query.NewQuery(repo, log)
	handler := server.NewGRPCHandler(commands, queries)

	waitGroup, ctx := errgroup.WithContext(ctx)

	server.Run(ctx, waitGroup, cfg, handler, log)

	err = waitGroup.Wait()
	if err != nil {
		log.Error("error from wait group: " + err.Error())
	}
}
