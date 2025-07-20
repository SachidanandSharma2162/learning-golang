package main

import (
	"Day22/mathsutils"
	"fmt"
)

func main(){

	var a,b int
	fmt.Println("Enter the value of a and b:")
	fmt.Scan(&a,&b)
	var ch int
	fmt.Println("Enter the choice\n 1. Add \n 2. Subtract \n 3. Multiply \n 4. Divide \n 5. Exit")
	fmt.Scan(&ch)

	switch ch {
	case 1:
		sum:=mathsutils.Add(a,b)
		fmt.Println(sum)
	case 2:
		sub:=mathsutils.Subtract(a,b)
		fmt.Println(sub)
	case 3:
		mul:=mathsutils.Multiply(a,b)
		fmt.Println(mul)
	case 4:
		div:=mathsutils.Divide(a,b)
		fmt.Println(div)
	case 5:
		fmt.Println("Thanks for using Calculator")
	default:
		fmt.Println("Enter valid choice!")
	}
}