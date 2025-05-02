package main

import "fmt"

func printFibonaaci(num int,first int,second int){
	for i := 1; i <= num; i++ {
		var third int = first+second
		fmt.Println(third)
		first=second
		second=third
	}
}
func main()  {
	var num int
	fmt.Println("Enter n: \n")
	fmt.Scanf("%d",&num)
	var first int = 0
	var second int = 1
	fmt.Println(first)
	fmt.Println(second)
	printFibonaaci(num,first,second)
}