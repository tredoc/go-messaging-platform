package logger

import (
	"github.com/tredoc/go-messaging-platform/orchestrator/internal/config"
	"log/slog"
	"os"
)

type TeeWriter struct {
	stdout *os.File
	file   *os.File
}

func (t *TeeWriter) Write(p []byte) (n int, err error) {
	n, err = t.stdout.Write(p)
	if err != nil {
		return n, err
	}
	n, err = t.file.Write(p)
	return n, err
}

func SetupLogger(env config.Environment) {
	file, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}

	writer := &TeeWriter{
		stdout: os.Stdout,
		file:   file,
	}

	var log *slog.Logger
	switch env {
	case config.Development:
		log = slog.New(slog.NewTextHandler(writer, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case config.Production:
		log = slog.New(slog.NewJSONHandler(writer, &slog.HandlerOptions{Level: slog.LevelInfo}))
	}

	slog.SetDefault(log)
}
