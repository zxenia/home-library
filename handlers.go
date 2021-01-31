package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homepage).Methods("GET")
	myRouter.HandleFunc("/books", allBooks).Methods("GET")
	myRouter.HandleFunc("/book/{title}/{author}", createBook).Methods("POST")
	myRouter.HandleFunc("/book/{title}/{author}", updateBook).Methods("PUT")
	myRouter.HandleFunc("/book/{title}", deleteBook).Methods("DELETE")

	myRouter.HandleFunc("/genres", allGenres).Methods("GET")
	myRouter.HandleFunc("/genre/{name}", createGenre).Methods("POST")
	myRouter.HandleFunc("/genre/{name}", updateGenre).Methods("PUT")
	myRouter.HandleFunc("/genre/{name}", deleteGenre).Methods("DELETE")

	myRouter.HandleFunc("/places", allPlaces).Methods("GET")
	myRouter.HandleFunc("/place/{name}", createPlace).Methods("POST")
	myRouter.HandleFunc("/place/{name}", updatePlace).Methods("PUT")
	myRouter.HandleFunc("/place/{name}", deletePlace).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8081", myRouter))
}
