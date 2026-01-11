package db

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Config struct to hold your database credentials
type Config struct {
	MongoURI string
	MongoDB  string
}

// Connect establishes a connection to the MongoDB database
func Connect(cfg Config) (*mongo.Client, *mongo.Database, error) {
	// Create a context with a 10-second timeout to prevent the app from freezing 
	// if the database is unreachable.
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Set client options using the URI from the config
	clientOption := options.Client().ApplyURI(cfg.MongoURI)

	// Attempt to connect to MongoDB
	client, err := mongo.Connect(ctx, clientOption)
	if err != nil {
		return nil, nil, fmt.Errorf("mongo connection failed: %w", err)
	}

	// Verify the connection is active by sending a Ping
	if err := client.Ping(ctx, nil); err != nil {
		return nil, nil, fmt.Errorf("mongo ping failed: %w", err)
	}

	// Select the specific database from the client
	database := client.Database(cfg.MongoDB)

	return client, database, nil
}

// Disconnect safely closes the MongoDB connection
func Disconnect(client *mongo.Client) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	
	return client.Disconnect(ctx)
}