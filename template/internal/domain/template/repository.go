package template

import "context"

type Repository interface {
	Save(context.Context, *Template) error
	DeleteByUUID(context.Context, string) error
	FindByUUID(context.Context, string) (*Template, error)
}
