package repositories

import (
	"context"
	"project-golang/internal/models"
	"time"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type (
	User = models.UserSchema
)

type UserRepository interface {
	Save(*User) error
	Update(string, *User) error
	Delete(string) error
	FindByID(string) (User, error)
	FindAll(bson.M, *options.FindOptions) ([]User, error)
	Count(bson.M) (int64, error)
}

type UserRepositoryMongo struct {
	db         *mongo.Database
	collection string
}

func NewUserRepositoryMongo(db *mongo.Database, collection string) *UserRepositoryMongo {
	return &UserRepositoryMongo{
		db:         db,
		collection: collection,
	}
}

func (r *UserRepositoryMongo) Save(user *User) error {
	user.UpdatedAt = time.Now().Unix()
	user.CreatedAt = time.Now().Unix()
	user.ID = uuid.New().String()
	_, err := r.db.Collection(r.collection).InsertOne(context.TODO(), user)
	return err
}

func (r *UserRepositoryMongo) Update(id string, user *User) error {
	user.UpdatedAt = time.Now().Unix()
	_, err := r.db.Collection(r.collection).UpdateOne(context.TODO(), bson.M{"_id": id}, user)
	return err
}

func (r *UserRepositoryMongo) Delete(id string) error {
	var user User
	err := r.db.Collection(r.collection).FindOneAndDelete(context.TODO(), bson.M{"_id": id}).Decode(&user)
	return err
}

func (r *UserRepositoryMongo) FindByID(id string) (User, error) {
	var user User
	projection := bson.M{"password": 0}
	err := r.db.Collection(r.collection).FindOne(
		context.TODO(),
		bson.M{"_id": id},
		options.FindOne().SetProjection(projection)).Decode(&user)
	return user, err
}

func (r *UserRepositoryMongo) FindByUsername(username string) (User, error) {
	var user User
	projection := bson.M{"password": 0}
	err := r.db.Collection(r.collection).FindOne(
		context.TODO(),
		bson.M{"username": username},
		options.FindOne().SetProjection(projection)).Decode(&user)
	return user, err
}

func (r *UserRepositoryMongo) FindAll(condition bson.M, options *options.FindOptions) ([]User, error) {
	var users []User
	projection := bson.M{"password": 0}
	options.SetProjection(projection)
	pointer, err := r.db.Collection(r.collection).Find(context.TODO(), condition, options)
	err = pointer.All(context.TODO(), &users)
	return users, err
}

func (r *UserRepositoryMongo) Count(condition bson.M) (int64, error) {
	result, err := r.db.Collection(r.collection).CountDocuments(context.TODO(), condition)
	return result, err
}
