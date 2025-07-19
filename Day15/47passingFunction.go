package main

import "fmt"

func calculator(a int,b int,op func(int, int) int)int{
	return op(a,b)
}
func add(a ,b int) int{
	return a+b
}
func multiply(a ,b int) int{
	return a*b
}
func divide(a ,b int) int{
	return a/b
}
func subtract(a ,b int) int{
	return a-b
}
func main() {
	fmt.Println("Passing Function as argument!")

	var ch int
	fmt.Println("Enter a choice: 1,2,3 and 4")
	fmt.Scanf("%d",&ch)

	switch ch{
	case 1:
		fmt.Println("Addition:", calculator(10,5,add))
	case 2:
		fmt.Println("Multiplication:", calculator(10,5,multiply))
	case 3:
		fmt.Println("Dvivision:", calculator(10,5,divide))
	case 4:
		fmt.Println("Subtraction:", calculator(10,5,subtract))
	default:
		fmt.Println("Enter a valid choice!")
	}
}