package command

import (
	"github.com/tredoc/go-messaging-platform/message/internal/domain/message"
	"github.com/tredoc/go-messaging-platform/message/internal/domain/status"
	"log/slog"
)

type MessageCommand struct {
	SaveMessage SaveMessageHandler
}

func NewMessageCommand(repo message.Repository, log *slog.Logger) MessageCommand {
	return MessageCommand{
		SaveMessage: NewSaveMessageHandler(repo, log),
	}
}

type StatusCommand struct {
	SaveStatus SaveStatusHandler
}

func NewStatusCommand(repo status.Repository, log *slog.Logger) StatusCommand {
	return StatusCommand{
		SaveStatus: NewSaveStatusHandler(repo, log),
	}
}
