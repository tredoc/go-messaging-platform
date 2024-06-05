package command

import (
	"context"
	"log"
)

type DeleteTemplate struct{}

type DeleteTemplateHandler struct{}

func NewDeleteTemplateHandler() DeleteTemplateHandler {
	return DeleteTemplateHandler{}
}

func (dt DeleteTemplateHandler) Handle(_ context.Context, _ DeleteTemplate) error {
	log.Println("Delete template command invoked")
	return nil
}
