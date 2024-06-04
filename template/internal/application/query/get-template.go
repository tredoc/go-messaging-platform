package query

import (
	"context"
	"log"
)

type GetTemplate struct{}

type GetTemplateHandler struct{}

func NewGetTemplateHandler() GetTemplateHandler {
	return GetTemplateHandler{}
}

func (dt GetTemplateHandler) Handle(_ context.Context, _ GetTemplate) error {
	log.Println("Get template query invoked")
	return nil
}
