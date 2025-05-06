package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main(){
	fmt.Println("Let's learn mod in golang.")
	r:=mux.NewRouter()
	r.HandleFunc("/",serverCode).Methods("GET")



	log.Fatal(http.ListenAndServe(":3000",r))
}

func serverCode(w http.ResponseWriter,r *http.Request){
	w.Write([]byte("<h1>Learning golang on YT ~ Hitesh Chaudhary</h1>"))
}