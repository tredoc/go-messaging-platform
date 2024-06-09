package command

import (
	"context"
	"github.com/tredoc/go-messaging-platform/template/internal/domain/template"
	"time"
)

type CreateTemplate struct {
	UUID      string
	Content   string
	TmplType  template.TmplType
	CreatedAt time.Time
}

type CreateTemplateHandler struct {
	repo template.Repository
}

func NewCreateTemplateHandler(repo template.Repository) CreateTemplateHandler {
	return CreateTemplateHandler{
		repo: repo,
	}
}

func (ct CreateTemplateHandler) Handle(ctx context.Context, cmd CreateTemplate) error {
	t, err := template.NewTemplate(cmd.UUID, cmd.Content, cmd.TmplType, cmd.CreatedAt)
	if err != nil {
		return err
	}

	return ct.repo.Save(ctx, t)
}
