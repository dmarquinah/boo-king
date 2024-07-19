package mongo

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectClient(dbURI string) (c *mongo.Client, err error) {
	// Set a timeout to allow the connection process to about if it takes too long
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	// Connect to the MongoDB server
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(dbURI))
	if err != nil {
		return nil, err
	}

	// Call the Ping method to verify that the connection has been stablished successfully.
	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, err
	}

	return client, nil
}
