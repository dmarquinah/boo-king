package user

import (
	"context"
	"errors"
	"fmt"

	"github.com/dmarquinah/boo-king/internal/app/core/domain/dto"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (r Repository) FindOne(id string) (userDTO *dto.FindUserDTO, err error) {
	userID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, fmt.Errorf("invalid id (%s): %w", id, err)
	}

	err = r.Database.Collection("user").FindOne(context.Background(), bson.M{"_id": userID}).Decode(&userDTO)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, fmt.Errorf("no items found: %w", err)
		}
		var commandErr *mongo.CommandError
		if errors.As(err, &commandErr) && commandErr.HasErrorLabel("NetworkTimeout") {
			return nil, fmt.Errorf("timeout: %w", err)
		}
		return nil, err
	}

	userDTO.Id = userID.Hex()
	return userDTO, nil
}

func (r Repository) FindAll() (userDTO *[]dto.FindUserDTO, err error) {

	findOptions := options.Find()

	cur, err := r.Database.Collection(r.UserCollection).Find(context.Background(), bson.D{}, findOptions)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, fmt.Errorf("no items found: %w", err)
		}
		var commandErr *mongo.CommandError
		if errors.As(err, &commandErr) && commandErr.HasErrorLabel("NetworkTimeout") {
			return nil, fmt.Errorf("timeout: %w", err)
		}
		return nil, err
	}

	var results []dto.FindUserDTO
	if err = cur.All(context.Background(), &results); err != nil {
		return nil, fmt.Errorf("can't read cursor: %w", err)
	}

	for _, result := range results {
		res, _ := bson.MarshalExtJSON(result, false, false)
		fmt.Println(res)
	}

	return &results, nil
}
