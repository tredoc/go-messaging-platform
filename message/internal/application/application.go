package application

import query "github.com/tredoc/go-messaging-platform/message/internal/application/queries"

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
