package repository

import (
	"context"
	"github.com/tredoc/go-messaging-platform/message/internal/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const messageDB = "message"

func RunMongo(ctx context.Context, cfg config.Config) (*mongo.Client, error) {
	clientOptions := options.Client().ApplyURI(cfg.MongoURI)
	client, err := mongo.Connect(ctx, clientOptions)

	if err != nil {
		return nil, err
	}

	err = client.Ping(ctx, nil)

	if err != nil {
		return nil, err
	}

	return client, nil
}
