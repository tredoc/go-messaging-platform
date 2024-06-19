package command

import (
	"context"
	"fmt"
	"github.com/tredoc/go-messaging-platform/message/internal/domain/status"
	"log/slog"
	"time"
)

type SaveStatus struct {
	UUID        string
	MessageUUID string
	Status      status.MessageStatus
	CreatedAt   time.Time
}

type SaveStatusHandler struct {
	repo status.Repository
	log  *slog.Logger
}

func NewSaveStatusHandler(repo status.Repository, log *slog.Logger) SaveStatusHandler {
	return SaveStatusHandler{
		repo: repo,
		log:  log.With(slog.String("command", "save_status")),
	}
}

func (sh SaveStatusHandler) Handle(ctx context.Context, cmd SaveStatus) error {
	sh.log.Debug("SaveStatusHandler.Handle", slog.Any("cmd", cmd))

	sts, err := status.NewStatus(cmd.UUID, cmd.Status, cmd.MessageUUID, cmd.CreatedAt)
	if err != nil {
		sh.log.Error("failed to create entity", slog.String("error", err.Error()))
		return err
	}

	err = sh.repo.SaveStatus(ctx, sts)
	if err != nil {
		sh.log.Error("failed to save status", slog.String("error", err.Error()))
		return fmt.Errorf("save_status: %w", err)
	}

	return nil
}
