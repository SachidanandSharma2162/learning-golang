package main

import (
	"fmt"
	"sync"
)


func main()  {
	websiteList:=[]string{
		"https://google.com",
		"https://fb.com",
		"https://github.com",
		"https://go.dev",
	}
	
	var wg sync.WaitGroup
	var mu sync.Mutex
	for _, web:=range websiteList{
		wg.Add(1)
		go getStatusCode(web,&wg,&mu)
	}
	wg.Wait()
}

func getStatusCode(web string, wg *sync.WaitGroup, mu *sync.Mutex){
	// res,err:=http.Get(web)

	// if err!=nil{
	// 	fmt.Printf("there was an error!\n")
	// } else {
	// 	fmt.Printf("%d is the status code\n",res.StatusCode)
	// }
	fmt.Println("Locked")
	mu.Lock()
	fmt.Println(web)
	fmt.Println("Unlocked")
	mu.Unlock()
	wg.Done()
}