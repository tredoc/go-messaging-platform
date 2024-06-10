package main

import (
	"context"
	"github.com/tredoc/go-messaging-platform/message/internal/command"
	"github.com/tredoc/go-messaging-platform/message/internal/config"
	"github.com/tredoc/go-messaging-platform/message/internal/query"
	"github.com/tredoc/go-messaging-platform/message/internal/repository"
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

	mng, err := repository.RunMongo(cfg)
	if err != nil {
		log.Fatal(err)
	}

	msgRepo := repository.NewMessageRepository(mng)
	stsRepo := repository.NewStatusRepository(mng)

	msgQueries := query.NewMessageQuery(msgRepo)
	msgCommands := command.NewMessageCommand(msgRepo)

	stsQueries := query.NewStatusQuery(stsRepo)
	stsCommands := command.NewStatusCommand(stsRepo)

	handler := server.NewGRPCHandler(msgQueries, msgCommands, stsQueries, stsCommands)

	err = server.Run(ctx, cfg, handler)
	if err != nil {
		log.Fatal(err)
	}
}
