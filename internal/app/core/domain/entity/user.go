package entity

import "time"

type User struct {
	Email     string     `json:"email" bson:"email"`
	Password  string     `json:"password" bson:"password"`
	CreatedAt *time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt *time.Time `json:"updated_at" bson:"updated_at"`
}
