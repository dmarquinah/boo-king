package user

import "go.mongodb.org/mongo-driver/mongo"

type Repository struct {
	Database       *mongo.Database
	UserCollection string `default:"user"`
}
