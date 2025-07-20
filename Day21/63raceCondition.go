package main

import (
	"fmt"
	"sync"
	
)
type Movie struct{
	views int
	mu sync.Mutex
}
func (m *Movie) incrementViews(wg *sync.WaitGroup) {
	defer func(){
		m.mu.Unlock()
		wg.Done()
	}()
	m.mu.Lock()
	m.views+=1
}


func main() {

	var wg sync.WaitGroup
	movie:=Movie{views: 0}
	for i:=0;i<100;i++{
		wg.Add(1)
		go movie.incrementViews(&wg)
	}
	
	wg.Wait()
	fmt.Println(movie.views)
}