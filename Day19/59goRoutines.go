package main

import (
	"fmt"
	"time"
)

func sayHelloForGoRoutine() {
	for i := 0; i < 5; i++ {
		fmt.Println("Hello from Routine",i)
		time.Sleep(500*time.Millisecond)
	}
}
func sayHelloFromMain(){
	for i := 0; i < 5; i++ {
		fmt.Println("Hello from Main",i)
		time.Sleep(500*time.Millisecond)
	}
}
func main() {
	go sayHelloForGoRoutine() //goroutine
	// sayHelloFromMain() // main routine
	time.Sleep(2*time.Second)
}