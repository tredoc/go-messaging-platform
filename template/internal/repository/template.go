package repository

import (
	"context"
	"errors"
	"github.com/tredoc/go-messaging-platform/template/internal/domain/template"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

const (
	templateCollection = "templates"
	templateDB         = "template"
)

type TemplateRepository struct {
	coll *mongo.Collection
}

type TemplateDocument struct {
	UUID      string            `bson:"uuid"`
	Content   string            `bson:"content"`
	TmplType  template.TmplType `bson:"type"`
	CreatedAt time.Time         `bson:"created_at"`
}

func NewTemplateRepository(db *mongo.Client) *TemplateRepository {
	return &TemplateRepository{coll: db.Database(templateDB).Collection(templateCollection)}
}

func (r *TemplateRepository) Save(ctx context.Context, t *template.Template) error {
	td := TemplateDocument{
		UUID:      t.UUID(),
		Content:   t.Content(),
		TmplType:  t.TmplType(),
		CreatedAt: t.CreatedAt(),
	}

	_, err := r.coll.InsertOne(ctx, td)
	if err != nil {
		return err
	}
	return nil
}

func (r *TemplateRepository) FindByUUID(ctx context.Context, uuid string) (*template.Template, error) {
	var tmpl TemplateDocument

	filter := bson.D{{"uuid", uuid}}
	err := r.coll.FindOne(ctx, filter).Decode(&tmpl)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, ErrNotFound
		}
		return nil, ErrNotFound
	}

	return template.UnmarshalFromDB(tmpl.UUID, tmpl.Content, tmpl.TmplType, tmpl.CreatedAt)
}

func (r *TemplateRepository) DeleteByUUID(ctx context.Context, uuid string) error {
	filter := bson.D{{"uuid", uuid}}
	_, err := r.coll.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}
	return nil
}
