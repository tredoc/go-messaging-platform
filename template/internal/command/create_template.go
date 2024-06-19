package command

import (
	"context"
	"fmt"
	"github.com/tredoc/go-messaging-platform/template/internal/domain/template"
	"log/slog"
	"time"
)

type CreateTemplate struct {
	UUID      string
	Content   string
	TmplType  template.TemplateType
	CreatedAt time.Time
}

type CreateTemplateHandler struct {
	repo template.Repository
	log  *slog.Logger
}

func NewCreateTemplateHandler(repo template.Repository, log *slog.Logger) CreateTemplateHandler {
	return CreateTemplateHandler{
		repo: repo,
		log:  log.With(slog.String("command", "create_template")),
	}
}

func (ct CreateTemplateHandler) Handle(ctx context.Context, cmd CreateTemplate) error {
	ct.log.Debug("CreateTemplateHandler.Handle", slog.Any("cmd", cmd))

	t, err := template.NewTemplate(cmd.UUID, cmd.Content, cmd.TmplType, cmd.CreatedAt)
	if err != nil {
		ct.log.Error("failed to create entity", slog.String("error", err.Error()))
		return fmt.Errorf("create_template: %w", err)
	}

	err = ct.repo.Save(ctx, t)
	if err != nil {
		ct.log.Error("failed to save template", slog.String("error", err.Error()))
		return fmt.Errorf("create_template: %w", err)
	}

	return nil
}
