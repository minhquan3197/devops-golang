package models

import (
	"time"
)

const (
	UsersCollection = "users"
)

// UserSchema schema
type UserSchema struct {
	ID        string    `json:"_id" bson:"_id"`
	Username  string    `json:"username" bson:"username"`
	Password  string    `json:"password" bson:"password"`
	CreatedAt time.Time `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}
