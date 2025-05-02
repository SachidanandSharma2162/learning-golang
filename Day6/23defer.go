package main

import (
	"fmt"
)

func main()  {
	fmt.Println("defer method")

	defer fmt.Println("One")
	defer fmt.Println("This is defered!!")
	defer fmt.Println("Three")

	// works like LIFO stack
	// usually used for file handling and database management
	for i := 0; i < 5; i++ {
		defer fmt.Print(i)
	}
}