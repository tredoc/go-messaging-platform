package query

import (
	"context"
	"github.com/tredoc/go-messaging-platform/message/internal/domain/message"
	"log"
)

type GetMessageStatus struct {
	UUID string
}

type GetMessageStatusHandler struct {
	repo message.Repository
}

func NewGetMessageStatusHandler(repo message.Repository) GetMessageStatusHandler {
	return GetMessageStatusHandler{
		repo: repo,
	}
}

func (gh GetMessageStatusHandler) Handle(_ context.Context, _ GetMessageStatus) error {
	log.Println("Get message status query invoked")
	return nil
}
