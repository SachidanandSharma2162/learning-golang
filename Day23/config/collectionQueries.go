package config

import (
	"Day23/model"
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type EmployeeConfig struct {
	MongoCollection *mongo.Collection
}

func (r *EmployeeConfig) InsertEmployee(emp *model.Employee)(interface{},error){
	result,err:=r.MongoCollection.InsertOne(context.Background(),emp)

	if err!=nil{
		return nil,err
	}
	return result.InsertedID,nil
}

func (r *EmployeeConfig) GetEmployeeById(empId string)(*model.Employee,error){
	var emp model.Employee

	err:=r.MongoCollection.FindOne(context.Background(),bson.D{{Key:"employee_id",Value:empId}}).Decode((&emp))

	if err!=nil{
		return nil,err
	}
	return &emp,nil
}

func (r *EmployeeConfig) GetAllEmployees()([]model.Employee,error){
	results,err:= r.MongoCollection.Find(context.Background(),bson.D{})

	if err!=nil{
		return nil,err
	}
	var employee []model.Employee

	err=results.All(context.Background(),&employee)
	if err!=nil{
		return nil,fmt.Errorf("resultdecode error: %s",err.Error())
	}
	return employee,nil
}

func (r *EmployeeConfig) UpdateById(empId string,updateEmp *model.Employee)(int64,error){
	result,err:=r.MongoCollection.UpdateOne(context.Background(),bson.D{{Key: "employee_id", Value: empId}},bson.D{{Key: "$set",Value:updateEmp }})

	if err!=nil{
		return 0,err
	}

	return result.ModifiedCount,nil
}

func (r *EmployeeConfig) DeleteById(empId string)(int64,error){
	result,err:=r.MongoCollection.DeleteOne(context.Background(),bson.D{{Key: "employee_id", Value: empId}})

	if err!=nil{
		return 0,err
	}
	return result.DeletedCount,nil
}

func (r *EmployeeConfig) DeleteAll()(int64,error){
	result,err:=r.MongoCollection.DeleteMany(context.Background(),bson.D{})

	if err!=nil{
		return 0,err
	}
	return result.DeletedCount,nil
}