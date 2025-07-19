package main

import "fmt"
func printslice1(items []interface{}){
	for _,v:= range items{
		fmt.Println(v)
	}
}
func printslice2[T any](items []T){
	for _,v:= range items{
		fmt.Println(v)
	}
}

// type constraints
func printslice3[T int | string | bool](items []T){
	for _,v:= range items{
		fmt.Println(v)
	}
}

type stack[T any] struct{
	elements []T
}
func main(){
	printslice3([]int{1,2,3,4,5,6})
	printslice3([]string{"Aman","Rohit","Rajeev"})
	printslice3([]bool{true,false})


	stack:= stack[any]{
		elements: []any{1,2,"apple",4,5,6},
	}

	fmt.Println(stack)
}