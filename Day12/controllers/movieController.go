// MONGO HELPERS
package controllers

import (
	"context"
	"day12/config"
	"day12/models"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func InsertOneMovie(movie models.Netflix) {
	result, err := config.Collection.InsertOne(context.Background(), movie)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted Successfulll, with id: ", result.InsertedID)
}

func UpdateOneMovie(movieId string) {
	id, _ := primitive.ObjectIDFromHex(movieId)
	filter := bson.M{
		"_id": id,
	}
	update := bson.M{
		"$set": bson.M{
			"watched": true,
		},
	}
	result,err:=config.Collection.UpdateOne(context.Background(),filter,update)

	if err!=nil{
		log.Fatal(err)
	}

	fmt.Println("Updated Successfully, count: ",result.ModifiedCount)
}


func DeleteOneMovie(movieId string) {
	id, _ := primitive.ObjectIDFromHex(movieId)
	filter:=bson.M{
		"_id":id,
	}
	deleteCount,err:=config.Collection.DeleteOne(context.Background(),filter)

	if err!=nil{
		log.Fatal(err)
	}
	fmt.Println("Successfully Deleted, Delete Count: ",deleteCount)
}

func DeleteAllMovies() int64 {
	filter:=bson.D{{}}
	deleteResult,err:=config.Collection.DeleteMany(context.Background(),filter,nil)

	if err!=nil{
		log.Fatal(err)
	}

	fmt.Println("Deleted Successfully, Delete Count: ",deleteResult.DeletedCount)
	return deleteResult.DeletedCount
}

func GetAllMovies() []primitive.M {
	cursor,err:=config.Collection.Find(context.Background(),bson.D{{}})
	if err!=nil{
		log.Fatal(err)
	}
	var movies []primitive.M
	for cursor.Next(context.Background()){
		var movie bson.M
		err:=cursor.Decode(&movie)
		if err!=nil{
			log.Fatal(err)
		}
		movies = append(movies, movie)
	}

	defer cursor.Close(context.Background())
	return movies
}
