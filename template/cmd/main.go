package main

import (
	"context"
	"github.com/tredoc/go-messaging-platform/template/internal/command"
	"github.com/tredoc/go-messaging-platform/template/internal/config"
	"github.com/tredoc/go-messaging-platform/template/internal/query"
	"github.com/tredoc/go-messaging-platform/template/internal/repository"
	"github.com/tredoc/go-messaging-platform/template/internal/server"
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

	repo := repository.NewTemplateRepository(mng)

	commands := command.NewCommand(repo)
	queries := query.NewQuery(repo)
	handler := server.NewGRPCHandler(commands, queries)

	err = server.Run(ctx, cfg, handler)
	if err != nil {
		log.Fatal(err)
	}
}
