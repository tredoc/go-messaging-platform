package application

import "github.com/tredoc/go-messaging-platform/message/internal/application/query"

type Application struct {
	Commands Commands
	Queries  Queries
}

func New() Application {
	return Application{
		Commands: Commands{},
		Queries: Queries{
			GetMessageStatus: query.NewGetMessageStatusHandler(),
		},
	}
}

type Commands struct{}

type Queries struct {
	GetMessageStatus query.GetMessageStatusHandler
}
