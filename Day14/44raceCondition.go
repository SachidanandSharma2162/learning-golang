package main

import (
	"fmt"
	"sync"
)

func main()  {
	fmt.Println("Race Condition!")

	var score=[]int{0}
	var wg sync.WaitGroup
	var mu sync.Mutex
	// effi functions
	wg.Add(3)
	go func(wg *sync.WaitGroup,mu *sync.Mutex){
		defer wg.Done()
		mu.Lock()
		fmt.Println("One R")
		score=append(score, 1)
		mu.Unlock()
	}(&wg,&mu)
	// wg.Add(1)
	go func(wg *sync.WaitGroup,mu *sync.Mutex){
		defer wg.Done()
		mu.Lock()
		fmt.Println("Two R")
		score=append(score, 2)
		mu.Unlock()
	}(&wg,&mu)
	// wg.Add(1)
	go func(wg *sync.WaitGroup,mu *sync.Mutex){
		defer wg.Done()
		mu.Lock()
		fmt.Println("Three R")
		score=append(score, 3)
		mu.Unlock()
	}(&wg,&mu)

	wg.Wait()

	fmt.Println(score)
}