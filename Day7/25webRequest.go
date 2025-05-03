package main

import (
	"fmt"
	"net/http"
)
func myName(w http.ResponseWriter,r *http.Request){
	fmt.Fprintln(w,"Hello this is me!!")
}
func main()  {
	http.HandleFunc("/",myName)
	fmt.Println("Server is running on port 3000\n")
	http.ListenAndServe(":3000",nil)
}