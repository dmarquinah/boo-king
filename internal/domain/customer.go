package domain

import "time"

type Customer struct {
	Name      string     `json:"name" bson:"name"`
	Email     string     `json:"email" bson:"email"`
	CreatedAt *time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt *time.Time `json:"updated_at" bson:"updated_at"`
}
