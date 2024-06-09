package repository

import "go.mongodb.org/mongo-driver/mongo"

const (
	messageCollection = "messages"
	messageDB         = "message"
)

type MessageRepository struct {
	coll *mongo.Collection
}

type MessageDocument struct{}

func NewMessageRepository(db *mongo.Client) *MessageRepository {
	return &MessageRepository{coll: db.Database(messageDB).Collection(messageCollection)}
}

func (r *MessageRepository) FindMessageStatusByUUID(uuid string) (string, error) {
	return "", nil
}
