package customer

import (
	"context"
	"fmt"
	"log"

	"github.com/dmarquinah/boo-king/internal/domain"
)

func (r Repository) Insert(customer domain.Customer) (id interface{}, err error) {

	collection := r.Client.Database("boo-king").Collection("customers")
	insertResult, err := collection.InsertOne(context.Background(), customer)
	if err != nil {
		log.Println(err.Error())
		return nil, fmt.Errorf("error inserting customer: %w", err)
	}

	return insertResult.InsertedID, nil
}
