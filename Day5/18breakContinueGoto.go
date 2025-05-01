package main

import "fmt"

func main()  {
	array:=[]int{1,2,3,4,5}
	fmt.Println(array)

	// break statement
	for i := 0; i < len(array); i++ {
		if(i>=3){
			break
		} else {
			fmt.Println(array[i])
		}
	}

	// continue statement
	for i := 0; i < len(array); i++ {
		if(i==3){
			continue
		} else {
			fmt.Println(array[i])
		}
	}

	// range for loop
	for index, value := range array {
		if index==3{
			break
		} else {
			fmt.Println(value)
		}
	}

	// goto statement:
	var num int=1

	for num<10{
		if(num==5){
			goto lco
		} else {
			fmt.Println(num)
			num++
		}
	}
	lco: 
	    fmt.Println("End")
}