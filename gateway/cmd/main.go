package main

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/tredoc/go-messaging-platform/gateway/internal/config"
	"github.com/tredoc/go-messaging-platform/gateway/internal/logger"
	"github.com/tredoc/go-messaging-platform/gateway/internal/server"
	"golang.org/x/sync/errgroup"
	"google.golang.org/protobuf/encoding/protojson"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
)

var interruptSignals = []os.Signal{
	os.Interrupt,
	syscall.SIGTERM,
	syscall.SIGINT,
}

func main() {
	cfg, err := config.GetConfig()
	if err != nil {
		panic("configuration error: " + err.Error())
	}

	logger.SetupLogger(cfg.Env)
	log := slog.Default()

	ctx, stop := signal.NotifyContext(context.Background(), interruptSignals...)
	defer stop()

	mux := runtime.NewServeMux(runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.JSONPb{
		MarshalOptions: protojson.MarshalOptions{
			UseProtoNames:   true,
			EmitUnpopulated: true,
		},
	}))

	waitGroup, ctx := errgroup.WithContext(ctx)

	server.Run(ctx, waitGroup, cfg, mux)

	err = waitGroup.Wait()
	if err != nil {
		log.Error("error from wait group: " + err.Error())
	}

}
