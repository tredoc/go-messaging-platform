package message

import "context"

type Repository interface {
	SaveMessage(context.Context, Message) error
	FindMessageByUUID(context.Context, string) (Message, error)
}
