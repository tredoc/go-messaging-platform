package server

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

func Run(ctx context.Context, cfg config.Config, mux *runtime.ServeMux) error {
	opts := []grpc.DialOption{grpc.WithInsecure()}

	err := pb.RegisterOrchestratorServiceHandlerFromEndpoint(ctx, mux, cfg.OrchestratorAddr, opts)
	if err != nil {
		return fmt.Errorf("failed to start HTTP/2 gateway: %w", err)
	}

	err = pb.RegisterTemplateServiceHandlerFromEndpoint(ctx, mux, cfg.TemplateAddr, opts)
	if err != nil {
		return fmt.Errorf("failed to start HTTP/2 gateway: %w", err)
	}

	err = pb.RegisterMessageServiceHandlerFromEndpoint(ctx, mux, cfg.MessageAddr, opts)
	if err != nil {
		return fmt.Errorf("failed to start HTTP/2 gateway: %w", err)
	}

	log.Printf("Starting grpc gateway on port: %s\n", cfg.Port)
	return http.ListenAndServe(fmt.Sprintf(":%s", cfg.Port), mux)
}
