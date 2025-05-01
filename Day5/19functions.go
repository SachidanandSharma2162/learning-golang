package main

import (
	"fmt"
)
 
func checkPrime(a int) bool {
	var count int=0
	for i := 2; i < a; i++ {
		if a%i==0{
			count++
		}
	}
	if count==0 {
		return true
	} else {
		return false
	}
}
func main()  {
	fmt.Println("Enter a number:")
	var b int
	fmt.Scanf("%d",&b)
	var isPrime bool=checkPrime(b)
	if isPrime{
		fmt.Println("Prime")
	} else {
		fmt.Println("Not Prime")
	}
}