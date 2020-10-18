package users

import (
	"context"
	"errors"
	"project-golang/internal/private/bcrypt"
	"project-golang/third_party/mongodb"
	"time"

	"github.com/labstack/gommon/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	// NotFound message for not found error
	NotFound = "User not found"
	// CreateFailed message for create error
	CreateFailed = "Create user failed"
)

// Create data
func Create(payload createUser) error {
	collection := mongodb.GetDB().Collection("users")
	hashedPassword := bcrypt.HashPassword(payload.Password)
	user := userSchema{
		CreatedAt: time.Now(),
		Password:  hashedPassword,
		Username:  payload.Username,
	}
	_, err := collection.InsertOne(context.TODO(), user)
	if err != nil {
		return errors.New(CreateFailed)
	}
	return nil
}

// FindOne data with condtion
func FindOne(condition bson.M) (userSchema, error) {
	collection := mongodb.GetDB().Collection("users")
	var cursor userSchema
	result := collection.FindOne(context.TODO(), condition)
	err := result.Decode(&cursor)
	if err != nil {
		return cursor, errors.New(NotFound)
	}
	return cursor, nil
}

// Count data with condtion
func Count(condtion bson.M) int64 {
	collection := mongodb.GetDB().Collection("users")
	result, _ := collection.CountDocuments(context.TODO(), condtion)
	return result
}

// Remove data with conditon
func Remove(condtion bson.M) error {
	collection := mongodb.GetDB().Collection("users")
	collection.FindOneAndDelete(context.TODO(), condtion)
	return nil
}

// Update data with ID
func Update(ID primitive.ObjectID, payload updateUser) error {
	collection := mongodb.GetDB().Collection("users")
	filter := bson.M{"_id": ID}
	update := bson.M{"$set": bson.M{
		"username":   payload.Username,
		"updated_at": time.Now(),
	}}
	_, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return errors.New(NotFound)
	}
	return nil
}

// Paginate func get list with paginate
func Paginate(limit, page int64, search string) (paginateUser, error) {
	collection := mongodb.GetDB().Collection("users")
	var results []userSchema
	var resp paginateUser
	conditionSearch := bson.M{"username": primitive.Regex{Pattern: "^" + search, Options: "i"}}
	if limit == 0 {
		limit = 12
	}
	if page == 0 {
		page = 1
	}
	findOptions := options.Find()
	findOptions.SetLimit(limit)
	findOptions.SetSkip(limit * (page - 1))
	pointer, err := collection.Find(context.TODO(), conditionSearch, findOptions)
	err = pointer.All(context.TODO(), &results)
	if err != nil {
		log.Errorf("Unable to read the cursor : %v", err)
		return resp, errors.New(NotFound)
	}
	resp = paginateUser{
		Data:    results,
		PerPage: limit,
		Total:   Count(conditionSearch),
	}
	return resp, nil
}
