package customer

import "go.mongodb.org/mongo-driver/mongo"

type Repository struct {
	Client *mongo.Client
}
