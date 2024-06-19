package command

import (
	"github.com/tredoc/go-messaging-platform/template/internal/domain/template"
	"log/slog"
)

type Command struct {
	CreateTemplate CreateTemplateHandler
	DeleteTemplate DeleteTemplateHandler
}

func NewCommand(repo template.Repository, log *slog.Logger) Command {
	return Command{
		CreateTemplate: NewCreateTemplateHandler(repo, log),
		DeleteTemplate: NewDeleteTemplateHandler(repo, log),
	}
}
