package usecase

import (
	"Day23/config"
	"Day23/model"
	"encoding/json"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
)

type EmployeeService struct {
	MongoCollection *mongo.Collection
}

type Response struct {
	Data  interface{} `json:"data,omitempty"`
	Error string      `json:"error,omitempty"`
}

func (svc *EmployeeService) InsertEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	res := &Response{}
	defer json.NewEncoder(w).Encode(res)

	var emp model.Employee

	err := json.NewDecoder(r.Body).Decode(&emp)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Invalid body", err)
		res.Error = err.Error()
		return
	}

	// assign new emp id

	emp.EmployeeId = uuid.NewString()

	repo := config.EmployeeConfig{MongoCollection: svc.MongoCollection}

	insertId, err := repo.InsertEmployee(&emp)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("insert error", err)
		return
	}

	res.Data = emp.EmployeeId

	w.WriteHeader(http.StatusOK)

	log.Println("Employee created", insertId, emp)
}
func (svc *EmployeeService) GetEmployeeById(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	res := &Response{}
	defer json.NewEncoder(w).Encode(res)

	empId := mux.Vars(r)["id"]

	log.Println("employee id", empId)

	repo := config.EmployeeConfig{MongoCollection: svc.MongoCollection}

	emp, err := repo.GetEmployeeById(empId)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Error in getting details", err)
		res.Error = err.Error()
		return
	}

	res.Data = emp
	w.WriteHeader(http.StatusOK)

}
func (svc *EmployeeService) GetAllEmployees(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	res := &Response{}
	defer json.NewEncoder(w).Encode(res)

	repo := config.EmployeeConfig{MongoCollection: svc.MongoCollection}

	emp, err := repo.GetAllEmployees()

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Error in getting details", err)
		res.Error = err.Error()
		return
	}

	res.Data = emp
	w.WriteHeader(http.StatusOK)
}
func (svc *EmployeeService) UpdateById(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	res := &Response{}
	defer json.NewEncoder(w).Encode(res)

	// get emp id

	empId:=mux.Vars(r)["id"]
	log.Println("employee id", empId)

	if empId==""{
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Invalid employee id")
		res.Error = "Invalid employee id"
		return
	}

	var emp model.Employee

	err := json.NewDecoder(r.Body).Decode(&emp)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Invalid body", err)
		res.Error = err.Error()
		return
	}

	emp.EmployeeId=empId

	repo:=config.EmployeeConfig{MongoCollection:svc.MongoCollection}

	count,err:=repo.UpdateById(empId,&emp)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Error", err)
		res.Error = err.Error()
		return
	}
	res.Data=count
	w.WriteHeader(http.StatusOK)
}

func (svc *EmployeeService) DeleteById(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	res := &Response{}
	defer json.NewEncoder(w).Encode(res)

	// get emp id

	empId:=mux.Vars(r)["id"]
	log.Println("employee id", empId)
	if empId==""{
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Invalid employee id")
		res.Error = "Invalid employee id"
		return
	}
	repo:=config.EmployeeConfig{MongoCollection:svc.MongoCollection}

	count,err:=repo.DeleteById(empId)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Error", err)
		res.Error = err.Error()
		return
	}
	res.Data=count
	w.WriteHeader(http.StatusOK)
}
func (svc *EmployeeService) DeleteAll(w http.ResponseWriter, r *http.Request)  {
	w.Header().Add("Content-Type", "application/json")
	res := &Response{}
	defer json.NewEncoder(w).Encode(res)

	repo:=config.EmployeeConfig{MongoCollection:svc.MongoCollection}

	count,err:=repo.DeleteAll()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Error", err)
		res.Error = err.Error()
		return
	}
	res.Data=count
	w.WriteHeader(http.StatusOK)
}
