package query

import (
	"context"
	"errors"
	"github.com/tredoc/go-messaging-platform/message/internal/domain/status"
	"github.com/tredoc/go-messaging-platform/message/internal/repository"
)

type GetMessageStatus struct {
	UUID string
}

type GetMessageStatusHandler struct {
	repo status.Repository
}

func NewGetMessageStatusHandler(repo status.Repository) GetMessageStatusHandler {
	return GetMessageStatusHandler{
		repo: repo,
	}
}

func (gh GetMessageStatusHandler) Handle(ctx context.Context, q GetMessageStatus) (status.Status, error) {
	sts, err := gh.repo.FindStatusByUUID(ctx, q.UUID)
	if err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			return status.Status{}, repository.ErrNotFound
		}

		return status.Status{}, err
	}

	return sts, nil
}
