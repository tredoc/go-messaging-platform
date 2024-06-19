package repository

import (
	"context"
	"errors"
	"fmt"
	"github.com/tredoc/go-messaging-platform/template/internal/domain/template"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log/slog"
	"time"
)

const (
	templateCollection = "templates"
	templateDB         = "template"
)

type TemplateRepository struct {
	coll *mongo.Collection
	log  *slog.Logger
}

type TemplateDocument struct {
	UUID      string                `bson:"_id"`
	Content   string                `bson:"content"`
	TmplType  template.TemplateType `bson:"type"`
	CreatedAt time.Time             `bson:"created_at"`
}

func NewTemplateRepository(db *mongo.Client, log *slog.Logger) *TemplateRepository {
	return &TemplateRepository{
		coll: db.Database(templateDB).Collection(templateCollection),
		log:  log.With(slog.String("repository", "template")),
	}
}

func (r *TemplateRepository) Save(ctx context.Context, t *template.Template) error {
	r.log.Debug("TemplateRepository.Save", slog.Any("template", t))

	td := TemplateDocument{
		UUID:      t.UUID(),
		Content:   t.Content(),
		TmplType:  t.TmplType(),
		CreatedAt: t.CreatedAt(),
	}

	_, err := r.coll.InsertOne(ctx, td)
	if err != nil {
		r.log.Error("error saving template", slog.String("error", err.Error()))
		return err
	}
	return nil
}

func (r *TemplateRepository) FindByUUID(ctx context.Context, uuid string) (*template.Template, error) {
	r.log.Debug("TemplateRepository.FindByUUID", slog.String("uuid", uuid))
	var tmpl TemplateDocument

	filter := bson.D{{"_id", uuid}}
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
	r.log.Debug("TemplateRepository.FindByUUID", slog.String("uuid", uuid))

	filter := bson.D{{"_id", uuid}}
	_, err := r.coll.DeleteOne(ctx, filter)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			r.log.Debug("template not found", slog.String("uuid", uuid))
			return ErrNotFound
		}

		r.log.Error("error deleting template", slog.String("error", err.Error()))
		return fmt.Errorf("error deleting template: %w", err)
	}

	return nil
}
