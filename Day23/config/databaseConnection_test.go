package config

import (
	"Day23/model"
	"context"
	"log"
	"testing"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func newMongoClient() *mongo.Client {
	mongoTestClient, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb+srv://sharmasachidanand1111:62kmMAqP@cluster0.8cakkxg.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0"))

	if err != nil {
		log.Fatal("Error while making connection", err)
	}

	log.Println("Connection Successfully Established")

	err = mongoTestClient.Ping(context.Background(), readpref.Primary())

	log.Println("Ping Success")

	return mongoTestClient
}

func TestMongoOperation(t *testing.T) {
	mongoTestClient := newMongoClient()

	defer mongoTestClient.Disconnect(context.Background())

	emp1 := uuid.New().String()
	coll := mongoTestClient.Database("crud-app").Collection("crud-app")

	empRepo := EmployeeConfig{MongoCollection: coll}

	t.Run("Insert Employee 1", func(t *testing.T) {
		emp := model.Employee{
			Name:       "Aaamn Yadav",
			Department: "Science",
			EmployeeId: emp1,
		}
		result, err := empRepo.InsertEmployee(&emp)

		if err != nil {
			t.Fatal("Insert 1 Failed", err)
		}
		t.Log("Insertion Successfull!", result)
	})

	t.Run("Get Employee 1",func(t *testing.T) {
		result,err:=empRepo.GetEmployeeById(emp1)

		if err!=nil{
			t.Fatal("failed to get employee 1",err)
		}

		t.Log("emp 1",result.Name)
	})

	t.Run("Get all employee",func(t *testing.T) {
		results,err:=empRepo.GetAllEmployees()

		if err!=nil{
			t.Fatal("get operation failed",err)
		}

		t.Log("Employees",results)
	})

	t.Run("Update Emp 1 name",func(t *testing.T) {
		emp:=model.Employee{
			Name: "Aman Yadav as AY",
			Department: "Science",
			EmployeeId: emp1,
		}

		result,err:=empRepo.UpdateById(emp1,&emp)

		if err!=nil{
			t.Fatal("Update operation falied",err)
		}

		t.Log("update count",result)
	})

	t.Run("Dlete emp and get all",func(t *testing.T) {
		result,err:=empRepo.GetAllEmployees()

		if err!=nil{
			t.Fatal("Delete operation failed",err)
		}

		t.Log("employees",result)
	})

	t.Run("Delete all employee",func(t *testing.T) {
		result,err:=empRepo.DeleteAll()

		if err!=nil{
			log.Fatal("Delete operation failed",err)
		}
		t.Log("results",result)
	})
}
