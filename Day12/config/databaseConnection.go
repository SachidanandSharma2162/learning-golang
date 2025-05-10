package config

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb+srv://sharmasachidanand1111:62km%40MAqP@cluster0.gxnxjpl.mongodb.net/?retryWrites=true&w=majority"
const dbName = "golang"
const collectionName = "watchlist"

var Collection *mongo.Collection

func init(){
	// client options
	clientOptions:= options.Client().ApplyURI(connectionString)


	// connect to mongoDB
	client,err:=mongo.Connect(context.TODO(),clientOptions)

	if err!=nil{
		log.Fatal(err) 
	}
    fmt.Println("Mongo Connection Done!!")

	Collection=client.Database(dbName).Collection(collectionName)

	// collection instance
	fmt.Println("collection instance is ready!!")
}