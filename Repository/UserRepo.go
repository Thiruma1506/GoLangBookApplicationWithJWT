package repository

import (
	model "bookapplication/Model"
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository interface {
	FindUserByEmail(email string) (*model.User,error)
	CreateNewUser(user *model.User) error
}

type UserRepositoryImpl struct {
	collection *mongo.Collection
}

func NewUserRepository(db *mongo.Database) UserRepository {
	return &UserRepositoryImpl{
		collection: db.Collection("UserCollection"),
	}
}

func createUserCtx() (context.Context, context.CancelFunc) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	return ctx, cancel
}

// CreateNewUser implements [Repository].
func (Urepo *UserRepositoryImpl) CreateNewUser(user *model.User) error {
	ctx, cancel := createUserCtx()
	defer cancel()

	resp, err := Urepo.collection.InsertOne(ctx, user)
	if err != nil {
		log.Printf("The User data was not persisted Successfully... %v", err)
		return err
	}
	fmt.Printf("The User with ID - %v has been successfully persisted", resp.InsertedID)
	return nil
}

// FindUserByEmail implements [Repository].
func (Urepo *UserRepositoryImpl) FindUserByEmail(email string) (*model.User, error) {
	ctx, cancel := createUserCtx()
	defer cancel()

	filter := bson.M{"emailId": email}
	var respUser model.User
	if err := Urepo.collection.FindOne(ctx, filter).Decode(&respUser); err != nil {
		log.Printf("The User with the given email is not present... %v", err)
		return nil,err
	}
	return &respUser,nil
}
