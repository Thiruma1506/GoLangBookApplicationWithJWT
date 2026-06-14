package config

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectToDB() *mongo.Database {
	var connectionString string = "mongodb://localhost:27017/"

	ctx,cancel := context.WithTimeout(context.Background() , time.Second*10)
	defer cancel()

	//connecting by creating client
	client, err := mongo.Connect(ctx,options.Client().ApplyURI(connectionString))
	if err != nil {
		log.Fatalf("The mongo connection has failed... %v", err.Error())
	}

	//send a ping check to clinet
	pingErr:=client.Ping(ctx,nil)
	if pingErr != nil{
		log.Fatalf("The ping check has failed... %v",pingErr.Error())
	}
	
	return client.Database("BookDb")
}