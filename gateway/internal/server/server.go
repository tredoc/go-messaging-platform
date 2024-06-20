package server

import (
	"context"
	"errors"
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/tredoc/go-messaging-platform/gateway/internal/config"
	"github.com/tredoc/go-messaging-platform/gateway/pb"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"log/slog"
	"net/http"
)

func Run(ctx context.Context, waitGroup *errgroup.Group, cfg config.Config, mux *runtime.ServeMux) {
	log := slog.Default()

	opts := []grpc.DialOption{grpc.WithInsecure()}

	err := pb.RegisterOrchestratorServiceHandlerFromEndpoint(ctx, mux, cfg.OrchestratorAddr, opts)
	if err != nil {
		log.Error("failed to start HTTP/2 gateway", slog.Any("error", err))
	}

	err = pb.RegisterTemplateServiceHandlerFromEndpoint(ctx, mux, cfg.TemplateAddr, opts)
	if err != nil {
		log.Error("failed to start HTTP/2 gateway", slog.Any("error", err))
	}

	err = pb.RegisterMessageServiceHandlerFromEndpoint(ctx, mux, cfg.MessageAddr, opts)
	if err != nil {
		log.Error("failed to start HTTP/2 gateway", slog.Any("error", err))
	}

	//slog.Info("Starting gateway server", slog.Any("environment", cfg.Env), slog.Int("port", cfg.Port))
	//return http.ListenAndServe(fmt.Sprintf(":%d", cfg.Port), HttpLogger(mux))

	httpServer := &http.Server{
		Handler: HttpLogger(mux),
		Addr:    fmt.Sprintf(":%d", cfg.Port),
	}

	waitGroup.Go(func() error {
		log.Info("Start HTTP gateway server", slog.Int("port", cfg.Port))
		err = httpServer.ListenAndServe()
		if err != nil {
			if errors.Is(err, http.ErrServerClosed) {
				return nil
			}
			log.Error("HTTP gateway server failed to serve")
			return err
		}
		return nil
	})

	waitGroup.Go(func() error {
		<-ctx.Done()
		log.Info("Graceful shutdown HTTP gateway server")

		err := httpServer.Shutdown(context.Background())
		if err != nil {
			log.Error("failed to shutdown HTTP gateway server")
			return err
		}

		log.Info("HTTP gateway server is stopped")
		return nil
	})
}
