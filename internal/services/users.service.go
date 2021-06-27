package services

import (
	"project-golang/internal/interfaces"
	"project-golang/internal/models"
	"project-golang/internal/repositories"
	"project-golang/pkg/bcrypt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type (
	userSchema = models.UserSchema
)

// RegisterUser func create new user
func RegisterUser(payload interfaces.CreateUser) error {
	payload.Password = bcrypt.HashPassword(payload.Password)
	err := repositories.CreateUser(payload)
	if err != nil {
		return err
	}
	return nil
}

// PaginateUsers func get list with paginate
func PaginateUsers(limit, page int64, search string) (interfaces.PaginateUser, error) {
	var resp interfaces.PaginateUser
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
	results, err := repositories.FindAllUser(conditionSearch, findOptions)
	if err != nil {
		return resp, err
	}
	countUser := repositories.CountUser(conditionSearch)
	resp = interfaces.PaginateUser{
		Data:    results,
		PerPage: limit,
		Total:   countUser,
	}
	return resp, nil
}

// FindUserByUsername func find user by username
func FindUserByUsername(username string) (userSchema, error) {
	conditionSearch := bson.M{"username": username}
	result, err := repositories.FindOneUser(conditionSearch)
	if err != nil {
		return result, err
	}
	return result, nil
}

// FindUserByID func find user by uuid
func FindUserByID(uuid string) (userSchema, error) {
	conditionSearch := bson.M{"_id": uuid}
	result, err := repositories.FindOneUser(conditionSearch)
	if err != nil {
		return result, err
	}
	return result, nil
}

// UpdateUserByID func update user by uuid
func UpdateUserByID(uuid string, payload interfaces.UpdateUser) error {
	updatePayload := bson.M{"$set": bson.M{
		"username":   payload.Username,
		"updated_at": time.Now(),
	}}
	err := repositories.UpdateUser(bson.M{"_id": uuid}, updatePayload)
	if err != nil {
		return err
	}
	return nil
}

// RemoveUserByID func update user by uuid
func RemoveUserByID(uuid string) error {
	err := repositories.RemoveUser(bson.M{"_id": uuid})
	if err != nil {
		return err
	}
	return nil
}
