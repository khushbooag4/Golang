package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

//struct
type Movie struct {
	ID       string    `json: "id"`
	Isbn     string    `json: "isbn"`
	Title    string    `json: "title"`
	Director *Director `json: "director"`
}

type Director struct {
	Firstname string `json: "firstnm"`
	Lastname  string `json: "lastnm"`
}

var movie []Movie

func getMovies(w http.ResponseWriter, r *http.Request) {
	//log.Fatal("Get")
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(movie) //to convert into json
}

func deleteMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	params := mux.Vars(r)

	//looping
	for index, item := range movie {
		if item.ID == params["id"] {
			//delete
			movie = append(movie[:index], movie[index+1:]...)
			break
		}
	}

	//left movie
	json.NewEncoder(w).Encode(movie)
}

func getMoviesById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	params := mux.Vars(r)

	// we are not using index so _ is complusory
	// because it show error for extra variable
	for _, item := range movie {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

func createMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var movies Movie
	_ = json.NewDecoder(r.Body).Decode(&movies)

	//id random value
	// movie.ID = strconv.Itoa(rand.Int(1000000))
	movies.ID = strconv.Itoa(1)
	movie = append(movie, movies)

	json.NewEncoder(w).Encode(movie)
}

func updateMovies(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update")
	w.Header().Set("Content-type", "application/json")
	params := mux.Vars(r)

	// we are not using index so _ is complusory
	// because it show error for extra variable
	for index, item := range movie {
		if item.ID == params["id"] {
			movie = append(movie[:index], movie[index+1:]...)
			var movies Movie
			_ = json.NewDecoder(r.Body).Decode(&movie)

			//id random value
			//movies.ID = strconv.Itoa(rand.Int(1000000))
			movies.ID = params["id"]
			movie = append(movie, movies)

			json.NewEncoder(w).Encode(movie)
			return
		}
	}
}

// * here is the pointer to the r
func main() {
	r := mux.NewRouter()

	movie = append(movie, Movie{
		ID:    "1",
		Isbn:  "234412",
		Title: "Movie 1",
		Director: &Director{
			Firstname: "John",
			Lastname:  "Smith",
		},
	})
	movie = append(movie, Movie{
		ID:    "2",
		Isbn:  "234490",
		Title: "Movie 2",
		Director: &Director{
			Firstname: "John",
			Lastname:  "Dom",
		},
	})
	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMoviesById).Methods("GET")
	r.HandleFunc("/movies", createMovies).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovies).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovies).Methods("DELETE")

	fmt.Printf("Starting server at port 8081\n")
	if err := http.ListenAndServe(":8081", r); err != nil {
		log.Fatal(err)
	}
}
