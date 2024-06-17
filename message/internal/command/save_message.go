package command

import (
	"context"
	"github.com/tredoc/go-messaging-platform/message/internal/domain/message"
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
}

func NewSaveMessageHandler(repo message.Repository) SaveMessageHandler {
	return SaveMessageHandler{
		repo: repo,
	}
}

func (gh SaveMessageHandler) Handle(ctx context.Context, cmd SaveMessage) error {
	dm, err := message.NewMessage(cmd.UUID, cmd.Message, cmd.TemplateUUID, cmd.Sender, cmd.Receiver, cmd.CreatedAt)
	if err != nil {
		return err
	}

	return gh.repo.SaveMessage(ctx, dm)
}
