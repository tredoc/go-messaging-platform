package query

import (
	"github.com/tredoc/go-messaging-platform/message/internal/domain/message"
	"github.com/tredoc/go-messaging-platform/message/internal/domain/status"
)

type MessageQuery struct {
	GetMessage GetMessageHandler
}

func NewMessageQuery(repo message.Repository) MessageQuery {
	return MessageQuery{
		GetMessage: NewGetMessageHandler(repo),
	}
}

type StatusQuery struct {
	GetMessageStatus GetMessageStatusHandler
}

func NewStatusQuery(repo status.Repository) StatusQuery {
	return StatusQuery{
		GetMessageStatus: NewGetMessageStatusHandler(repo),
	}
}
