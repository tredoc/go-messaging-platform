package main

import (
	"context"
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

	repo := repository.NewMessageRepository(mng)

	queries := query.NewQuery(repo)
	handler := server.NewGRPCHandler(queries)

	err = server.Run(ctx, cfg, handler)
	if err != nil {
		log.Fatal(err)
	}
}
