package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"math/rand"
	"net/http"
	"strconv"
)

type Movie struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

var movies []Movie

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	err := json.NewEncoder(w).Encode(movies)

	if err != nil {
		return
	}
}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	parameters := mux.Vars(r)

	for index, item := range movies {
		if item.ID == parameters["id"] {
			movies = append(movies[:index], movies[index+1:]...)

			break
		}
	}

	err := json.NewEncoder(w).Encode(movies)

	if err != nil {
		return
	}
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	for _, item := range movies {
		if item.ID == params["id"] {
			err := json.NewEncoder(w).Encode(item)

			if err != nil {
				return
			}
		}
	}
}

func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var movie Movie

	_ = json.NewDecoder(r.Body).Decode(&movie)

	movie.ID = strconv.Itoa(rand.Intn(100000000))
	movies = append(movies, movie)

	err := json.NewEncoder(w).Encode(movie)

	if err != nil {
		return
	}
}

func updateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)

			var movie Movie

			_ = json.NewDecoder(r.Body).Decode(&movie)

			movie.ID = params["id"]
			movie.Isbn = params["isbn"]
			movie.Title = params["title"]
			movie.Director = &Director{
				FirstName: params["first_name"],
				LastName:  params["last_name"],
			}

			movies = append(movies, movie)

			err := json.NewEncoder(w).Encode(movie)

			if err != nil {
				return
			}
		}
	}
}

func main() {
	r := mux.NewRouter()

	movies = append(movies, Movie{ID: "1", Isbn: "312345", Title: "Some movie", Director: &Director{
		FirstName: "First name",
		LastName:  "Last name",
	}})

	movies = append(movies, Movie{ID: "2", Isbn: "316345", Title: "Some movie, second", Director: &Director{
		FirstName: "First name, second",
		LastName:  "Last name, second",
	}})

	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Println("Starting server on port 8080")

	log.Fatal(http.ListenAndServe(":8080", r))
}
