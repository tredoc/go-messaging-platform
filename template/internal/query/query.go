package query

import (
	"github.com/tredoc/go-messaging-platform/template/internal/domain/template"
	"log/slog"
)

type Query struct {
	GetTemplate GetTemplateHandler
}

func NewQuery(repo template.Repository, log *slog.Logger) Query {
	return Query{
		GetTemplate: NewGetTemplateHandler(repo, log),
	}
}
