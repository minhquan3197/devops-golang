package repositories

import (
	"context"
	"errors"
	"project-golang/internal/interfaces"
	"project-golang/internal/models"
	"project-golang/pkg/constants"
	"project-golang/pkg/mongodb"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/gommon/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type (
	userSchema = models.UserSchema
)

const (
	UsersCollection = models.UsersCollection
)

// Create data
func CreateUser(payload interfaces.CreateUser) error {
	collection := mongodb.GetDB().Collection(UsersCollection)
	user := userSchema{
		CreatedAt: time.Now(),
		ID:        uuid.New().String(),
		Password:  payload.Password,
		Username:  payload.Username,
	}
	_, err := collection.InsertOne(context.TODO(), user)
	if err != nil {
		return errors.New(constants.UserCreatedFailed)
	}
	return nil
}

// Remove data with conditon
func RemoveUser(condtion bson.M) error {
	collection := mongodb.GetDB().Collection(UsersCollection)
	collection.FindOneAndDelete(context.TODO(), condtion)
	return nil
}

// Update data with conditon
func UpdateUser(condition, payload bson.M) error {
	collection := mongodb.GetDB().Collection(UsersCollection)
	_, err := collection.UpdateOne(context.TODO(), condition, payload)
	if err != nil {
		return errors.New(constants.UserNotFound)
	}
	return nil
}

// FindOne data with condtion
func FindOneUser(condition bson.M) (userSchema, error) {
	collection := mongodb.GetDB().Collection(UsersCollection)
	var cursor userSchema
	result := collection.FindOne(context.TODO(), condition)
	err := result.Decode(&cursor)
	if err != nil {
		return cursor, errors.New(constants.UserNotFound)
	}
	return cursor, nil
}

// FindAll data with condtion
func FindAllUser(condition bson.M, options *options.FindOptions) ([]userSchema, error) {
	collection := mongodb.GetDB().Collection(UsersCollection)
	var results []userSchema
	pointer, err := collection.Find(context.TODO(), condition, options)
	if err != nil {
		log.Errorf("Unable to read the cursor : %v", err)
		return results, errors.New(constants.UserQueryIncorrect)
	}
	err = pointer.All(context.TODO(), &results)
	if err != nil {
		log.Errorf("Unable to read the cursor : %v", err)
		return results, errors.New(constants.UserQueryIncorrect)
	}
	return results, nil
}

// Count data with condtion
func CountUser(condtion bson.M) int64 {
	collection := mongodb.GetDB().Collection(UsersCollection)
	result, _ := collection.CountDocuments(context.TODO(), condtion)
	return result
}
