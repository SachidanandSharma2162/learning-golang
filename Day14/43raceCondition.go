package main

import (
	"fmt"
	"sync"
)

var count = 0 
func increment(id string,wg *sync.WaitGroup,mu *sync.Mutex)  {
	for i:=1;i<=20;i++{
		mu.Lock()
		fmt.Println(id)
        count++
        mu.Unlock() 
	}
	wg.Done()
}
func main(){
	fmt.Println("Race Condition!")
	var wg sync.WaitGroup
	var mu sync.Mutex
	wg.Add(2)
	go increment("first",&wg,&mu)
	go increment("second",&wg,&mu)

	wg.Wait()
	fmt.Println(count)
}