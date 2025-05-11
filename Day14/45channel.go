package main

import (
	"fmt"
	"sync"

)


func main()  {
	fmt.Println("Channel in golang")
	var wg sync.WaitGroup
	// declara a channel
	ch:=make(chan string,2)
	wg.Add(1)
	go func(wg *sync.WaitGroup){
		defer wg.Done()
		ch <- "Hello"
		ch <- "World"
	}(&wg)
	wg.Add(1)
	go func(wg *sync.WaitGroup){
		defer wg.Done()
		msg:= <-ch
		fmt.Println(msg)
		msg1:= <-ch
		fmt.Println(msg1)
	}(&wg)
	wg.Wait()
	close(ch)

	_,isChannelOpen := <- ch

	if !isChannelOpen {
		fmt.Println("Channel Closed!")
	}

}