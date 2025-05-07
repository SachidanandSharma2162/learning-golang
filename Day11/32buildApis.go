package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

// model for courses
type Course struct{
	CourseId string `json:"courseid"`
	CourseName string `json:"coursename"`
	CoursePrice int `json:"courseprice"`
	Author *Author `json:"author"`
	
}

type Author struct{
	Fullname string `json:"fulllname"`
	Website string `json:"website"`
}

// fake DB
var courses []Course

// middlewares
func (course *Course) isEmpty() bool{
	return course.CourseName==""
}
func main() {
	fmt.Println("learning APIs")
	r:=mux.NewRouter()

	// seeding
	courses = append(courses, Course{
		CourseId:"1",
	    CourseName:"Javascript",
	    CoursePrice:3000,
	    Author: &Author{
			Fullname:"Sachidanand Sharma",
	        Website:"github/author_ss",},
	})
	courses = append(courses, Course{
		CourseId:"2",
	    CourseName:"Mern",
	    CoursePrice:2000,
	    Author: &Author{
			Fullname:"Perman Peter",
	        Website:"github/author_pp",},
	})
	r.HandleFunc("/",serverHomeRoute).Methods("GET")
	r.HandleFunc("/courses",getAllCourses).Methods("GET")
	r.HandleFunc("/course/{id}",getOneCourse).Methods("GET")
	r.HandleFunc("/course",createOneCourse).Methods("POST")
	r.HandleFunc("/course/{id}",updateOneCourse).Methods("PUT")
	r.HandleFunc("/course/{id}",deleteOneCourse).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":3000",r))
}

// controllers
func serverHomeRoute(w http.ResponseWriter, r *http.Request)  {
	w.Write([]byte("<h1>Welcome to the home route!</h1>"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request)  {
	fmt.Println("get all courses")
	w.Header().Set("Content-type","application/json")
	json.NewEncoder(w).Encode(courses)
}

func getOneCourse(w http.ResponseWriter, r *http.Request){
	fmt.Println("Get one course")
	w.Header().Set("Content-type","application/json")

	// grab id
	params := mux.Vars(r)
	// loop through course and find matching course

	for _, course := range courses {
		if course.CourseId==params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("No course found with given id!")
	return
}

func createOneCourse(w http.ResponseWriter, r *http.Request)  {
	fmt.Println("Create one course")
	w.Header().Set("Content-type","application/json")

	// what if entire body is empty

	if r.Body==nil{
		json.NewEncoder(w).Encode("Please send some data")
		return
	}
	// what about-{}
	var course Course
	json.NewDecoder(r.Body).Decode(&course)
 
	if course.isEmpty(){
		json.NewEncoder(w).Encode("No data in json")
		return
	}
	
	// to generate a unique
	rand.Seed(time.Now().UnixNano())
	course.CourseId=strconv.Itoa(rand.Intn(100))
	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)
	return
}

func updateOneCourse(w http.ResponseWriter, r *http.Request)  {
	fmt.Println("Update one course")
	w.Header().Set("Content-type","application/json")

	params:=mux.Vars(r)

	// loop -> id -> remove -> add with id
	for index, course:= range courses{
		if course.CourseId==params["id"]{
			courses = append(courses[:index],courses[index+1:]...)
			_=json.NewDecoder(r.Body).Decode(&course)
			course.CourseId=params["id"]
			courses = append(courses, course)
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	// send a res when id id not found
	json.NewEncoder(w).Encode("course with this id not found")
	return
}

func deleteOneCourse(w http.ResponseWriter, r *http.Request)  {
	fmt.Println("Delete one course")
	w.Header().Set("Content-type","application/json")

	params:=mux.Vars(r)

	for idx,course:=range courses{
	 	if course.CourseId==params["id"]{
			courses=append(courses[:idx],courses[idx+1:]...)
			json.NewEncoder(w).Encode("Successfully removed the course!")
			break
		}
	}
}