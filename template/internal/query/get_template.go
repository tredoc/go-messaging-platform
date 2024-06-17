package query

import (
	"context"
	"github.com/tredoc/go-messaging-platform/template/internal/domain/template"
)

type GetTemplate struct {
	UUID string
}

type GetTemplateHandler struct {
	repo template.Repository
}

func NewGetTemplateHandler(repo template.Repository) GetTemplateHandler {
	return GetTemplateHandler{
		repo: repo,
	}
}

func (gh GetTemplateHandler) Handle(ctx context.Context, cmd GetTemplate) (*template.Template, error) {
	return gh.repo.FindByUUID(ctx, cmd.UUID)
}
