package query

import (
	"github.com/tredoc/go-messaging-platform/message/internal/domain/message"
	"github.com/tredoc/go-messaging-platform/message/internal/domain/status"
	"log/slog"
)

type MessageQuery struct {
	GetMessage GetMessageHandler
}

func NewMessageQuery(repo message.Repository, log *slog.Logger) MessageQuery {
	return MessageQuery{
		GetMessage: NewGetMessageHandler(repo, log),
	}
}

type StatusQuery struct {
	GetMessageStatus GetMessageStatusHandler
}

func NewStatusQuery(repo status.Repository, log *slog.Logger) StatusQuery {
	return StatusQuery{
		GetMessageStatus: NewGetMessageStatusHandler(repo, log),
	}
}
