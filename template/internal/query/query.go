package query

type Query struct {
	GetTemplate GetTemplateHandler
}

func NewQuery() Query {
	return Query{
		GetTemplate: NewGetTemplateHandler(),
	}
}
