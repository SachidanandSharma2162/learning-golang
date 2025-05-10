package routes

import (

	"github.com/gorilla/mux"
)

func Router() *mux.Router{
	r := mux.NewRouter()
	r.HandleFunc("/allmovies",GetAllMoviesFormDatabase).Methods("GET")
	r.HandleFunc("/createmovie",CraeteMovieInDatabase).Methods("POST")
	r.HandleFunc("/markaswatched/{id}",MarkAsWatchedInDatabase).Methods("PUT")
	r.HandleFunc("/deletemovie/{id}",DeleteOneMovieFromDatabase).Methods("DELETE")
	r.HandleFunc("/deleteallmovies",DeleteAllMoviesFromDatabase).Methods("DELETE")
	return r
}