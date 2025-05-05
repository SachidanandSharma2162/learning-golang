package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

func main() { 
	fmt.Println("lets learn how to make Get request using golang!!")
	go startServer()
	time.Sleep(1 * time.Second)
	// PerformGetRequest()
	PerformPostJsonRequest()
}

func startServer() {
	http.HandleFunc("/get", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello from Go server!")
	})
	fmt.Println("🟢 Server running on http://127.0.0.1:3000")
	err := http.ListenAndServe("127.0.0.1:3000", nil)
	if err != nil {
		panic(err)
	}
}
func PerformGetRequest() {
	const url = "http://127.0.0.1:3000/get" 

	res, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	fmt.Println(res.StatusCode)
	fmt.Println(res.ContentLength)

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	fmt.Println("✅ Response Body:")
	fmt.Println(string(body))
}

func PerformPostJsonRequest(){
	const myurl="http://localhost:3000/signup"

	resData:=strings.NewReader(`{
	"username":"userOne",
	"password":"123",
	"course":"python"
	}`)
	res,err:=http.Post(myurl,"application/json",resData)
	if err!=nil{
		panic(err)
	}
	defer res.Body.Close()
	
	data,_:=io.ReadAll(res.Body)

	fmt.Println(string(data))

}