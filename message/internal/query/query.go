package query

import "github.com/tredoc/go-messaging-platform/message/internal/domain/message"

type Query struct {
	GetMessageStatus GetMessageStatusHandler
}

func NewQuery(repo message.Repository) Query {
	return Query{GetMessageStatus: NewGetMessageStatusHandler(repo)}
}
