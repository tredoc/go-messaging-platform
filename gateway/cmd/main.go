package main

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/tredoc/go-messaging-platform/gateway/internal/config"
	"github.com/tredoc/go-messaging-platform/gateway/internal/logger"
	"github.com/tredoc/go-messaging-platform/gateway/internal/server"
	"google.golang.org/protobuf/encoding/protojson"
	"log/slog"
	"os"
)

func main() {
	cfg, err := config.GetConfig()
	if err != nil {
		panic("configuration error: " + err.Error())
	}

	logger.SetupLogger(cfg.Env)

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux(runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.JSONPb{
		MarshalOptions: protojson.MarshalOptions{
			UseProtoNames:   true,
			EmitUnpopulated: true,
		},
	}))

	err = server.Run(ctx, cfg, mux)
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
}
