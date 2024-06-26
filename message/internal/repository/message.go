package repository

import (
	"context"
	"errors"
	"github.com/tredoc/go-messaging-platform/message/internal/domain/message"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log/slog"
	"time"
)

const messageCollection = "messages"

type MessageRepository struct {
	coll *mongo.Collection
	log  *slog.Logger
}

type MessageDocument struct {
	UUID         string    `bson:"_id"`
	Message      string    `bson:"message"`
	TemplateUUID string    `bson:"template_uuid"`
	Sender       string    `bson:"sender"`
	Receiver     string    `bson:"receiver"`
	CreatedAt    time.Time `bson:"created_at"`
}

func NewMessageRepository(db *mongo.Client, log *slog.Logger) *MessageRepository {
	return &MessageRepository{
		coll: db.Database(messageDB).Collection(messageCollection),
		log:  log.With(slog.String("repository", "message")),
	}
}

func (r MessageRepository) SaveMessage(ctx context.Context, msg message.Message) error {
	r.log.Debug("MessageRepository.SaveMessage", slog.Any("message", msg))

	_, err := r.coll.InsertOne(ctx, MessageDocument{
		UUID:         msg.UUID(),
		Message:      msg.Message(),
		TemplateUUID: msg.TemplateUUID(),
		Sender:       msg.Sender(),
		Receiver:     msg.Receiver(),
		CreatedAt:    msg.CreatedAt(),
	})

	return err
}

func (r MessageRepository) FindMessageByUUID(ctx context.Context, uuid string) (message.Message, error) {
	r.log.Debug("MessageRepository.FindMessageByUUID", slog.String("uuid", uuid))

	var msg MessageDocument

	filter := bson.D{{"_id", uuid}}
	err := r.coll.FindOne(ctx, filter).Decode(&msg)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return message.Message{}, ErrMsgNotFound
		}

		return message.Message{}, err
	}

	return message.UnmarshalFromDB(msg.UUID, msg.Message, msg.TemplateUUID, msg.Sender, msg.Receiver, msg.CreatedAt)
}
