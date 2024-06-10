package repository

import (
	"context"
	"github.com/tredoc/go-messaging-platform/message/internal/domain/status"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

const statusCollection = "statuses"

type StatusRepository struct {
	coll *mongo.Collection
}

type StatusDocument struct {
	UUID        string           `bson:"_id"`
	MessageUUID string           `bson:"message_uuid"`
	Status      status.MsgStatus `bson:"status"`
	CreatedAt   time.Time        `bson:"created_at"`
}

func NewStatusRepository(db *mongo.Client) *StatusRepository {
	return &StatusRepository{
		coll: db.Database(messageDB).Collection(statusCollection)}
}

func (r StatusRepository) SaveStatus(ctx context.Context, s status.Status) error {
	return nil
}

func (r StatusRepository) FindStatusByUUID(ctx context.Context, uuid string) (status.Status, error) {
	var s StatusDocument

	filter := bson.D{{"_id", uuid}}
	err := r.coll.FindOne(ctx, filter).Decode(&s)
	if err != nil {
		return status.Status{}, ErrNotFound
	}

	return status.UnmarshalFromDB(s.UUID, s.Status, s.MessageUUID, s.CreatedAt)
}

func (r StatusRepository) FindStatusByMessageUUID(ctx context.Context, uuid string) (status.Status, error) {
	var s StatusDocument

	filter := bson.D{{"message_uuid", uuid}}
	err := r.coll.FindOne(ctx, filter).Decode(&s)
	if err != nil {
		return status.Status{}, ErrNotFound
	}

	return status.UnmarshalFromDB(s.UUID, s.Status, s.MessageUUID, s.CreatedAt)
}
