package query

import "github.com/tredoc/go-messaging-platform/template/internal/domain/template"

type Query struct {
	GetTemplate GetTemplateHandler
}

func NewQuery(repo template.Repository) Query {
	return Query{
		GetTemplate: NewGetTemplateHandler(repo),
	}
}
