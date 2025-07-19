package main

import "fmt"
  
func sum(values ...int) int {
	tot:=0
	for _, v := range values {
		tot = tot + v
	}
	return tot
}
func main() {
	sumOfNumbers := sum(1,2,3,4,5,6,7,8,9,10)
	fmt.Println("Sum of numbers:",sumOfNumbers)
}
