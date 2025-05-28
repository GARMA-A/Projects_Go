package types

import (
	"context"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	CollectionName = "Todos"
	DatabaseName   = "Golang"
	GlobalClient   *mongo.Client
)

func ConnectToMongoDB() (*mongo.Client, error) {
	MonogDBURI := os.Getenv("DB_URL")
	clientOptions := options.Client().ApplyURI(MonogDBURI)

	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return nil, err
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		return nil, err
	}

	return client, nil
}

func CurrentCollection(client *mongo.Client) *mongo.Collection {
	return client.Database(DatabaseName).Collection(CollectionName)
}
