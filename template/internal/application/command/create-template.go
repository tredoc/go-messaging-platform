package command

import (
	"context"
	"log"
)

type CreateTemplate struct{}

type CreateTemplateHandler struct{}

func NewCreateTemplateHandler() CreateTemplateHandler {
	return CreateTemplateHandler{}
}

func (ct CreateTemplateHandler) Handle(_ context.Context, _ CreateTemplate) error {
	log.Println("Create template command invoked")
	return nil
}
