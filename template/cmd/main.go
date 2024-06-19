package main

import (
	"context"
	"github.com/tredoc/go-messaging-platform/template/internal/command"
	"github.com/tredoc/go-messaging-platform/template/internal/config"
	"github.com/tredoc/go-messaging-platform/template/internal/logger"
	"github.com/tredoc/go-messaging-platform/template/internal/query"
	"github.com/tredoc/go-messaging-platform/template/internal/repository"
	"github.com/tredoc/go-messaging-platform/template/internal/server"
	"log/slog"
	"os"
)

func main() {
	cfg, err := config.GetConfig()
	if err != nil {
		panic("configuration error: " + err.Error())
	}

	log := logger.GetLogger(cfg.Env)

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mng, err := repository.RunMongo(cfg)
	if err != nil {
		log.Error("mongo connection error: "+err.Error(), slog.Any("env", cfg.Env))
		os.Exit(1)
	}

	log.Info("Connected to MongoDB", slog.Any("environment", cfg.Env))

	repo := repository.NewTemplateRepository(mng, log)

	commands := command.NewCommand(repo, log)
	queries := query.NewQuery(repo, log)
	handler := server.NewGRPCHandler(commands, queries)

	err = server.Run(ctx, cfg, handler, log)
	if err != nil {
		log.Error("message server run error: "+err.Error(), slog.Any("env", cfg.Env))
		os.Exit(1)
	}
}
