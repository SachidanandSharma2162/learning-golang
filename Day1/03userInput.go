package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	// Scanln is used to take user inputs

	var name string
	fmt.Println("Enter your name: \n")
	fmt.Scanln(&name)
	fmt.Println("Your name: \n", name)

	var number int
	fmt.Println("Enter number: \n")
	fmt.Scanln(&number)
	fmt.Println("The number is: \n", number)

	// to take input in format specified

	fmt.Printf("Your name is: %s and your age is: %d \n", name, number)

	// bufio ka bhi use kr ke ham user inout ke skte hai
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter Your Name: \n")
	userName, _ := reader.ReadString('\n')
	fmt.Println(userName)
}
