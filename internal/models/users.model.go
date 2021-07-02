package models

const (
	UsersCollection = "users"
)

// UserSchema schema
type UserSchema struct {
	ID        string `json:"_id" bson:"_id"`
	Username  string `json:"username" bson:"username"`
	Password  string `json:"password" bson:"password"`
	CreatedAt int64  `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt int64  `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}
