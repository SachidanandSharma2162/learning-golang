package main

import (
    "fmt"
    "time"
)

func sayHello() {
    for i := 0; i < 5; i++ {
        fmt.Println("Hello")
        time.Sleep(100 * time.Millisecond)
    }
}

func sayWorld() {
    for i := 0; i < 5; i++ {
        fmt.Println("World")
        time.Sleep(100 * time.Millisecond)
    }
}

func main() {
    go sayHello()
    sayWorld()    

    time.Sleep(1 * time.Second) // wait for goroutines to finish
}
