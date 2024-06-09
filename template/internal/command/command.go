package command

import "github.com/tredoc/go-messaging-platform/template/internal/domain/template"

type Command struct {
	CreateTemplate CreateTemplateHandler
	DeleteTemplate DeleteTemplateHandler
}

func NewCommand(repo template.Repository) Command {
	return Command{
		CreateTemplate: NewCreateTemplateHandler(repo),
		DeleteTemplate: NewDeleteTemplateHandler(repo),
	}
}
