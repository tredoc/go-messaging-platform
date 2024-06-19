package command

import (
	"context"
	"fmt"
	"github.com/tredoc/go-messaging-platform/template/internal/domain/template"
	"log/slog"
)

type DeleteTemplate struct {
	UUID string
}

type DeleteTemplateHandler struct {
	repo template.Repository
	log  *slog.Logger
}

func NewDeleteTemplateHandler(repo template.Repository, log *slog.Logger) DeleteTemplateHandler {
	return DeleteTemplateHandler{
		repo: repo,
		log:  log.With(slog.String("command", "delete_template")),
	}
}

func (dt DeleteTemplateHandler) Handle(ctx context.Context, cmd DeleteTemplate) error {
	dt.log.Debug("DeleteTemplateHandler.Handle", slog.Any("cmd", cmd))
	err := dt.repo.DeleteByUUID(ctx, cmd.UUID)
	if err != nil {
		dt.log.Error("failed to delete template", slog.String("error", err.Error()))
		return fmt.Errorf("delete_template: %w", err)
	}

	return nil
}
