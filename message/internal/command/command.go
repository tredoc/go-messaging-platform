package command

import (
	"github.com/tredoc/go-messaging-platform/message/internal/domain/message"
	"github.com/tredoc/go-messaging-platform/message/internal/domain/status"
)

type MessageCommand struct {
	SaveMessage SaveMessageHandler
}

func NewMessageCommand(repo message.Repository) MessageCommand {
	return MessageCommand{
		SaveMessage: NewSaveMessageHandler(repo),
	}
}

type StatusCommand struct {
	SaveStatus SaveStatusHandler
}

func NewStatusCommand(repo status.Repository) StatusCommand {
	return StatusCommand{
		SaveStatus: NewSaveStatusHandler(repo),
	}
}
