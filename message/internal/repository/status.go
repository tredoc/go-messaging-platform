package repository

import (
	"context"
	"github.com/tredoc/go-messaging-platform/message/internal/domain/status"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log/slog"
	"time"
)

const statusCollection = "statuses"

type StatusRepository struct {
	coll *mongo.Collection
	log  *slog.Logger
}

type StatusDocument struct {
	UUID        string               `bson:"_id"`
	MessageUUID string               `bson:"message_uuid"`
	Status      status.MessageStatus `bson:"status"`
	CreatedAt   time.Time            `bson:"created_at"`
}

func NewStatusRepository(db *mongo.Client, log *slog.Logger) *StatusRepository {
	return &StatusRepository{
		coll: db.Database(messageDB).Collection(statusCollection),
		log:  log.With(slog.String("repository", "status")),
	}
}

func (r StatusRepository) SaveStatus(_ context.Context, _ status.Status) error {
	return nil
}

func (r StatusRepository) FindStatusByUUID(ctx context.Context, uuid string) (status.Status, error) {
	r.log.Debug("StatusRepository.FindStatusByUUID", slog.String("uuid", uuid))

	var s StatusDocument

	filter := bson.D{{"_id", uuid}}
	err := r.coll.FindOne(ctx, filter).Decode(&s)
	if err != nil {
		return status.Status{}, ErrStatusNotFound
	}

	return status.UnmarshalFromDB(s.UUID, s.Status, s.MessageUUID, s.CreatedAt)
}

func (r StatusRepository) FindStatusByMessageUUID(ctx context.Context, uuid string) (status.Status, error) {
	r.log.Debug("StatusRepository.FindStatusByMessageUUID", slog.String("uuid", uuid))

	var s StatusDocument

	filter := bson.D{{"message_uuid", uuid}}
	err := r.coll.FindOne(ctx, filter).Decode(&s)
	if err != nil {
		return status.Status{}, ErrStatusNotFound
	}

	return status.UnmarshalFromDB(s.UUID, s.Status, s.MessageUUID, s.CreatedAt)
}
