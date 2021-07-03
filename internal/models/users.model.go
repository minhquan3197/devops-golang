package models

const (
	UsersCollection = "users"
)

// UserSchema schema
type UserSchema struct {
	ID        string `json:"_id" bson:"_id"`
	Username  string `json:"username" bson:"username"`
	Password  string `json:"password,omitempty" bson:"password"`
	CreatedAt int64  `json:"created_at" bson:"created_at"`
	UpdatedAt int64  `json:"updated_at" bson:"updated_at"`
}
