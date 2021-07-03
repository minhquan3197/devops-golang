package services

import (
	"errors"
	"project-golang/internal/interfaces"
	"project-golang/internal/models"
	"project-golang/internal/repositories"
	"project-golang/pkg/bcrypt"
	"project-golang/pkg/constants"
	"project-golang/pkg/mongodb"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func userRepository() *repositories.UserRepositoryMongo {
	userRepository := repositories.NewUserRepositoryMongo(mongodb.GetDB(), models.UsersCollection)
	return userRepository
}

// Register func create new user
func Register(payload interfaces.CreateUser) error {
	var user *models.UserSchema
	user = &models.UserSchema{
		Username: payload.Username,
		Password: bcrypt.HashPassword(payload.Password),
	}
	err := userRepository().Save(user)
	if err != nil {
		return errors.New(constants.UserCreatedFailed)
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
	results, err := userRepository().FindAll(conditionSearch, findOptions)
	if err != nil {
		return resp, errors.New(constants.UserQueryIncorrect)
	}
	countUser, err := userRepository().Count(conditionSearch)
	if err != nil {
		return resp, errors.New(constants.UserQueryIncorrect)
	}
	resp = interfaces.PaginateUser{
		Data:    results,
		PerPage: limit,
		Total:   countUser,
	}
	return resp, nil
}

// FindUserByUsername func find user by username
func FindUserByUsername(username string) (models.UserSchema, error) {
	result, err := userRepository().FindByUsername(username)
	if err != nil {
		return result, errors.New(constants.UserNotFound)
	}
	return result, nil
}

// FindUserByUsername func find user by username
func FindUserByUsernameLogin(username string) (models.UserSchema, error) {
	result, err := userRepository().FindByUsernameLogin(username)
	if err != nil {
		return result, errors.New(constants.UserNotFound)
	}
	return result, nil
}

// FindUserByID func find user by uuid
func FindUserByID(uuid string) (models.UserSchema, error) {
	result, err := userRepository().FindByID(uuid)
	if err != nil {
		return result, errors.New(constants.UserNotFound)
	}
	return result, nil
}

// UpdateUserByID func update user by uuid
func UpdateUserByID(uuid string, payload interfaces.UpdateUser) error {
	var user *models.UserSchema
	user = &models.UserSchema{
		Username: payload.Username,
	}
	err := userRepository().UpdateInfo(uuid, user)
	if err != nil {
		return errors.New(constants.UserNotFound)
	}
	return nil
}

// RemoveUserByID func update user by uuid
func RemoveUserByID(uuid string) error {
	err := userRepository().DeleteByID(uuid)
	if err != nil {
		return errors.New(constants.UserNotFound)
	}
	return nil
}
