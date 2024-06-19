package command

import (
	"context"
	"fmt"
	"github.com/tredoc/go-messaging-platform/message/internal/domain/message"
	"log/slog"
	"time"
)

type SaveMessage struct {
	UUID         string
	Message      string
	TemplateUUID string
	Sender       string
	Receiver     string
	CreatedAt    time.Time
}

type SaveMessageHandler struct {
	repo message.Repository
	log  *slog.Logger
}

func NewSaveMessageHandler(repo message.Repository, log *slog.Logger) SaveMessageHandler {
	return SaveMessageHandler{
		repo: repo,
		log:  log.With(slog.String("command", "save_message")),
	}
}

func (gh SaveMessageHandler) Handle(ctx context.Context, cmd SaveMessage) error {
	gh.log.Debug("SaveMessageHandler.Handle", slog.Any("cmd", cmd))

	dm, err := message.NewMessage(cmd.UUID, cmd.Message, cmd.TemplateUUID, cmd.Sender, cmd.Receiver, cmd.CreatedAt)
	if err != nil {
		gh.log.Error("failed to create entity", slog.Any("error", err))
		return fmt.Errorf("save_message: %w", err)
	}

	err = gh.repo.SaveMessage(ctx, dm)
	if err != nil {
		gh.log.Error("failed to save message", slog.Any("error", err))
		return fmt.Errorf("save_message: %w", err)
	}

	return nil
}
