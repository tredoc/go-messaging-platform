package command

import (
	"context"
	"github.com/tredoc/go-messaging-platform/message/internal/domain/status"
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
}

func NewSaveStatusHandler(repo status.Repository) SaveStatusHandler {
	return SaveStatusHandler{
		repo: repo,
	}
}

func (sh SaveStatusHandler) Handle(ctx context.Context, cmd SaveStatus) error {
	sts, err := status.NewStatus(cmd.UUID, cmd.Status, cmd.MessageUUID, cmd.CreatedAt)
	if err != nil {
		return err
	}

	return sh.repo.SaveStatus(ctx, sts)
}
