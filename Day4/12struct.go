package main

import (
	"fmt"
)
// name of the structure must be in capital latters 
type Users struct{
	name string
	age int
}
func main()  {
	userOne:=Users{
		name:"Aman",
		age: 23,
	}
	fmt.Println(userOne)
	
	// accessing values
	fmt.Println(userOne.name,userOne.age)
	
	// updating values
	userOne.name="Adil"
	fmt.Println(userOne)

	// Anonymous syntax
	emp := struct{
		Id   int
		Name string
	}{101,"Rahul"}
	fmt.Println(emp)
    fmt.Printf("%+v",userOne)

}