package main

import (
	"context"
	"github.com/tredoc/go-messaging-platform/template/internal/application"
	"github.com/tredoc/go-messaging-platform/template/internal/config"
	"github.com/tredoc/go-messaging-platform/template/internal/handler"
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

	app := application.New()
	grpcHandler := handler.NewGRPCTemplateHandler(app)
	GRPCServer := server.NewGRPCServer(grpcHandler)

	err = server.Run(ctx, GRPCServer, cfg)
	if err != nil {
		log.Fatal(err)
	}
}
