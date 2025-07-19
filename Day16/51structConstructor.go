package main

import "fmt"

type Person struct {
	name string
	age  int
	city string
}

// constructor function for Person struct
func NewPerson(name string, age int, city string) *Person {
	newPerson := Person{
		name: name,
		age:  age,
		city: city,
	}
	// you can return the pointer of the struct
	return &newPerson
}

func main() {
	person1 := NewPerson("Manoj", 25, "Delhi")

	fmt.Println(person1)
}
