package main

import (
	"Day23/usecase"
	"context"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var mongoClient *mongo.Client
func init(){
	// load env variable
	err:=godotenv.Load()

	if err!=nil{
		log.Fatal("Error loading env file")
	}
	log.Println("env file loaded")

	// Connection

	mongoClient,err=mongo.Connect(context.Background(),options.Client().ApplyURI(os.Getenv("MONGO_URI")))

	if err!=nil{
		log.Fatal("Error connecting mongoDB",err)
	}
	err=mongoClient.Ping(context.Background(),readpref.Primary())
	if err!=nil{
		log.Fatal("ping failed",err)
	}

	log.Println("Connection established")
}
func main() {
	// close the mongo connection
	defer mongoClient.Disconnect(context.Background())

	coll:= mongoClient.Database(os.Getenv("DB_NAME")).Collection(os.Getenv("COLLECTION_NAME"))
	// Employee Service

	empService:=usecase.EmployeeService{MongoCollection: coll}
	r:=mux.NewRouter()

	r.HandleFunc("/health",healthHandler).Methods(http.MethodGet)
	r.HandleFunc("/insertemployee",empService.InsertEmployee).Methods(http.MethodPost)
	r.HandleFunc("/getemployeebyid/{id}",empService.GetEmployeeById).Methods(http.MethodGet)
	r.HandleFunc("/getallemployee",empService.GetAllEmployees).Methods(http.MethodGet)
	r.HandleFunc("/updateemployeebyid/{id}",empService.UpdateById).Methods(http.MethodPut)
	r.HandleFunc("/deleteemployeebyid/{id}",empService.DeleteById).Methods(http.MethodDelete)
	r.HandleFunc("/deleteallemployee",empService.DeleteAll).Methods(http.MethodDelete)


	log.Println("Server is running at port 4444")
	http.ListenAndServe(":4444",r)

}

func healthHandler(w http.ResponseWriter, r *http.Request){
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("running..."))
}