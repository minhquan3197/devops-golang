package auth

import (
	"context"
	"errors"
	"net/http"
	"project-golang/internal/private/bcrypt"
	"project-golang/internal/private/jwt"
	"project-golang/third_party/mongodb"
	"project-golang/types"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
)

const (
	// NotFound message for not found error
	NotFound = "User not found"
	// Unauthorized message for login failed
	Unauthorized = "Unauthorized"
)

type (
	userSchema = types.UserSchema
)

// Login func login for user
func Login(req RequestLogin) (string, error) {
	var token string
	res, err := FindUserWithUsername(req.Username)
	if err != nil {
		return token, err
	}
	if !bcrypt.ComparePassword(req.Password, res.Password) {
		return token, echo.NewHTTPError(http.StatusUnauthorized, Unauthorized)
	}
	token = jwt.Encrypt("token", res.Username)
	return token, nil
}

// FindUserWithUsername func find user with username
// If use func in pkg users, will apply middleware become cycle import
func FindUserWithUsername(username string) (userSchema, error) {
	var cursor userSchema
	collection := mongodb.GetDB().Collection("users")
	result := collection.FindOne(context.TODO(), bson.M{"username": username})
	err := result.Decode(&cursor)
	if err != nil {
		return cursor, errors.New(NotFound)
	}
	return cursor, nil
}
