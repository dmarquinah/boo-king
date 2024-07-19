package user

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/dmarquinah/boo-king/internal/app/core/domain/dto"
	"github.com/dmarquinah/boo-king/internal/app/core/domain/entity"
)

func (r Repository) Insert(user entity.User) (userDTO *dto.FindUserDTO, err error) {
	now := time.Now().UTC()
	user.CreatedAt = &now
	user.UpdatedAt = &now

	fmt.Println("Collection:", r.UserCollection)
	insertResult, err := r.Database.Collection("user").InsertOne(context.Background(), user)
	if err != nil {
		log.Println(err.Error())
		return nil, fmt.Errorf("error inserting customer: %w", err)
	}

	return &dto.FindUserDTO{Id: insertResult.InsertedID}, nil
}
