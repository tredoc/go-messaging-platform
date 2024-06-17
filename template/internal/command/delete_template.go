package command

import (
	"context"
	"github.com/tredoc/go-messaging-platform/template/internal/domain/template"
)

type DeleteTemplate struct {
	UUID string
}

type DeleteTemplateHandler struct {
	repo template.Repository
}

func NewDeleteTemplateHandler(repo template.Repository) DeleteTemplateHandler {
	return DeleteTemplateHandler{
		repo: repo,
	}
}

func (dt DeleteTemplateHandler) Handle(ctx context.Context, cmd DeleteTemplate) error {
	return dt.repo.DeleteByUUID(ctx, cmd.UUID)
}
