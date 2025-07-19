package main

import (
	"fmt"
	"sync"
)

func sayHello(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		fmt.Println("Hello")
	}
}
func sayWorld(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		fmt.Println("World")
	}
}
func main() {
	var wg sync.WaitGroup
	
	wg.Add(2)
	go sayHello(&wg)
	go sayWorld(&wg)
	wg.Wait()

}