package query

import (
	"context"
	"errors"
	"fmt"
	"github.com/tredoc/go-messaging-platform/message/internal/domain/message"
	"github.com/tredoc/go-messaging-platform/message/internal/repository"
	"log/slog"
)

type GetMessage struct {
	UUID string
}

type GetMessageHandler struct {
	repo message.Repository
	log  *slog.Logger
}

func NewGetMessageHandler(repo message.Repository, log *slog.Logger) GetMessageHandler {
	return GetMessageHandler{
		repo: repo,
		log:  log.With(slog.String("query", "get_message")),
	}
}

func (gh GetMessageHandler) Handle(ctx context.Context, q GetMessage) (message.Message, error) {
	gh.log.Debug("GetMessageHandler.Handle", slog.String("uuid", q.UUID))

	dm, err := gh.repo.FindMessageByUUID(ctx, q.UUID)
	if err != nil {
		if errors.Is(err, repository.ErrMsgNotFound) {
			gh.log.Debug("message not found", slog.String("uuid", q.UUID))
			return message.Message{}, repository.ErrMsgNotFound
		}

		gh.log.Error("failed to find message", slog.String("error", err.Error()))
		return message.Message{}, fmt.Errorf("get_message: %w", err)
	}

	return dm, nil
}
