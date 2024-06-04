package application

import (
	"github.com/tredoc/go-messaging-platform/template/internal/application/command"
	"github.com/tredoc/go-messaging-platform/template/internal/application/query"
)

type Application struct {
	Commands Commands
	Queries  Queries
}

func New() Application {
	return Application{
		Commands: Commands{
			CreateTemplate: command.NewCreateTemplateHandler(),
			DeleteTemplate: command.NewDeleteTemplateHandler(),
		},
		Queries: Queries{
			GetTemplate: query.NewGetTemplateHandler(),
		},
	}
}

type Commands struct {
	CreateTemplate command.CreateTemplateHandler
	DeleteTemplate command.DeleteTemplateHandler
}

type Queries struct {
	GetTemplate query.GetTemplateHandler
}
