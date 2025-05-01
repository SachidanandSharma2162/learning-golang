package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main()  {
	fmt.Println("We will be learning conditionals now!!")

	fmt.Println("Enter a number:\n")
	reader:=bufio.NewReader(os.Stdin)
	a,_:=reader.ReadString('\n')
	b,_:=strconv.ParseInt(strings.TrimSpace(a),10,64)

	if b%2==0 {
		fmt.Printf("%d is even number \n",b)
	} else {
		fmt.Printf("%d is odd number \n",b)
	}

	// if with short variable decleration

	if num:=23; num%3==0{
		fmt.Println("Divisible by 3 \n")
	} else {
		fmt.Println("Not divisible by 3 \n")

	}
}

