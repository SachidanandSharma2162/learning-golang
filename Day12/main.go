package main

import (
	"day12/routes"
	"fmt"
	"log"
	"net/http"
)

func main()  {
	fmt.Println("mongoDB for go!")
	fmt.Println("Server is getting started...")
	r:=routes.Router()
	log.Fatal(http.ListenAndServe(":3000",r))
	fmt.Println("Listening to port 3000...")

}