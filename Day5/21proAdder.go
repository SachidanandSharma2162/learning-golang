package main

import "fmt"

// here we get a slice and we loop through to get the total, values is a slice

// we can also return various values like this
 
func proAdder(values...int)(int,string)  {
	total:=0

	for _, val := range values {
		total=total+val
	}

	return total,"the total we have calculated"
}
func main()  {
	// a:=5
	// b:=10
	c,_:=proAdder(1,2,3,4,5,6,7,8,9,10)
	fmt.Printf("the sumṇ is: %d",c)
}