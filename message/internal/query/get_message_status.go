package query

import (
	"context"
	"errors"
	"fmt"
	"github.com/tredoc/go-messaging-platform/message/internal/domain/status"
	"github.com/tredoc/go-messaging-platform/message/internal/repository"
	"log/slog"
)

type GetMessageStatus struct {
	UUID string
}

type GetMessageStatusHandler struct {
	repo status.Repository
	log  *slog.Logger
}

func NewGetMessageStatusHandler(repo status.Repository, log *slog.Logger) GetMessageStatusHandler {
	return GetMessageStatusHandler{
		repo: repo,
		log:  log.With(slog.String("query", "get_message_status")),
	}
}

func (gh GetMessageStatusHandler) Handle(ctx context.Context, q GetMessageStatus) (status.Status, error) {
	gh.log.Debug("GetMessageStatusHandler.Handle", slog.String("uuid", q.UUID))

	sts, err := gh.repo.FindStatusByUUID(ctx, q.UUID)
	if err != nil {
		if errors.Is(err, repository.ErrStatusNotFound) {
			gh.log.Debug("status not found", slog.String("uuid", q.UUID))
			return status.Status{}, repository.ErrStatusNotFound
		}

		gh.log.Error("failed to find status", slog.String("error", err.Error()))
		return status.Status{}, fmt.Errorf("get_message_status: %w", err)
	}

	return sts, nil
}
