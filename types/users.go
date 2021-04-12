package types

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// UserSchema schema
type UserSchema struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	UUID      string             `json:"uuid" bson:"uuid"`
	Username  string             `json:"username" bson:"username"`
	Password  string             `json:"password" bson:"password"`
	CreatedAt time.Time          `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt time.Time          `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}
