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

type Repo interface {
	CreateNewBook(book *model.Book) error
	UpdateBook(book *model.Book, id string) error
	GetBookById(id string) (*model.Book, error)
	DeleteBookById(id string) error
	CreateListOfBooks(books []*model.Book) error
}

// Injecting the mongo collection as a dependency
type MongoRepository struct {
	collection *mongo.Collection
}

// NewMongoRepository constructor to assign the collection name to the database object passed from main()
func NewMongoRepository(db *mongo.Database) Repo {
	return &MongoRepository{
		collection: db.Collection("BookCollection"),
	}
}

func createCtx() (context.Context, context.CancelFunc) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	return ctx, cancel
}

// repository functions
func (repo *MongoRepository) CreateNewBook(book *model.Book) error {
	ctx, cancel := createCtx()
	defer cancel()

	resp, err := repo.collection.InsertOne(ctx, book)
	if err != nil {
		log.Fatalf("The data was not persisted in the collection... %v", err)
		return err
	}
	log.Printf("The data was persisted in the collection with id: %v", resp.InsertedID)
	return nil
}

func (repo *MongoRepository) UpdateBook(book *model.Book, id string) error {

	ctx, cancel := createCtx()
	defer cancel()

	filter := bson.M{"_id": id}
	update := bson.M{"$set": book}
	resp, err := repo.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		log.Fatalf("Update operation failed %v", err)
		return err
	}
	fmt.Print(resp.ModifiedCount)
	return nil
}

func (repo *MongoRepository) GetBookById(id string) (*model.Book, error) {

	ctx, cancel := createCtx()
	defer cancel()

	var Book model.Book
	filter := bson.M{"_id": id}
	err := repo.collection.FindOne(ctx, filter).Decode(&Book)
	if err != nil {
		log.Fatalf("Id not found in DB... %v", err)
		return nil, err
	}
	return &Book, nil
}

func (repo *MongoRepository) DeleteBookById(id string) error {
	ctx, cancel := createCtx()
	defer cancel()

	filter := bson.M{"_id": id}
	resp, err := repo.collection.DeleteOne(ctx, filter)
	if err != nil {
		log.Fatalf("The Book with Id %v HAS NOT BEEN DELETED %v...", id, err)
		return err
	}
	fmt.Print(resp)
	return nil
}

func (repo *MongoRepository) CreateListOfBooks(books []*model.Book) error {
	ctx, cancel := createCtx()
	defer cancel()

	BookList := make([]any, len(books))
	for i:=range books{
		BookList[i]=books[i]
	}

	result, err := repo.collection.InsertMany(ctx, BookList)
	if err != nil {
		log.Printf("The List of Books are not persisted successfully %v", err)
		return err
	}
	fmt.Print(result.InsertedIDs...)
	return nil
}
