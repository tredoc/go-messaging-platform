package main

import (
	"context"
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/tredoc/go-messaging-platform/gateway/internal/config"
	"github.com/tredoc/go-messaging-platform/gateway/pb"
	"google.golang.org/grpc"
	"log"
	"net/http"
)

func main() {
	cfg, err := config.GetConfig()
	if err != nil {
		log.Fatalf("configuration error: %v", err)
	}

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}

	err = pb.RegisterOrchestratorServiceHandlerFromEndpoint(ctx, mux, "orchestrator:"+cfg.OrchestratorPort, opts)
	if err != nil {
		log.Fatalf("failed to start HTTP/2 gateway: %v", err)
	}

	err = pb.RegisterTemplateServiceHandlerFromEndpoint(ctx, mux, "template:"+cfg.TemplatePort, opts)
	if err != nil {
		log.Fatalf("failed to start HTTP/2 gateway: %v", err)
	}

	err = pb.RegisterMessageServiceHandlerFromEndpoint(ctx, mux, "message:"+cfg.MessagePort, opts)
	if err != nil {
		log.Fatalf("failed to start HTTP/2 gateway: %v", err)
	}

	log.Printf("Starting grpc gateway on port: %s\n", cfg.Port)
	http.ListenAndServe(fmt.Sprintf(":%s", cfg.Port), mux)

}
