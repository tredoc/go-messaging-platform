package query

import (
	"context"
	"github.com/tredoc/go-messaging-platform/message/internal/domain/message"
)

type GetMessage struct {
	UUID string
}

type GetMessageHandler struct {
	repo message.Repository
}

func NewGetMessageHandler(repo message.Repository) GetMessageHandler {
	return GetMessageHandler{
		repo: repo,
	}
}

func (gh GetMessageHandler) Handle(ctx context.Context, q GetMessage) (message.Message, error) {
	dm, err := gh.repo.FindMessageByUUID(ctx, q.UUID)
	if err != nil {
		return message.Message{}, err
	}
	return dm, nil
}
