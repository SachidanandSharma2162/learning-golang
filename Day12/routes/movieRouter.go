package routes

import (
	"day12/controllers"
	"day12/models"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func GetAllMoviesFormDatabase(w http.ResponseWriter,r *http.Request) {
	w.Header().Set("Content-type","application/x-www-form-urlencode")
	allMovies:=controllers.GetAllMovies()

	json.NewEncoder(w).Encode(allMovies)
}

func CraeteMovieInDatabase(w http.ResponseWriter,r *http.Request)  {
	w.Header().Set("Content-type","application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods","POST")
	var movie models.Netflix
	_=json.NewDecoder(r.Body).Decode(&movie)
	controllers.InsertOneMovie(movie)
	json.NewEncoder(w).Encode(movie)

}

func MarkAsWatchedInDatabase(w http.ResponseWriter,r *http.Request)  {
	w.Header().Set("Content-type","application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods","PUT")


	params:=mux.Vars(r)
	controllers.UpdateOneMovie(params["id"])
	json.NewEncoder(w).Encode("Updated Successfully!")
}

func DeleteOneMovieFromDatabase(w http.ResponseWriter,r *http.Request)  {
	w.Header().Set("Content-type","application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods","DELETE")

	params:=mux.Vars(r)

	controllers.DeleteOneMovie(params["id"])

	json.NewEncoder(w).Encode("Deleted Successfully!")
}
func DeleteAllMoviesFromDatabase(w http.ResponseWriter,r *http.Request)  {
	w.Header().Set("Content-type","application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods","DELETE")

	count:=controllers.DeleteAllMovies()
	json.NewEncoder(w).Encode(count)
}