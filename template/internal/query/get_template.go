package query

import (
	"context"
	"fmt"
	"github.com/tredoc/go-messaging-platform/template/internal/domain/template"
	"log/slog"
)

type GetTemplate struct {
	UUID string
}

type GetTemplateHandler struct {
	repo template.Repository
	log  *slog.Logger
}

func NewGetTemplateHandler(repo template.Repository, log *slog.Logger) GetTemplateHandler {
	return GetTemplateHandler{
		repo: repo,
		log:  log.With(slog.String("query", "get_template")),
	}
}

func (gh GetTemplateHandler) Handle(ctx context.Context, cmd GetTemplate) (*template.Template, error) {
	gh.log.Debug("NewGetTemplateHandler.Handle", slog.Any("cmd", cmd))

	t, err := gh.repo.FindByUUID(ctx, cmd.UUID)
	if err != nil {
		gh.log.Error("failed to get template", slog.String("error", err.Error()))
		return nil, fmt.Errorf("create_template: %w", err)
	}

	return t, nil
}
