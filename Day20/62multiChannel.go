package main

import "fmt"

func main() {

	chan1 := make(chan int)
	chan2 := make(chan string)

	go func() {
		chan1 <- 10
	}()
	go func() {
		chan2 <- "pong"
	}()
	defer func(){
		close(chan1)
		fmt.Println("chan1 closed")
		close(chan2)
		fmt.Println("chan2 closed")
	}()
	for i := 1; i <= 2; i++ {
		select {
		case chan1Val := <-chan1:
			fmt.Println("Received from chan1:", chan1Val)
		case chan2Val := <-chan2:
			fmt.Println("Received from chan2:", chan2Val)
		}
	}
}