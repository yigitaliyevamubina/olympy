package storage

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoClient struct {
	client *mongo.Client
}

func NewMongoClient(uri string) (*MongoClient, error) {
	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return nil, fmt.Errorf("error connecting to MongoDB: %v", err)
	}

	if err := client.Ping(context.Background(), nil); err != nil {
		return nil, fmt.Errorf("error pinging MongoDB: %v", err)
	}

	fmt.Println("Successfully connected to MongoDB")
	return &MongoClient{client: client}, nil
}

func (m *MongoClient) InsertEvent(ctx context.Context, event bson.D) error {
	collection := m.client.Database("streaming").Collection("events")
	_, err := collection.InsertOne(ctx, event)
	return err
}
