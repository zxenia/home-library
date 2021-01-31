package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

// var db *gorm.DB
// var err error

func homepage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Home Library")
}

// BOOKS

func allBooks(w http.ResponseWriter, r *http.Request) {
	db, err = gorm.Open("sqlite3", "library.db")
	if err != nil {
		panic("Could not connect to the database")
	}
	defer db.Close()

	var books []Book
	db.Find(&books)
	json.NewEncoder(w).Encode(books)
}

func createBook(w http.ResponseWriter, r *http.Request) {
	db, err = gorm.Open("sqlite3", "library.db")
	if err != nil {
		panic("Could not connect to the database")
	}
	defer db.Close()

	vars := mux.Vars(r)
	title := vars["title"]
	author := vars["author"]

	db.Create(&Book{Title: title, Author: author})
	fmt.Fprint(w, "New book created")

}

func deleteBook(w http.ResponseWriter, r *http.Request) {
	db, err = gorm.Open("sqlite3", "library.db")
	if err != nil {
		panic("Could not connect to the database")
	}
	defer db.Close()

	vars := mux.Vars(r)
	title := vars["title"]
	var book Book
	db.Where("title = ?", title).Find(&book)
	db.Delete(&book)
	fmt.Fprintf(w, "The book %s deleted", title)
}

func updateBook(w http.ResponseWriter, r *http.Request) {
	db, err = gorm.Open("sqlite3", "library.db")
	if err != nil {
		panic("Could not connect to the database")
	}
	defer db.Close()

	vars := mux.Vars(r)
	title := vars["title"]

	var book Book
	db.Where("title = ?", title).Find(&book)

	// user.Email = email

	db.Save(&book)
	fmt.Fprintf(w, "The book %s was updated", title)

}

// GENRES

func allGenres(w http.ResponseWriter, r *http.Request) {
	db, err = gorm.Open("sqlite3", "library.db")
	if err != nil {
		panic("Could not connect to the database")
	}
	defer db.Close()

	var genres []Genre
	db.Find(&genres)
	json.NewEncoder(w).Encode(genres)
}

func createGenre(w http.ResponseWriter, r *http.Request) {
	db, err = gorm.Open("sqlite3", "library.db")
	if err != nil {
		panic("Could not connect to the database")
	}
	defer db.Close()

	vars := mux.Vars(r)
	name := vars["name"]

	db.Create(&Genre{Name: name})
	fmt.Fprint(w, "New genre created")

}

func deleteGenre(w http.ResponseWriter, r *http.Request) {
	db, err = gorm.Open("sqlite3", "library.db")
	if err != nil {
		panic("Could not connect to the database")
	}
	defer db.Close()

	vars := mux.Vars(r)
	name := vars["name"]
	var genre Genre
	db.Where("name = ?", name).Find(&genre)
	db.Delete(&genre)
	fmt.Fprintf(w, "The genre %s deleted", name)
}

func updateGenre(w http.ResponseWriter, r *http.Request) {
	db, err = gorm.Open("sqlite3", "library.db")
	if err != nil {
		panic("Could not connect to the database")
	}
	defer db.Close()

	vars := mux.Vars(r)
	name := vars["name"]
	var genre Genre
	db.Where("name = ?", name).Find(&genre)

	// user.Email = email

	db.Save(&genre)
	fmt.Fprintf(w, "The genre %s was updated", name)

}

// PLACES

func allPlaces(w http.ResponseWriter, r *http.Request) {
	db, err = gorm.Open("sqlite3", "library.db")
	if err != nil {
		panic("Could not connect to the database")
	}
	defer db.Close()

	var places []Place
	db.Find(&places)
	json.NewEncoder(w).Encode(places)
}

func createPlace(w http.ResponseWriter, r *http.Request) {
	db, err = gorm.Open("sqlite3", "library.db")
	if err != nil {
		panic("Could not connect to the database")
	}
	defer db.Close()

	vars := mux.Vars(r)
	name := vars["name"]

	db.Create(&Place{Name: name})
	fmt.Fprint(w, "New place created")

}

func deletePlace(w http.ResponseWriter, r *http.Request) {
	db, err = gorm.Open("sqlite3", "library.db")
	if err != nil {
		panic("Could not connect to the database")
	}
	defer db.Close()

	vars := mux.Vars(r)
	name := vars["name"]
	var place Place
	db.Where("name = ?", name).Find(&place)
	db.Delete(&place)
	fmt.Fprintf(w, "The place %s deleted", name)
}

func updatePlace(w http.ResponseWriter, r *http.Request) {
	db, err = gorm.Open("sqlite3", "library.db")
	if err != nil {
		panic("Could not connect to the database")
	}
	defer db.Close()

	vars := mux.Vars(r)
	name := vars["name"]
	var place Place
	db.Where("name = ?", name).Find(&place)

	// user.Email = email

	db.Save(&place)
	fmt.Fprintf(w, "The place %s was updated", name)

}
