package main

import "fmt"

func main(){
	//Inline Struct
	person:= struct {
		name string
		age  int
	}{"Amit",44}

	fmt.Println(person)
	fmt.Println(person.name)
	fmt.Println(person.age)
}