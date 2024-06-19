package main

import (
	"context"
	"github.com/tredoc/go-messaging-platform/message/internal/command"
	"github.com/tredoc/go-messaging-platform/message/internal/config"
	"github.com/tredoc/go-messaging-platform/message/internal/logger"
	"github.com/tredoc/go-messaging-platform/message/internal/query"
	"github.com/tredoc/go-messaging-platform/message/internal/repository"
	"github.com/tredoc/go-messaging-platform/message/internal/server"
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

	err = server.Run(ctx, cfg, handler, log)
	if err != nil {
		log.Error("message server run error: "+err.Error(), slog.Any("env", cfg.Env))
		os.Exit(1)
	}
}
