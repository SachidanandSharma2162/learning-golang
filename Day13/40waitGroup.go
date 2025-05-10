package main

import (
	"fmt"
	"sync"
	"time"
)

func sayHello(wg *sync.WaitGroup) {
	defer wg.Done()
    for i := 0; i < 5; i++ {
        fmt.Println("Hello")
        time.Sleep(100 * time.Millisecond)
    }
}

func sayWorld(wg *sync.WaitGroup) {
	defer wg.Done()
    for i := 0; i < 5; i++ {
        fmt.Println("World")
        time.Sleep(100 * time.Millisecond)
    }
}

func main()  {

	fmt.Println("go routine and waitGroup")
	var wg sync.WaitGroup

	wg.Add(2)
	go sayHello(&wg)
	go sayWorld(&wg)
	
	wg.Wait()
	fmt.Println("Rountines done!");
}