package query

import (
	"context"
	"log"
)

type GetMessageStatus struct{}

type GetMessageStatusHandler struct{}

func NewGetMessageStatusHandler() GetMessageStatusHandler {
	return GetMessageStatusHandler{}
}

func (gh GetMessageStatusHandler) Handle(_ context.Context, _ GetMessageStatus) error {
	log.Println("Get message status query invoked")
	return nil
}
